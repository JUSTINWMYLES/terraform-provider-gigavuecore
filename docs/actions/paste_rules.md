---
page_title: "gigavuecore_paste_rules Action - gigavuecore"
subcategory: ""
description: |-
  Paste rules from the user-specific clipboard into one or more target traffic flow policies. Supports duplicate detection with two options: - DETECT\_DUPLICATES: Returns duplicate rules without saving (409 Conflict) - PROCEED\_WITHOUT\_DUPLICATES: Saves only non-duplicate rules Uses optimistic locking for concurrent conflict detection.
---

# gigavuecore_paste_rules Action

Paste rules from the user-specific clipboard into one or more target traffic flow policies. Supports duplicate detection with two options: - DETECT\_DUPLICATES: Returns duplicate rules without saving (409 Conflict) - PROCEED\_WITHOUT\_DUPLICATES: Saves only non-duplicate rules Uses optimistic locking for concurrent conflict detection.

## Example Usage

```terraform
action "gigavuecore_paste_rules" "example" {
  config {
    configs = "example"
    option = "example"
    rule_category = "example"
    rule_type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `configs` (List(Dynamic), required) - List of paste configurations
* `option` (String, optional) - Option for handling duplicate rules during paste operation
* `rule_category` (String, required) - Category of rules to copy/paste
* `rule_type` (String, required) - Type of application rule
