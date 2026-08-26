---
page_title: "gigavuecore_redefine_gs_group_gs_group_sa_apf_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's saApf Params(Deprecated: use PUT/gsGroups/{alias}/params/resource)
---

# gigavuecore_redefine_gs_group_gs_group_sa_apf_params Action

Redefine GS Group's saApf Params(Deprecated: use PUT/gsGroups/{alias}/params/resource)

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_gs_group_sa_apf_params" "example" {
  config {
    alias       = "example"
    buffer_size = 1
    cluster_id  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `buffer_size` (Number, optional) - Session aware APF buffer. Valid values: 2-5: in million sessions (0 = disabled). Changes take effect after card or system reboot
* `cluster_id` (String, required) - Target Cluster ID


