---
page_title: "gigavuecore_migrate_fabric_maps_to_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Migrate specific fabric maps to traffic flows in a cluster
---

# gigavuecore_migrate_fabric_maps_to_traffic_flows Action

Migrate specific fabric maps to traffic flows in a cluster

## Example Usage

```terraform
action "gigavuecore_migrate_fabric_maps_to_traffic_flows" "example" {
  config {
    alias = [ "example" ]
    cluster_id = "example"
    cluster_ids = [ "example" ]
    dry_run = true
    force = true
    migration_alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (List(String), optional)
* `cluster_id` (String, required) - Target Cluster ID
* `cluster_ids` (List(String), optional)
* `dry_run` (Bool, optional) - Simulate migration without making changes
* `force` (Bool, optional) - Force migration even if conflicts exist
* `migration_alias` (String, optional)
