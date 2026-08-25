---
page_title: "gigavuecore_update_map_group Action - gigavuecore"
subcategory: ""
description: |-
  Update map group for a given policy and source rule
---

# gigavuecore_update_map_group Action

Update map group for a given policy and source rule

## Example Usage

```terraform
action "gigavuecore_update_map_group" "example" {
  config {
    alias              = "example"
    map_group_config   = "example"
    policy_alias       = "example"
    source_alias       = "example"
    source_rules_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `map_group_config` (Dynamic, required)
* `policy_alias` (String, required) - Policy alias
* `source_alias` (String, required)
* `source_rules_alias` (String, required) - Source rules alias


