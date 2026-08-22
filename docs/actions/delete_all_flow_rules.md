---
page_title: "gigavuecore_delete_all_flow_rules Action - gigavuecore"
subcategory: ""
description: |-
  Delete all flow rules for a subflow
---

# gigavuecore_delete_all_flow_rules Action

Delete all flow rules for a subflow

## Example Usage

```terraform
action "gigavuecore_delete_all_flow_rules" "example" {
  config {
    flow_alias = "example"
    policy_alias = "example"
    sub_flow_alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `flow_alias` (String, required) - Flow identifier
* `policy_alias` (String, required) - Policy identifier
* `sub_flow_alias` (String, required) - Subflow identifier
