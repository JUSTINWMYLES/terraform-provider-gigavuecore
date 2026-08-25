---
page_title: "gigavuecore_import_license Action - gigavuecore"
subcategory: ""
description: |-
  Import FM License
---

# gigavuecore_import_license Action

Import FM License

## Example Usage

```terraform
action "gigavuecore_import_license" "example" {
  config {
    lic_file_name = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `lic_file_name` (String, required) - User uploaded license file


