---
page_title: "gigavuecore_remove_all_map_flow_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all flowRules from a 'secondLevel/flowFilter' map
---

# gigavuecore_remove_all_map_flow_rules Action

Remove all flowRules from a 'secondLevel/flowFilter' map

## Example Usage

```terraform
action "gigavuecore_remove_all_map_flow_rules" "example" {
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


