---
page_title: "gigavuecore_migrate_fabric_maps_to_traffic_flows_1 Action - gigavuecore"
subcategory: ""
description: |-
  Migrate specific fabric maps to traffic flows
---

# gigavuecore_migrate_fabric_maps_to_traffic_flows_1 Action

Migrate specific fabric maps to traffic flows

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_migrate_fabric_maps_to_traffic_flows_1" "example" {
  config {
    alias           = ["example"]
    cluster_ids     = ["example"]
    dry_run         = true
    force           = true
    migration_alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (List of String, optional)
* `cluster_ids` (List of String, optional)
* `dry_run` (Boolean, optional) - Simulate migration without making changes
* `force` (Boolean, optional) - Force migration even if conflicts exist
* `migration_alias` (String, optional)


