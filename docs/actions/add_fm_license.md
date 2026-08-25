---
page_title: "gigavuecore_add_fm_license Action - gigavuecore"
subcategory: ""
description: |-
  Add FM License
---

# gigavuecore_add_fm_license Action

Add FM License

## Example Usage

```terraform
action "gigavuecore_add_fm_license" "example" {
  config {
    license_key = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `license_key` (String, required) - FM License Key to add


