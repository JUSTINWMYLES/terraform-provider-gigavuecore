---
page_title: "gigavuecore_redefine_gs_group_params_gtp_control_sampling Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Gtp Control Sampling Params
---

# gigavuecore_redefine_gs_group_params_gtp_control_sampling Action

Redefine GS Group's Gtp Control Sampling Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_params_gtp_control_sampling" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    enabled = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Bool, optional) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.
