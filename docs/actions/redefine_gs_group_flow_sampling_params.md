---
page_title: "gigavuecore_redefine_gs_group_flow_sampling_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Flow Sampling Params
---

# gigavuecore_redefine_gs_group_flow_sampling_params Action

Redefine GS Group's Flow Sampling Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_flow_sampling_params" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    ip_ranges  = [ "example" ]
    rate       = 5
    timeout    = 1
    type       = "deviceIp"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `ip_ranges` (List of Dynamic, optional)
* `rate` (Number, optional) - in percent
* `timeout` (Number, optional) - in minutes
* `type` (String, optional)


