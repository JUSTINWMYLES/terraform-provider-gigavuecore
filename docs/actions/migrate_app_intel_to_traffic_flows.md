---
page_title: "gigavuecore_migrate_app_intel_to_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Migrate AppIntel solution to Traffic Flows
---

# gigavuecore_migrate_app_intel_to_traffic_flows Action

Migrate AppIntel solution to Traffic Flows

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_migrate_app_intel_to_traffic_flows" "example" {
  config {
    alias           = "example"
    cluster_id      = "example"
    dry_run         = true
    force           = true
    migration_alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional)
* `cluster_id` (String, optional)
* `dry_run` (Boolean, optional) - Simulate migration without making changes
* `force` (Boolean, optional) - Force migration even if conflicts exist
* `migration_alias` (String, optional)


