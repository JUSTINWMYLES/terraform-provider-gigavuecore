---
page_title: "gigavuecore_redefine_gs_group_xpkt_match Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Xpkt match
---

# gigavuecore_redefine_gs_group_xpkt_match Action

Redefine GS Group's Xpkt match

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_xpkt_match" "example" {
  config {
    alias   = "example"
    enabled = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `enabled` (Boolean, optional)


