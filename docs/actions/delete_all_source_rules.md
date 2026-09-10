---
page_title: "gigavuecore_delete_all_source_rules Action - gigavuecore"
subcategory: ""
description: |-
  Delete all source rules for a policy
---

# gigavuecore_delete_all_source_rules Action

Delete all source rules for a policy

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

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


