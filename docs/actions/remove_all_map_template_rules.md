---
page_title: "gigavuecore_remove_all_map_template_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all map template rules
---

# gigavuecore_remove_all_map_template_rules Action

Remove all map template rules

## Example Usage

```terraform
action "gigavuecore_remove_all_map_template_rules" "example" {
  config {
    alias = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map template
* `cluster_id` (String, required) - Target Cluster ID
