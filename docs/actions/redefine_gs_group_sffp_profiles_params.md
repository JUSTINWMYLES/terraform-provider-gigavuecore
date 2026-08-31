---
page_title: "gigavuecore_redefine_gs_group_sffp_profiles_params Action - gigavuecore"
subcategory: ""
description: |-
  new in H 5.8
---

# gigavuecore_redefine_gs_group_sffp_profiles_params Action

new in H 5.8

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_sffp_profiles_params" "example" {
  config {
    alias         = "example"
    sffp_profiles = [ "example" ]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `sffp_profiles` (List of String, optional)


