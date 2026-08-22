---
page_title: "gigavuecore_rollback_traffic_flows_migration Action - gigavuecore"
subcategory: ""
description: |-
  Rollback traffic flows migration
---

# gigavuecore_rollback_traffic_flows_migration Action

Rollback traffic flows migration

## Example Usage

```terraform
action "gigavuecore_rollback_traffic_flows_migration" "example" {
  config {
    force = true
    traffic_flows = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `force` (Bool, optional) - Force rollback even if conflicts exist
* `traffic_flows` (List(String), optional)
