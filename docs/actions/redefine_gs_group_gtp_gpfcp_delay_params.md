---
page_title: "gigavuecore_redefine_gs_group_gtp_gpfcp_delay_params Action - gigavuecore"
subcategory: ""
description: |-
  new in H 5.8
---

# gigavuecore_redefine_gs_group_gtp_gpfcp_delay_params Action

new in H 5.8

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_gtp_gpfcp_delay_params" "example" {
  config {
    alias = "example"
    timeout = 1
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `timeout` (Number, optional)
