#!/usr/bin/env python3
"""
Generate TypeScript Zod schemas and a typed API client from CiviCRM API v4 getFields.

Output (written to --output-dir):
  civicrm.schemas.ts  — Zod schemas + inferred TypeScript types for every entity
  civicrm.client.ts   — typed CiviCRM API client (get/getById/create/update/delete)

Usage:
    python generate_civicrm_types.py --url https://example.org --api-key YOUR_KEY
    python generate_civicrm_types.py --url https://example.org --api-key YOUR_KEY \\
        --entities Tag,CaseType,Contact --output-dir src/api
    python generate_civicrm_types.py --url https://example.org --api-key YOUR_KEY --insecure

By default all entities are discovered via Entity.get and iterated automatically.
"""

import argparse
import json
import os
import sys
import urllib.request
import urllib.parse
import ssl

# CiviCRM data_type → Zod validator
ZOD_MAP: dict[str, str] = {
    "String":     "z.string()",
    "Integer":    "z.number().int()",
    "Float":      "z.number()",
    "Boolean":    "z.boolean()",
    "Date":       "z.string()",
    "Timestamp":  "z.string()",
    "Text":       "z.string()",
    "Blob":       "z.string()",
    "Money":      "z.number()",
    "Array":      "z.array(z.unknown())",
    "Serialized": "z.unknown()",
}

# ─────────────────────────────────────────────
# HTTP helpers
# ─────────────────────────────────────────────

def _ssl_ctx(insecure: bool) -> ssl.SSLContext:
    ctx = ssl.create_default_context()
    if insecure:
        ctx.check_hostname = False
        ctx.verify_mode = ssl.CERT_NONE
    return ctx


def _do_request(url: str, api_key: str, params_json: str, method: str, insecure: bool) -> bytes:
    ctx = _ssl_ctx(insecure)
    if method == "GET":
        full_url = url + "?" + urllib.parse.urlencode({"params": params_json})
        req = urllib.request.Request(full_url, method="GET")
    else:
        data = urllib.parse.urlencode({"params": params_json}).encode()
        req = urllib.request.Request(url, data=data, method="POST")
        req.add_header("Content-Type", "application/x-www-form-urlencoded")

    req.add_header("Authorization", f"Bearer {api_key}")
    req.add_header("X-Requested-With", "XMLHttpRequest")
    req.add_header("Accept", "application/json")

    with urllib.request.urlopen(req, context=ctx, timeout=15) as resp:
        return resp.read()


def _call(base_url: str, api_key: str, entity: str, action: str,
          params: dict, insecure: bool, label: str) -> list:
    url = f"{base_url.rstrip('/')}/civicrm/ajax/api4/{entity}/{action}"
    params_json = json.dumps(params)

    body = None
    for method in ("POST", "GET"):
        try:
            body = _do_request(url, api_key, params_json, method, insecure)
            break
        except urllib.error.HTTPError as e:
            if e.code == 405 and method == "POST":
                continue
            print(f"  HTTP {e.code} {label}: {e.reason}", file=sys.stderr)
            return []
        except Exception as e:
            print(f"  Error {label}: {e}", file=sys.stderr)
            return []

    if body is None:
        return []

    try:
        parsed = json.loads(body.decode())
    except json.JSONDecodeError:
        print(f"  Could not parse response for {label}", file=sys.stderr)
        return []

    if parsed.get("error_message"):
        print(f"  API error {label}: {parsed['error_message']}", file=sys.stderr)
        return []

    return parsed.get("values", [])

# ─────────────────────────────────────────────
# Discovery
# ─────────────────────────────────────────────

def fetch_all_entities(base_url: str, api_key: str, insecure: bool) -> list[str]:
    values = _call(base_url, api_key, "Entity", "get",
                   {"checkPermissions": False, "select": ["name"]},
                   insecure, "Entity.get")
    names = []
    for entry in values:
        name = entry.get("name") if isinstance(entry, dict) else entry
        if name:
            names.append(name)
    return sorted(names)


def fetch_fields(base_url: str, api_key: str, entity: str, insecure: bool) -> list[dict]:
    return _call(base_url, api_key, entity, "getFields",
                 {"checkPermissions": False},
                 insecure, f"{entity}.getFields")

# ─────────────────────────────────────────────
# Zod schema rendering
# ─────────────────────────────────────────────

