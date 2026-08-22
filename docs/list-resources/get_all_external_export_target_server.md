---
page_title: "gigavuecore_get_all_external_export_target_server List Resource - gigavuecore"
subcategory: ""
description: |-
  Load all External Export Server
---

# gigavuecore_get_all_external_export_target_server List Resource

Load all External Export Server

## Example Usage

```terraform
list "gigavuecore_get_all_external_export_target_server" "example" {
  provider = gigavuecore
  limit = 100
}

```

## Schema

### Arguments

The following arguments are supported:


### Identity Attributes

The following identity attributes are exported for each matching result:

* `export_target_alias` (String, computed)
