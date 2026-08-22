---
page_title: "gigavuecore_validate_external_export_server Data Source - gigavuecore"
subcategory: ""
description: |-
  Validate External Export Server
---

# gigavuecore_validate_external_export_server Data Source

Validate External Export Server

## Example Usage

```terraform
data "gigavuecore_validate_external_export_server" "example" {
  export_target_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `export_target_alias` (String, required) - Alias for external export target target server

### Attributes

In addition to all arguments above, the following attributes are exported:


