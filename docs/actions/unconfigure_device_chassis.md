---
page_title: "gigavuecore_unconfigure_device_chassis Action - gigavuecore"
subcategory: ""
description: |-
  Unconfigure Device Chassis
---

# gigavuecore_unconfigure_device_chassis Action

Unconfigure Device Chassis

## Example Usage

```terraform
action "gigavuecore_unconfigure_device_chassis" "example" {
  config {
    box_id = "example"
    cluster_id = "example"
    node_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - Box ID of the target device
* `cluster_id` (String, required) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device
