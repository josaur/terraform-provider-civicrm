# CiviCRM Terraform Provider Code Generator

Generates Go source files for the Terraform provider directly from the CiviCRM API v4
`getFields` endpoint. One command produces a complete resource and data source for any
CiviCRM entity — no boilerplate to write by hand.

## Scripts

| Script | Output |
|---|---|
| `generate_terraform_provider.py` | `resource_<entity>.go` + `data_source_<entity>.go` |
| `generate_civicrm_types.py` | TypeScript Zod schemas + typed API client |
| `generate_postman_collection.py` | Postman Collection v2.1 JSON |

This document covers `generate_terraform_provider.py` and `generate_postman_collection.py`.

## Requirements

- Python 3.9+
- Network access to a running CiviCRM instance
- A CiviCRM API key with permission to call `getFields`

No third-party Python packages are needed.

## Usage

```
python generate_terraform_provider.py \
    --url <CIVICRM_URL> \
    --api-key <API_KEY> \
    --entities <ENTITY>[,<ENTITY>...] \
    [--output-dir <DIR>] \
    [--insecure] \
    [--dry-run]
```

### Options

| Flag | Required | Default | Description |
|---|---|---|---|
| `--url` | yes | — | Base URL of the CiviCRM instance (e.g. `https://example.org`) |
| `--api-key` | yes | — | CiviCRM API key |
| `--entities` | yes | — | Comma-separated list of CiviCRM entity names |
| `--output-dir` | no | `internal/provider` | Directory to write the generated `.go` files |
| `--insecure` | no | false | Skip TLS certificate verification (dev instances only) |
| `--dry-run` | no | false | Print generated code to stdout instead of writing files |
| `--provider-go` | no | `internal/provider/provider.go` | Path to `provider.go` for automatic registration |
| `--no-register` | no | false | Skip automatic registration in `provider.go` |
| `--docs-dir` | no | _(not set)_ | Generate Markdown docs under `<dir>/resources/` and `<dir>/data-sources/` |

### Examples

Generate a single entity:

```bash
python tools/generate-types/generate_terraform_provider.py \
    --url https://civicrm.example.org \
    --api-key abc123 \
    --entities MembershipStatus
```

Generate several entities at once:

```bash
python tools/generate-types/generate_terraform_provider.py \
    --url https://civicrm.example.org \
    --api-key abc123 \
    --entities PriceSet,PriceField,PriceFieldValue
```

Preview what would be generated without writing any files:

```bash
python tools/generate-types/generate_terraform_provider.py \
    --url https://civicrm.example.org \
    --api-key abc123 \
    --entities Tag \
    --dry-run
```

Write output to a custom directory:

```bash
python tools/generate-types/generate_terraform_provider.py \
    --url https://civicrm.example.org \
    --api-key abc123 \
    --entities OptionValue \
    --output-dir /tmp/generated
```

## What gets generated

For each entity the script writes two files:

### `resource_<entity_snake>.go`

A full CRUD Terraform resource implementing:

- `Create` — calls `entity.create` with all required and optional fields
- `Read` — calls `entity.get` by ID and maps the result back into state
- `Update` — calls `entity.update`, explicitly setting removed optional fields to `nil`
- `Delete` — calls `entity.delete`
- `ImportState` — supports `terraform import` by numeric ID
- `mapResultToState` — maps the raw API response onto the Go model struct

### `data_source_<entity_snake>.go`

A read-only data source that looks up a single record. If the entity has a `name` field,
both `id` and `name` can be used as lookup keys (at least one is required). Otherwise only
`id` is accepted.

## Type mapping

CiviCRM `data_type` values are mapped to Terraform plugin framework types as follows:

| CiviCRM `data_type` | Go type | Schema attribute |
|---|---|---|
| `Integer` | `types.Int64` | `schema.Int64Attribute` |
| `Float` | `types.Float64` | `schema.Float64Attribute` |
| `Money` | `types.Float64` | `schema.Float64Attribute` |
| `Boolean` | `types.Bool` | `schema.BoolAttribute` |
| `String` | `types.String` | `schema.StringAttribute` |
| `Text` | `types.String` | `schema.StringAttribute` |
| `Date` / `Timestamp` | `types.String` | `schema.StringAttribute` |
| `Array` / `Blob` | `types.String` | `schema.StringAttribute` |

`Money` fields receive special handling: CiviCRM can return them as either a JSON number
or a string (e.g. `"12.50"`), so the generated reader handles both cases.

## Field cardinality rules

The generator derives `Required` / `Optional+Computed` from the API metadata:

