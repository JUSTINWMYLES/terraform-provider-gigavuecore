---
page_title: "gigavuecore_copy_rules Action - gigavuecore"
subcategory: ""
description: |-
  Copy rules from one or more source traffic flow policies to a user-specific clipboard. Supports both SOURCE and APPLICATION rule categories with optimistic locking for concurrent conflict detection.
---

# gigavuecore_copy_rules Action

Copy rules from one or more source traffic flow policies to a user-specific clipboard.
Supports both SOURCE and APPLICATION rule categories with optimistic locking for concurrent conflict detection.

## Example Usage

```terraform
action "gigavuecore_copy_rules" "example" {
  config {
    configs = [{
      flow_alias            = "example"
      is_all_rules          = true
      policy_id_or_alias    = "example"
      policy_updated_time   = 0
      rule_ids              = [ 0 ]
      source_and_rule_alias = "example"
      sub_flow_alias        = "example"
    }]
    rule_category = "example"
    rule_type     = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `configs` (Attributes List, required) - List of copy configurations (see [below for nested schema](#nestedatt--configs))
* `rule_category` (String, required) - Category of rules to copy/paste
* `rule_type` (String, required) - Type of application rule

<a id="nestedatt--configs"></a>
### Nested Schema for `configs`

Required:

* `policy_id_or_alias` (String) - Policy ID or alias from which to copy rules
* `policy_updated_time` (Number) - Timestamp of the policy version for optimistic locking

Optional:

* `flow_alias` (String) - Flow alias (required for APPLICATION category)
* `is_all_rules` (Boolean) - Whether to copy all rules or specific rule IDs
* `rule_ids` (List of Number) - Specific rule IDs to copy (if isAllRules is false)
* `source_and_rule_alias` (String) - Source and rule alias (required for SOURCE category)
* `sub_flow_alias` (String) - Sub-flow alias (required for APPLICATION category)

