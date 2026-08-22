---
page_title: "gigavuecore_redefine_gs_group_netflow_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Netflow Params
---

# gigavuecore_redefine_gs_group_netflow_params Action

Redefine GS Group's Netflow Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_netflow_params" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    monitor = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `monitor` (String, optional) - Alias of referenced Netflow Monitor
