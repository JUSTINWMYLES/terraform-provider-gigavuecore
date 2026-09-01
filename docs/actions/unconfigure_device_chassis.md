---
page_title: "gigavuecore_unconfigure_device_chassis Action - gigavuecore"
subcategory: ""
description: |-
  Unconfigure Device Chassis
---

# gigavuecore_unconfigure_device_chassis Action

Unconfigure Device Chassis

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_unconfigure_device_chassis" "example" {
  config {
    box_id     = "example"
    cluster_id = "example"
    node_id    = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - Box ID of the target device
* `cluster_id` (String, required) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device


