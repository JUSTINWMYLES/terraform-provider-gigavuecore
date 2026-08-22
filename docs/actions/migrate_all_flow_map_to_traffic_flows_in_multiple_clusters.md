---
page_title: "gigavuecore_migrate_all_flow_map_to_traffic_flows_in_multiple_clusters Action - gigavuecore"
subcategory: ""
description: |-
  Migrate all flow maps to traffic flows in multiple clusters
---

# gigavuecore_migrate_all_flow_map_to_traffic_flows_in_multiple_clusters Action

Migrate all flow maps to traffic flows in multiple clusters

## Example Usage

```terraform
action "gigavuecore_migrate_all_flow_map_to_traffic_flows_in_multiple_clusters" "example" {
  config {
    alias = [ "example" ]
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
* `cluster_ids` (List(String), optional)
* `dry_run` (Bool, optional) - Simulate migration without making changes
* `force` (Bool, optional) - Force migration even if conflicts exist
* `migration_alias` (String, optional)
