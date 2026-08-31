---
page_title: "gigavuecore_unconfigure_all_chassis_cards Action - gigavuecore"
subcategory: ""
description: |-
  Unconfigure all device Chassis Cards
---

# gigavuecore_unconfigure_all_chassis_cards Action

Unconfigure all device Chassis Cards

## Example Usage

```terraform
action "gigavuecore_unconfigure_all_chassis_cards" "example" {
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


