---
page_title: "gigavuecore_add_source_rules Action - gigavuecore"
subcategory: ""
description: |-
  Add source rules to a policy and sourceRulesAlias
---

# gigavuecore_add_source_rules Action

Add source rules to a policy and sourceRulesAlias

## Example Usage

```terraform
action "gigavuecore_add_source_rules" "example" {
  config {
    drop_rules = null
    pass_rules = null
    policy_alias = "example"
    source_rules_alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `drop_rules` (Set(Dynamic), optional)
* `pass_rules` (Set(Dynamic), optional)
* `policy_alias` (String, required) - Policy alias
* `source_rules_alias` (String, required) - Source rules alias
