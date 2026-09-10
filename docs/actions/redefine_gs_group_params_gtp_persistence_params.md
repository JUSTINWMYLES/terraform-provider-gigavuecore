---
page_title: "gigavuecore_redefine_gs_group_params_gtp_persistence_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Gtp Persistence Params
---

# gigavuecore_redefine_gs_group_params_gtp_persistence_params Action

Redefine GS Group's Gtp Persistence Params

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_params_gtp_persistence_params" "example" {
  config {
    alias            = "example"
    cluster_id       = "example"
    enabled          = true
    file_age_timeout = 10
    interval         = 10
    restart_age_time = 10
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Boolean, optional) - GTP Persistence Status
* `file_age_timeout` (Number, optional) - GTP Persistence File Age Timeout(mins)
* `interval` (Number, optional) - GTP Persistence Interval(mins)
* `restart_age_time` (Number, optional) - GTP Persistence Restart Age Time(mins)


