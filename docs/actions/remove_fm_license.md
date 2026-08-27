---
page_title: "gigavuecore_remove_fm_license Action - gigavuecore"
subcategory: ""
description: |-
  Remove FM License
---

# gigavuecore_remove_fm_license Action

Remove FM License

## Example Usage

```terraform
action "gigavuecore_remove_fm_license" "example" {
  config {
    license_key = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `license_key` (String, required) - FM License Key to remove


