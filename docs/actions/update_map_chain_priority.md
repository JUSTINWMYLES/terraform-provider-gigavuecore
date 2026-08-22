---
page_title: "gigavuecore_update_map_chain_priority Action - gigavuecore"
subcategory: ""
description: |-
  Update map chain priority for a given policy and source rule
---

# gigavuecore_update_map_chain_priority Action

Update map chain priority for a given policy and source rule

## Example Usage

```terraform
action "gigavuecore_update_map_chain_priority" "example" {
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

* `alias` (String, required) - Alias is mandatory, cannot be null or blank
* `policy_alias` (String, required) - Policy alias
* `priority_configs` (List(Dynamic), optional) - List of clustered priority specifications
* `source_alias` (String, required) - SourceAlias is mandatory, cannot be null or blank
* `source_rules_alias` (String, required) - Source rules alias
