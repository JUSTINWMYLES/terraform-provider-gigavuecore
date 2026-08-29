---
page_title: "gigavuecore_redefine_gs_group_dedup_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Dedup Params
---

# gigavuecore_redefine_gs_group_dedup_params Action

Redefine GS Group's Dedup Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_dedup_params" "example" {
  config {
    action     = "example"
    alias      = "example"
    cluster_id = "example"
    ip_tclass  = "example"
    ip_tos     = "example"
    tcp_seq    = "example"
    timer      = 0
    vlan       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `action` (String, optional)
* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `ip_tclass` (String, optional)
* `ip_tos` (String, optional)
* `tcp_seq` (String, optional)
* `timer` (Number, optional) - in microseconds
* `vlan` (String, optional)


