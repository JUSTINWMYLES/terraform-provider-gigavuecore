---
page_title: "gigavuecore_add_map_rules_from_template Action - gigavuecore"
subcategory: ""
description: |-
  add template rules to a 'regular/byRule', 'inline/byRule' or 'firstLevel/byRule' map
---

# gigavuecore_add_map_rules_from_template Action

add template rules to a 'regular/byRule', 'inline/byRule' or 'firstLevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_add_map_rules_from_template" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    template_alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
* `cluster_id` (String, required) - Target Cluster ID
* `template_alias` (String, required) - alias of the rules template
