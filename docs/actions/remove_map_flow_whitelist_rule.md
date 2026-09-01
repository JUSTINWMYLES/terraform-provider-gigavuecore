---
page_title: "gigavuecore_remove_map_flow_whitelist_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a flowRule from a 'secondLevel/flowWhitelist' map
---

# gigavuecore_remove_map_flow_whitelist_rule Action

Remove a flowRule from a 'secondLevel/flowWhitelist' map

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_remove_map_flow_whitelist_rule" "example" {
  config {
    alias   = "example"
    rule_id = 0
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `rule_id` (Number, required) - flowWhitelistRule id


