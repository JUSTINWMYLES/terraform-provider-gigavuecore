---
page_title: "gigavuecore_redefine_gs_group_params_gtp_random_sampling Action - gigavuecore"
subcategory: ""
description: |-
  new in H 5.7
---

# gigavuecore_redefine_gs_group_params_gtp_random_sampling Action

new in H 5.7

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_params_gtp_random_sampling" "example" {
  config {
    alias    = "example"
    enabled  = true
    interval = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `enabled` (Boolean, optional) - When enabled, sampling of subscriber's sessions happens in random fashion
* `interval` (Number, optional) - Rotation Interval in multiples of 12 (hrs)