def _zod_base(field: dict) -> str:
    data_type = field.get("data_type", "String")
    opts = field.get("options")
    if opts and isinstance(opts, dict) and opts:
        literals = ", ".join(f'"{k}"' for k in opts.keys())
        return f"z.enum([{literals}])"
    return ZOD_MAP.get(data_type, "z.unknown()")


def _zod_chain(field: dict) -> str:
    required = field.get("required", False)
    nullable = field.get("nullable", True)
    default_value = field.get("default_value")

    chain = _zod_base(field)
    if nullable and not required:
        chain += ".nullable()"
    if not required:
        chain += ".optional()"
    if default_value is not None and not required:
        if isinstance(default_value, str):
            chain += f'.default("{default_value}")'
        elif isinstance(default_value, bool):
            chain += f".default({'true' if default_value else 'false'})"
        elif isinstance(default_value, (int, float)):
            chain += f".default({default_value})"
    return chain


def render_schema(entity: str, fields: list[dict]) -> str:
    lines = [f"export const {entity}Schema = z.object({{"]
    for f in fields:
        name = f.get("name", "")
        desc = f.get("description") or f.get("title") or ""
        if desc:
            lines.append(f"  // {desc}")
        lines.append(f"  {name}: {_zod_chain(f)},")
    lines.append("})")
    lines.append(f"export type {entity} = z.infer<typeof {entity}Schema>")
    return "\n".join(lines)


def generate_schemas(entities_fields: list[tuple[str, list[dict]]]) -> str:
    blocks = [
        "// Auto-generated — do not edit manually",
        "// Run generate_civicrm_types.py to regenerate",
        'import { z } from "zod"',
        "",
    ]
    for entity, fields in entities_fields:
        blocks.append(f"\n// {'─'*56}")
        blocks.append(f"// {entity}")
        blocks.append(f"// {'─'*56}\n")
        blocks.append(render_schema(entity, fields))
        blocks.append("")
    return "\n".join(blocks)

# ─────────────────────────────────────────────
# Client rendering
# ─────────────────────────────────────────────

# The runtime base client is embedded directly into the generated file so the
# output is self-contained — no extra runtime package needed.
BASE_CLIENT = '''\
// Auto-generated — do not edit manually
// Run generate_civicrm_types.py to regenerate

export type WhereClause = [string, string, unknown?][]
export type OrderByClause = Record<string, "ASC" | "DESC">

export interface CiviCRMClientOptions {
  baseUrl: string
  apiKey: string
  /** Skip TLS verification (Node.js only, dev use) */
  checkPermissions?: boolean
}

interface ApiResponse<T> {
  values: T[]
  count: number
  error_code?: number
  error_message?: string
}

async function apiCall<T>(
  opts: CiviCRMClientOptions,
  entity: string,
  action: string,
  params: Record<string, unknown>,
): Promise<T[]> {
  const url = `${opts.baseUrl.replace(/\\/$/, "")}/civicrm/ajax/api4/${entity}/${action}`
  const body = new URLSearchParams({
    params: JSON.stringify({
      checkPermissions: opts.checkPermissions ?? false,
      ...params,
    }),
  })

  const res = await fetch(url, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${opts.apiKey}`,
      "X-Requested-With": "XMLHttpRequest",
      "Content-Type": "application/x-www-form-urlencoded",
      Accept: "application/json",
    },
    body: body.toString(),
  })

  if (!res.ok) {
    throw new Error(`CiviCRM API HTTP ${res.status}: ${res.statusText}`)
  }

  const json: ApiResponse<T> = await res.json()
  if (json.error_message) {
    throw new Error(`CiviCRM API error ${json.error_code}: ${json.error_message}`)
  }

  return json.values
}
'''


