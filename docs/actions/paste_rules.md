---
page_title: "gigavuecore_paste_rules Action - gigavuecore"
subcategory: ""
description: |-
  Paste rules from the user-specific clipboard into one or more target traffic flow policies. Supports duplicate detection with two options: - DETECT\_DUPLICATES: Returns duplicate rules without saving (409 Conflict) - PROCEED\_WITHOUT\_DUPLICATES: Saves only non-duplicate rules Uses optimistic locking for concurrent conflict detection.
---

# gigavuecore_paste_rules Action

Paste rules from the user-specific clipboard into one or more target traffic flow policies.
Supports duplicate detection with two options:
- DETECT_DUPLICATES: Returns duplicate rules without saving (409 Conflict)
- PROCEED_WITHOUT_DUPLICATES: Saves only non-duplicate rules
Uses optimistic locking for concurrent conflict detection.

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_paste_rules" "example" {
  config {
    configs = [{
      flow_alias            = "example"
      policy_id_or_alias    = "example"
      policy_updated_time   = 0
      source_and_rule_alias = "example"
      sub_flow_alias        = "example"
    }]
    option        = "DETECT_DUPLICATES"
    rule_category = "SOURCE"
    rule_type     = "byRule"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `configs` (Attributes List, required) - List of paste configurations (see [below for nested schema](#nestedatt--configs))
* `option` (String, optional) - Option for handling duplicate rules during paste operation
* `rule_category` (String, required) - Category of rules to copy/paste
* `rule_type` (String, required) - Type of application rule

<a id="nestedatt--configs"></a>
### Nested Schema for `configs`

Required:

* `policy_id_or_alias` (String) - Policy ID or alias to which to paste rules
* `policy_updated_time` (Number) - Timestamp of the policy version for optimistic locking

Optional:

* `flow_alias` (String) - Flow alias (required for APPLICATION category)
* `source_and_rule_alias` (String) - Source and rule alias (required for SOURCE category)
* `sub_flow_alias` (String) - Sub-flow alias (required for APPLICATION category)

