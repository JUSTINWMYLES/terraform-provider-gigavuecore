---
page_title: "gigavuecore_delete_all_source_rules Action - gigavuecore"
subcategory: ""
description: |-
  Delete all source rules for a policy
---

# gigavuecore_delete_all_source_rules Action

Delete all source rules for a policy

## Example Usage

```terraform
action "gigavuecore_delete_all_source_rules" "example" {
  config {
    policy_alias       = "example"
    source_rules_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `policy_alias` (String, required) - Policy identifier
* `source_rules_alias` (String, required) - Source rules identifier


