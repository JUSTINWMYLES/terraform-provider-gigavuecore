---
page_title: "gigavuecore_rollback_traffic_flows_migration Action - gigavuecore"
subcategory: ""
description: |-
  Rollback traffic flows migration
---

# gigavuecore_rollback_traffic_flows_migration Action

Rollback traffic flows migration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_rollback_traffic_flows_migration" "example" {
  config {
    force         = true
    traffic_flows = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `force` (Boolean, optional) - Force rollback even if conflicts exist
* `traffic_flows` (List of String, optional)


