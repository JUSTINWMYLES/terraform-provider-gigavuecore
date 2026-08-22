---
page_title: "gigavuecore_unconfigure_chassis_card Action - gigavuecore"
subcategory: ""
description: |-
  Unconfigure a device Chassis Card
---

# gigavuecore_unconfigure_chassis_card Action

Unconfigure a device Chassis Card

## Example Usage

```terraform
action "gigavuecore_unconfigure_chassis_card" "example" {
  config {
    cluster_id = "example"
    node_id = "example"
    slot_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device
* `slot_id` (String, required) - Device card slot ID
