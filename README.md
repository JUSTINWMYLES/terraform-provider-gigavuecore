# terraform-provider-gigavuecore
Gigamon GigaVUE-FM Core Terraform Provider

This repository holds the OpenAPI specification sources and the
[eidos](https://github.com/signalbreak-labs/eidos) `generator.yaml` used to
generate the Terraform provider code. The generated provider code itself is not
committed.

## Layout

- `generator.yaml` — eidos generator configuration (117 resources, 457 data
  sources, 554 actions, 2 ephemeral resources, 119 list resources, 3 functions).
- `spec/openapi.fm.yaml` — the upstream GigaVUE-FM 6.14.00 Core OpenAPI spec
  (contains external JSON Schema `$ref`s).
- `spec/source/**` — the external JSON Schema files referenced by the spec.
- `spec/openapi.fm.bundled.yaml` — the spec with every external ref inlined into
  local `#/components/schemas` refs (the file eidos consumes).
- `tools/bundle.py` — the bundler that produces the bundled spec.

## Regenerating

```sh
python3 tools/bundle.py          # rebuild spec/openapi.fm.bundled.yaml
eidos generate --config generator.yaml --output ./gen
```

Run `eidos generate` from the repository root so the relative `spec.path` in
`generator.yaml` resolves.
