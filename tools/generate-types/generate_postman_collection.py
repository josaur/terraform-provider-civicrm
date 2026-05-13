#!/usr/bin/env python3
"""
Generate a Postman Collection v2.1 from the CiviCRM API v4 entity/field definitions.

For every discovered (or specified) entity the collection contains five requests:
  GET    <entity>.get
  GET    <entity>.getById   (where id = {{<entity>_id}})
  POST   <entity>.create
  PATCH  <entity>.update
  DELETE <entity>.delete

Postman variables used in the collection:
  {{baseUrl}}   – CiviCRM base URL, e.g. https://example.org
  {{apiKey}}    – CiviCRM Bearer API key

Usage:
    python generate_postman_collection.py --url https://example.org --api-key YOUR_KEY
    python generate_postman_collection.py --url https://example.org --api-key YOUR_KEY \\
        --entities Tag,CaseType,Contact --output collection.json
    python generate_postman_collection.py --url https://example.org --api-key YOUR_KEY --insecure
"""

import argparse
import json
import os
import sys
import urllib.request
import urllib.parse
import ssl
import uuid

# ─────────────────────────────────────────────
# HTTP helpers (identical pattern to generate_civicrm_types.py)
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
# Postman Collection builders
# ─────────────────────────────────────────────

_COMMON_HEADERS = [
    {"key": "Authorization", "value": "Bearer {{apiKey}}"},
    {"key": "X-Civi-Key", "value": "{{siteKey}}"},
    {"key": "X-Requested-With", "value": "XMLHttpRequest"},
    {"key": "Accept", "value": "application/json"},
]

_FORM_HEADER = {"key": "Content-Type", "value": "application/x-www-form-urlencoded"}


def _uid() -> str:
    return str(uuid.uuid4())


def _api_url(entity: str, action: str) -> dict:
    """Return a Postman URL object for {{baseUrl}}/civicrm/ajax/api4/<entity>/<action>."""
    return {
        "raw": f"{{{{baseUrl}}}}/civicrm/ajax/api4/{entity}/{action}",
        "host": ["{{baseUrl}}"],
        "path": ["civicrm", "ajax", "api4", entity, action],
    }


def _params_body(params: dict) -> dict:
    """Postman urlencoded body with a single 'params' key."""
    return {
        "mode": "urlencoded",
        "urlencoded": [
            {
                "key": "params",
                "value": json.dumps(params, ensure_ascii=False),
                "type": "text",
            }
        ],
    }


def _build_create_values(fields: list[dict]) -> dict:
    """Build a minimal example values object from required / non-nullable fields."""
    values: dict = {}
    for f in fields:
        name = f.get("name", "")
        if name == "id":
            continue
        if f.get("required") or not f.get("nullable", True):
            data_type = f.get("data_type", "String")
            if data_type in ("Integer", "Float", "Money"):
                values[name] = 0
            elif data_type == "Boolean":
                values[name] = False
            else:
                values[name] = f"<{name}>"
    return values


def _build_select_fields(fields: list[dict]) -> list[str]:
    return [f["name"] for f in fields if f.get("name")]


def _request(name: str, method: str, entity: str, action: str,
             params: dict, description: str = "") -> dict:
    headers = list(_COMMON_HEADERS)
    body = None

    if method == "POST":
        headers.append(_FORM_HEADER)
        body = _params_body(params)
        url_obj = _api_url(entity, action)
    else:
        # GET — encode params as query string
        url_obj = dict(_api_url(entity, action))
        url_obj["query"] = [
            {
                "key": "params",
                "value": json.dumps(params, ensure_ascii=False),
            }
        ]

    item: dict = {
        "id": _uid(),
        "name": name,
        "request": {
            "method": method,
            "header": headers,
            "url": url_obj,
        },
        "response": [],
    }
    if body:
        item["request"]["body"] = body
    if description:
        item["request"]["description"] = description

    return item


