---
page_title: "gigavuecore_configure_chassis_card Action - gigavuecore"
subcategory: ""
description: |-
  Configure a device Chassis Card
---

# gigavuecore_configure_chassis_card Action

Configure a device Chassis Card

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_configure_chassis_card" "example" {
  config {
    cluster_id = "example"
    node_id    = "example"
    slot_id    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device
* `slot_id` (String, required) - Device card slot ID


