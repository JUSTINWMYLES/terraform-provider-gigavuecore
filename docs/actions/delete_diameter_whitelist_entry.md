---
page_title: "gigavuecore_delete_diameter_whitelist_entry Action - gigavuecore"
subcategory: ""
description: |-
  Delete Diameter Whitelist Entry
---

# gigavuecore_delete_diameter_whitelist_entry Action

Delete Diameter Whitelist Entry

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