def render_entity_client(entity: str) -> str:
    return f"""\
export function create{entity}Client(opts: CiviCRMClientOptions) {{
  return {{
    async get(params: {{
      where?: WhereClause
      select?: (keyof {entity})[]
      orderBy?: OrderByClause
      limit?: number
      offset?: number
    }} = {{}}): Promise<{entity}[]> {{
      const raw = await apiCall<unknown>(opts, "{entity}", "get", params)
      return raw.map((v) => {entity}Schema.parse(v))
    }},

    async getById(id: number, select?: (keyof {entity})[]): Promise<{entity}> {{
      const results = await apiCall<unknown>(opts, "{entity}", "get", {{
        where: [["id", "=", id]],
        ...(select ? {{ select }} : {{}}),
      }})
      if (!results.length) throw new Error(`{entity} ${{id}} not found`)
      return {entity}Schema.parse(results[0])
    }},

    async create(values: Partial<Omit<{entity}, "id">>): Promise<{entity}> {{
      const results = await apiCall<unknown>(opts, "{entity}", "create", {{ values }})
      if (!results.length) throw new Error("No value returned from {entity}.create")
      return {entity}Schema.parse(results[0])
    }},

    async update(id: number, values: Partial<Omit<{entity}, "id">>): Promise<{entity}> {{
      const results = await apiCall<unknown>(opts, "{entity}", "update", {{
        where: [["id", "=", id]],
        values,
      }})
      if (!results.length) throw new Error("No value returned from {entity}.update")
      return {entity}Schema.parse(results[0])
    }},

    async delete(id: number): Promise<void> {{
      await apiCall<unknown>(opts, "{entity}", "delete", {{
        where: [["id", "=", id]],
      }})
    }},
  }}
}}
"""


def generate_client(entities: list[str]) -> str:
    blocks = [BASE_CLIENT]

    # Schema imports
    schema_imports = ", ".join(
        f"{e}Schema, type {e}" for e in entities
    )
    blocks.insert(1, f'import {{ {schema_imports} }} from "./civicrm.schemas"\n')

    for entity in entities:
        blocks.append(render_entity_client(entity))

    # Top-level factory that wires everything together
    factory_lines = ["export function createCiviCRMClient(opts: CiviCRMClientOptions) {", "  return {"]
    for entity in entities:
        factory_lines.append(f"    {entity}: create{entity}Client(opts),")
    factory_lines.append("  }")
    factory_lines.append("}")
    blocks.append("\n".join(factory_lines))

    return "\n".join(blocks)

# ─────────────────────────────────────────────
# Main
# ─────────────────────────────────────────────

def main():
    parser = argparse.ArgumentParser(
        description="Generate Zod schemas + typed CiviCRM API client from getFields"
    )
    parser.add_argument("--url", required=True, help="CiviCRM base URL (e.g. https://example.org)")
    parser.add_argument("--api-key", required=True, help="CiviCRM API key")
    parser.add_argument(
        "--entities",
        help="Comma-separated entity names (e.g. Tag,CaseType). "
             "If omitted, all entities are discovered via Entity.get.",
    )
    parser.add_argument(
        "--output-dir",
        default=".",
        help="Directory to write civicrm.schemas.ts and civicrm.client.ts (default: current dir)",
    )
    parser.add_argument("--insecure", action="store_true", help="Skip TLS certificate verification")
    args = parser.parse_args()

    if args.entities:
        entity_names = [e.strip() for e in args.entities.split(",")]
    else:
        print("Discovering all entities via Entity.get...", file=sys.stderr)
        entity_names = fetch_all_entities(args.url, args.api_key, args.insecure)
        if not entity_names:
            print("No entities found — check URL and API key.", file=sys.stderr)
            sys.exit(1)
        print(f"Found {len(entity_names)} entities.", file=sys.stderr)

    # Fetch fields for each entity, drop those with no fields
    entities_fields: list[tuple[str, list[dict]]] = []
    for entity in entity_names:
        print(f"  getFields {entity}...", file=sys.stderr)
        fields = fetch_fields(args.url, args.api_key, entity, args.insecure)
        if fields:
            entities_fields.append((entity, fields))
        else:
            print(f"    skipped (no fields)", file=sys.stderr)

    successful_entities = [e for e, _ in entities_fields]
    print(f"Generating types for {len(successful_entities)} entities...", file=sys.stderr)

    os.makedirs(args.output_dir, exist_ok=True)

    schemas_path = os.path.join(args.output_dir, "civicrm.schemas.ts")
    with open(schemas_path, "w", encoding="utf-8") as f:
        f.write(generate_schemas(entities_fields))
    print(f"Written: {schemas_path}", file=sys.stderr)

    client_path = os.path.join(args.output_dir, "civicrm.client.ts")
    with open(client_path, "w", encoding="utf-8") as f:
        f.write(generate_client(successful_entities))
    print(f"Written: {client_path}", file=sys.stderr)


if __name__ == "__main__":
    main()
