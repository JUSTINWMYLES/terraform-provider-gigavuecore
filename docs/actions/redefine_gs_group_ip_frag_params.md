---
page_title: "gigavuecore_redefine_gs_group_ip_frag_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's IP Fragmentation Params
---

# gigavuecore_redefine_gs_group_ip_frag_params Action

Redefine GS Group's IP Fragmentation Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_ip_frag_params" "example" {
  config {
    alias                = "example"
    cluster_id           = "example"
    forward              = true
    head_session_timeout = 15
    timeout              = 5
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `forward` (Boolean, optional)
* `head_session_timeout` (Number, optional) - in seconds
* `timeout` (Number, optional) - in seconds


