---
page_title: "gigavuecore_delete_map_group_traffic_flows_map_groups_policy_alias_source_rules_source_rules_alias Action - gigavuecore"
subcategory: ""
description: |-
  Delete map group by source
---

# gigavuecore_delete_map_group_traffic_flows_map_groups_policy_alias_source_rules_source_rules_alias Action

Delete map group by source

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_map_group_traffic_flows_map_groups_policy_alias_source_rules_source_rules_alias" "example" {
  config {
    policy_alias       = "example"
    source_rules_alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `policy_alias` (String, required) - Policy alias
* `source_rules_alias` (String, required) - Source rules alias


