---
page_title: "gigavuecore_rollback_traffic_flows Action - gigavuecore"
subcategory: ""
description: |-
  Roll back Traffic Flow migration
---

# gigavuecore_rollback_traffic_flows Action

Roll back Traffic Flow migration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_rollback_traffic_flows" "example" {
  config {
    traffic_flows = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `traffic_flows` (List of String, optional)