| Condition | Schema modifier |
|---|---|
| `required: true` | `Required: true` |
| `required: false`, any type except Bool | `Optional: true, Computed: true` |
| `required: false`, Bool | `Optional: true, Computed: true, Default: booldefault.StaticBool(false)` |
| `id` field | `Computed: true` + `UseStateForUnknown` plan modifier |

## After generation

### 1. Register the new types in `provider.go`

The script prints a hint at the end:

```
── Registration hint for provider.go ──────────────────
Resources():
  NewMembershipStatusResource,
DataSources():
  NewMembershipStatusDataSource,
```

Add these to the `Resources()` and `DataSources()` slices in `internal/provider/provider.go`.

### 2. Verify it compiles

```bash
go build ./...
```

### 3. Review defaults for Bool fields

The generator uses `booldefault.StaticBool(false)` for all Boolean fields. Check the
CiviCRM documentation or a live record for fields where the real default is `true`
(common examples: `is_active`, `is_required`, `is_display_amounts`) and adjust accordingly.

### 4. Review Money / serialized fields

Fields with `serialize: 1` in the API (e.g. `extends` on `PriceSet`) are treated as plain
strings — this matches how the CiviCRM API v4 returns them. No change needed in most cases.

## Limitations

- **No custom field support** — custom fields (`custom.*`) are not fetched or generated.
- **No FK label resolution** — foreign key fields (e.g. `financial_type_id`) are generated
  as plain integers. The human-readable `:name` / `:label` suffixes are not wired up.
- **Bool defaults require review** — the API metadata does not reliably expose default
  values, so all booleans default to `false` in the generated code.
- **No list/set attributes** — fields typed as `Array` are rendered as strings containing
  the serialized value, consistent with how the provider handles `extends` on `PriceSet`.

---

## Postman Collection Generator

`generate_postman_collection.py` queries the same CiviCRM `getFields` endpoint and writes
a **Postman Collection v2.1** JSON file that can be imported directly into Postman or Bruno.

### What gets generated

For every entity a folder is created containing five requests:

| Request | Action | Description |
|---|---|---|
| `<Entity> – get` | `entity.get` | List records (first 25, all fields) |
| `<Entity> – getById` | `entity.get` | Fetch one record by `{{<Entity>_id}}` |
| `<Entity> – create` | `entity.create` | Create a record (required fields pre-filled) |
| `<Entity> – update` | `entity.update` | Update a record by `{{<Entity>_id}}` |
| `<Entity> – delete` | `entity.delete` | Delete a record by `{{<Entity>_id}}` |

All requests use `POST` with the CiviCRM `params` payload as `application/x-www-form-urlencoded`.

### Collection variables

| Variable | Description |
|---|---|
| `baseUrl` | CiviCRM base URL, e.g. `https://example.org` |
| `apiKey` | CiviCRM Bearer API key (sent as `Authorization: Bearer …`) |
| `siteKey` | CiviCRM Site Key (sent as `X-Civi-Key` header) |
| `<Entity>_id` | Record ID used in `getById`, `update`, and `delete` (one per entity, default `1`) |

### Usage

**Generate the collection:**

```bash
# All entities (auto-discovered)
python tools/generate-types/generate_postman_collection.py \
    --url https://civicrm.example.org \
    --api-key YOUR_KEY

# Specific entities only
python tools/generate-types/generate_postman_collection.py \
    --url https://civicrm.example.org \
    --api-key YOUR_KEY \
    --entities Tag,Contact,Case \
    --output my_collection.json

# Self-signed / dev TLS
python tools/generate-types/generate_postman_collection.py \
    --url https://civicrm.example.org \
    --api-key YOUR_KEY \
    --insecure
```

**Import into Postman:**

1. Open Postman → **Import** → drag the `.json` file onto the dialog.
2. Open the collection → **Variables** tab.
3. Set `baseUrl`, `apiKey`, and `siteKey` in the **Variables** tab.
4. To test a specific record set the matching `<Entity>_id` variable.
5. Run any request — or use the **Collection Runner** to execute a whole entity folder.

**Import into Bruno:**

1. **Import Collection** → select the `.json` file → choose *Postman Collection v2*.
2. Set `baseUrl` and `apiKey` as environment variables.
3. Use Bruno's environment switcher to point at dev/staging/prod.

### CLI reference

| Flag | Required | Default | Description |
|---|---|---|---|
| `--url` | yes | — | CiviCRM base URL |
| `--api-key` | yes | — | CiviCRM API key |
| `--entities` | no | all | Comma-separated entity names |
| `--output` | no | `civicrm_postman.json` | Output file path |
| `--insecure` | no | false | Skip TLS verification |
