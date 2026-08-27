---
page_title: "gigavuecore_delete_diameter_whitelist_entry Action - gigavuecore"
subcategory: ""
description: |-
  Delete Diameter Whitelist Entry
---

# gigavuecore_delete_diameter_whitelist_entry Action

Delete Diameter Whitelist Entry

## Example Usage

```terraform
action "gigavuecore_delete_diameter_whitelist_entry" "example" {
  config {
    alias     = "example"
    user_name = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Diameter Whitelist
* `user_name` (String, optional) - imsi-based whitelist entry being deleted.if not provided deletes all entries