def build_entity_folder(entity: str, fields: list[dict]) -> dict:
    select_fields = _build_select_fields(fields)
    create_values = _build_create_values(fields)
    id_var = f"{{{{{entity}_id}}}}"

    items = [
        _request(
            f"{entity} – get",
            "POST",
            entity, "get",
            {
                "checkPermissions": False,
                "select": select_fields[:20],  # cap for readability
                "limit": 25,
            },
            f"Fetch a list of {entity} records.",
        ),
        _request(
            f"{entity} – getById",
            "POST",
            entity, "get",
            {
                "checkPermissions": False,
                "where": [["id", "=", id_var]],
                "select": select_fields[:20],
            },
            f"Fetch a single {entity} by ID. Set the collection variable {entity}_id.",
        ),
        _request(
            f"{entity} – create",
            "POST",
            entity, "create",
            {
                "checkPermissions": False,
                "values": create_values if create_values else {"<field>": "<value>"},
            },
            f"Create a new {entity} record.",
        ),
        _request(
            f"{entity} – update",
            "POST",
            entity, "update",
            {
                "checkPermissions": False,
                "where": [["id", "=", id_var]],
                "values": {"<field>": "<new_value>"},
            },
            f"Update an existing {entity} record. Set the collection variable {entity}_id.",
        ),
        _request(
            f"{entity} – delete",
            "POST",
            entity, "delete",
            {
                "checkPermissions": False,
                "where": [["id", "=", id_var]],
            },
            f"Delete a {entity} record. Set the collection variable {entity}_id.",
        ),
    ]

    return {
        "id": _uid(),
        "name": entity,
        "item": items,
        "description": f"CiviCRM API v4 – {entity}",
    }


def build_collection(entities_fields: list[tuple[str, list[dict]]]) -> dict:
    folders = [build_entity_folder(entity, fields) for entity, fields in entities_fields]

    # One variable per entity for the ID placeholder, plus base vars
    variables = [
        {"id": _uid(), "key": "baseUrl", "value": "", "type": "string",
         "description": "CiviCRM base URL, e.g. https://example.org"},
        {"id": _uid(), "key": "apiKey", "value": "", "type": "secret",
         "description": "CiviCRM Bearer API key"},
        {"id": _uid(), "key": "siteKey", "value": "", "type": "secret",
         "description": "CiviCRM Site Key (sent as X-Civi-Key header)"},
    ]
    for entity, _ in entities_fields:
        variables.append({
            "id": _uid(),
            "key": f"{entity}_id",
            "value": "1",
            "type": "string",
            "description": f"ID used in getById / update / delete for {entity}",
        })

    return {
        "info": {
            "_postman_id": _uid(),
            "name": "CiviCRM API v4",
            "description": (
                "Auto-generated Postman Collection for CiviCRM API v4.\n"
                "Set the **baseUrl**, **apiKey**, and **siteKey** collection variables before running."
            ),
            "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
        },
        "item": folders,
        "variable": variables,
    }

# ─────────────────────────────────────────────
# Main
# ─────────────────────────────────────────────

def main():
    parser = argparse.ArgumentParser(
        description="Generate a Postman Collection v2.1 from CiviCRM API v4 getFields"
    )
    parser.add_argument("--url", required=True, help="CiviCRM base URL (e.g. https://example.org)")
    parser.add_argument("--api-key", required=True, help="CiviCRM API key")
    parser.add_argument(
        "--entities",
        help="Comma-separated entity names (e.g. Tag,CaseType). "
             "If omitted, all entities are discovered via Entity.get.",
    )
    parser.add_argument(
        "--output",
        default="civicrm_postman.json",
        help="Output file path (default: civicrm_postman.json)",
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

    entities_fields: list[tuple[str, list[dict]]] = []
    for entity in entity_names:
        print(f"  getFields {entity}...", file=sys.stderr)
        fields = fetch_fields(args.url, args.api_key, entity, args.insecure)
        if fields:
            entities_fields.append((entity, fields))
        else:
            print(f"    skipped (no fields)", file=sys.stderr)

    print(f"Building collection for {len(entities_fields)} entities...", file=sys.stderr)

    collection = build_collection(entities_fields)

    out_dir = os.path.dirname(args.output)
    if out_dir:
        os.makedirs(out_dir, exist_ok=True)

    with open(args.output, "w", encoding="utf-8") as f:
        json.dump(collection, f, ensure_ascii=False, indent=2)

    print(f"Written: {args.output}", file=sys.stderr)
    print(f"  {len(entities_fields)} entity folders, "
          f"{len(entities_fields) * 5} requests total.", file=sys.stderr)


if __name__ == "__main__":
    main()
