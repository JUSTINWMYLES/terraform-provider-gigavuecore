---
page_title: "gigavuecore_remove_all_map_ap_rules Action - gigavuecore"
subcategory: ""
description: |-
  Remove all apRules from a 'secondlevel/byRule' map
---

# gigavuecore_remove_all_map_ap_rules Action

Remove all apRules from a 'secondlevel/byRule' map

## Example Usage

```terraform
action "gigavuecore_remove_all_map_ap_rules" "example" {
  config {
    alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target map
