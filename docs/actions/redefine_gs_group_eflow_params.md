---
page_title: "gigavuecore_redefine_gs_group_eflow_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Eflow Params
---

# gigavuecore_redefine_gs_group_eflow_params Action

Redefine GS Group's Eflow Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_eflow_params" "example" {
  config {
    alias        = "example"
    cluster_id   = "example"
    enabled      = true
    interval     = 1
    log_enabled  = true
    packet_count = 1
    packet_ratio = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Boolean, optional) - Enable/Disable elephant flow detection and handling
* `interval` (Number, optional) - time interval in seconds
* `log_enabled` (Boolean, optional) - Enable/Disable logging of elephant flow parameters into gs logs
* `packet_count` (Number, optional) - Number of packets to be received by a flow
* `packet_ratio` (Number, optional) - Percentage of packets in a flow vs overall packet count


