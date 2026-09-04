---
page_title: "gigavuecore_redefine_gs_group_diameter_whitelist Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Diameter whitelist
---

# gigavuecore_redefine_gs_group_diameter_whitelist Action

Redefine GS Group's Diameter whitelist

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_diameter_whitelist" "example" {
  config {
    alias     = "example"
    whitelist = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `whitelist` (String, required) - Alias of referenced diameter Whitelist


