---
page_title: "gigavuecore_remove_map_flow_rule Action - gigavuecore"
subcategory: ""
description: |-
  Remove a flowRule from a 'secondLevel/flowFilter' map
---

# gigavuecore_remove_map_flow_rule Action

Remove a flowRule from a 'secondLevel/flowFilter' map

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_remove_map_flow_rule" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rule_id    = 0
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `cluster_id` (String, required) - Target Cluster ID
* `rule_id` (Number, required) - flowRule id


