---
page_title: "gigavuecore_remove_all_map_gs_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all gsRules from a 'secondLevel/byRule' map
---

# gigavuecore_remove_all_map_gs_rules Action

Remove all gsRules from a 'secondLevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_remove_all_map_gs_rules" "example" {
  config {
    alias = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `cluster_id` (String, required) - Target Cluster ID
