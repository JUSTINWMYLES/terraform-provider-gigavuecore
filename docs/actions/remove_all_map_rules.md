---
page_title: "gigavuecore_remove_all_map_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all rules from a 'regular/byRule', 'inline/byRule' or 'firstLevel/byRule' map
---

# gigavuecore_remove_all_map_rules Action

Remove all rules from a 'regular/byRule', 'inline/byRule' or 'firstLevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_remove_all_map_rules" "example" {
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
