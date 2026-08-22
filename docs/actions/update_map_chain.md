---
page_title: "gigavuecore_update_map_chain Action - gigavuecore"
subcategory: ""
description: |-
  Update map chain for a given policy and source rule
---

# gigavuecore_update_map_chain Action

Update map chain for a given policy and source rule

## Example Usage

```terraform
action "gigavuecore_update_map_chain" "example" {
  config {
    alias = "example"
    policy_alias = "example"
    priority_configs = null
    source_alias = "example"
    source_rules_alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `policy_alias` (String, required) - Policy alias
* `priority_configs` (List(Dynamic), optional)
* `source_alias` (String, required)
* `source_rules_alias` (String, required) - Source rules alias
