---
page_title: "gigavuecore_del_key_map_from_hsm_group Action - gigavuecore"
subcategory: ""
description: |-
  Delete keymaps from HSM Group
---

# gigavuecore_del_key_map_from_hsm_group Action

Delete keymaps from HSM Group

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_del_key_map_from_hsm_group" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rule_id    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.
* `rule_id` (String, required) - ruleId of keymap to delete, 0 = all


