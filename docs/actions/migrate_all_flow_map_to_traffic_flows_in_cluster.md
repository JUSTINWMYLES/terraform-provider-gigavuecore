---
page_title: "gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster Action - gigavuecore"
subcategory: ""
description: |-
  Migrate all flow maps to traffic flows in a cluster
---

# gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster Action

Migrate all flow maps to traffic flows in a cluster

## Example Usage

```terraform
action "gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster" "example" {
  config {
    cluster_id = "example"
    dry_run = true
    force = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `dry_run` (Bool, optional) - Simulate migration without making changes
* `force` (Bool, optional) - Force migration even if conflicts exist
