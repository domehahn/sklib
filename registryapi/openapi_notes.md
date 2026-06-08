# Registry API – OpenAPI Notes

This file documents design decisions for the SkillForge registry HTTP API DTOs
defined in `registryapi/`. It is not a generated file; it is maintained by hand.

## Versioning

- All responses include `api_version` in the capabilities response.
- Breaking changes must increment the major version segment of `api_version`.

## Endpoint conventions

| Verb   | Path                                          | Purpose                    |
|--------|-----------------------------------------------|----------------------------|
| GET    | `/capabilities`                               | Describe registry features |
| GET    | `/skills`                                     | Search skills              |
| GET    | `/skills/{namespace}/{name}`                  | Skill info + version list  |
| GET    | `/skills/{namespace}/{name}/versions`         | Version list only          |
| POST   | `/skills/{namespace}/{name}/resolve`          | Resolve a constraint       |
| GET    | `/skills/{namespace}/{name}/{version}`        | Version info               |
| GET    | `/skills/{namespace}/{name}/{version}/dl`     | Download redirect          |
| PUT    | `/skills/{namespace}/{name}/{version}`        | Publish a version          |
| POST   | `/skills/{namespace}/{name}/{version}/deprecate` | Mark deprecated         |
| POST   | `/skills/{namespace}/{name}/{version}/yank`   | Yank a version             |
| POST   | `/skills/{namespace}/{name}/{version}/unyank` | Un-yank a version          |

## Error format

All error responses use `ErrorResponse`:

```json
{
  "error": "human readable message",
  "code": "machine_readable_code",
  "details": { "field": "extra context" }
}
```

## Package upload

Publish uses multipart/form-data with the package archive as the file part.
The registry computes and verifies the SHA-256 server-side.
Clients may provide an expected `sha256` in a form field for additional verification.
