---
page_title: "gigavuecore_migrate_all_fabric_maps_to_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Migrate all fabric maps to traffic flows
---

# gigavuecore_migrate_all_fabric_maps_to_traffic_flows Action

Migrate all fabric maps to traffic flows

## Example Usage

```terraform
action "gigavuecore_migrate_all_fabric_maps_to_traffic_flows" "example" {
  config {
    dry_run = true
    force = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `dry_run` (Bool, optional) - Simulate migration without making changes
* `force` (Bool, optional) - Force migration even if conflicts exist
