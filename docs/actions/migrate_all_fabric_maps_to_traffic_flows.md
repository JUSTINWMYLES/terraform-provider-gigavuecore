---
page_title: "gigavuecore_migrate_all_fabric_maps_to_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Migrate all fabric maps to traffic flows
---

# gigavuecore_migrate_all_fabric_maps_to_traffic_flows Action

Migrate all fabric maps to traffic flows

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_migrate_all_fabric_maps_to_traffic_flows" "example" {
  config {
    dry_run = true
    force   = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `dry_run` (Boolean, optional) - Simulate migration without making changes
* `force` (Boolean, optional) - Force migration even if conflicts exist


