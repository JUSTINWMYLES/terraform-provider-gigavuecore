---
page_title: "gigavuecore_remove_all_map_gs_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all gsRules from a 'secondLevel/byRule' map
---

# gigavuecore_remove_all_map_gs_rules Action

Remove all gsRules from a 'secondLevel/byRule' map

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_remove_all_map_gs_rules" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `cluster_id` (String, required) - Target Cluster ID


