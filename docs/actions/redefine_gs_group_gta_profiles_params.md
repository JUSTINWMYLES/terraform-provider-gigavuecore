---
page_title: "gigavuecore_redefine_gs_group_gta_profiles_params Action - gigavuecore"
subcategory: ""
description: |-
  Redefine GS Group's Gta Profiles Params
---

# gigavuecore_redefine_gs_group_gta_profiles_params Action

Redefine GS Group's Gta Profiles Params

## Example Usage

```terraform
action "gigavuecore_redefine_gs_group_gta_profiles_params" "example" {
  config {
    alias = "example"
    gta_profiles = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GS Group
* `gta_profiles` (List(String), optional)
