---
page_title: "gigavuecore_copy_rules Action - gigavuecore"
subcategory: ""
description: |-
  Copy rules from one or more source traffic flow policies to a user-specific clipboard. Supports both SOURCE and APPLICATION rule categories with optimistic locking for concurrent conflict detection.
---

# gigavuecore_copy_rules Action

Copy rules from one or more source traffic flow policies to a user-specific clipboard. Supports both SOURCE and APPLICATION rule categories with optimistic locking for concurrent conflict detection.

## Example Usage

```terraform
action "gigavuecore_copy_rules" "example" {
  config {
    configs = "example"
    rule_category = "example"
    rule_type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `configs` (List(Dynamic), required) - List of copy configurations
* `rule_category` (String, required) - Category of rules to copy/paste
* `rule_type` (String, required) - Type of application rule
