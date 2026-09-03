---
page_title: "gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster Action - gigavuecore"
subcategory: ""
description: |-
  Migrate all flow maps to traffic flows in a cluster
---

# gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster Action

Migrate all flow maps to traffic flows in a cluster

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_migrate_all_flow_map_to_traffic_flows_in_cluster" "example" {
  config {
    cluster_id = "example"
    dry_run    = true
    force      = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `dry_run` (Boolean, optional) - Simulate migration without making changes
* `force` (Boolean, optional) - Force migration even if conflicts exist


