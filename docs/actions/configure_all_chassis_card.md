---
page_title: "gigavuecore_configure_all_chassis_card Action - gigavuecore"
subcategory: ""
description: |-
  Configure all device Chassis Cards
---

# gigavuecore_configure_all_chassis_card Action

Configure all device Chassis Cards

## Example Usage

```terraform
action "gigavuecore_configure_all_chassis_card" "example" {
  config {
    cluster_id = "example"
    node_id    = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID. Either 'clusterId' or 'nodeId' is required
* `node_id` (String, optional) - ID of the target device


