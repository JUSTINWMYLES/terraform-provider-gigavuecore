---
page_title: "gigavuecore_reload_fm_instances Action - gigavuecore"
subcategory: ""
description: |-
  Reload all FM Instances in HA Group
---

# gigavuecore_reload_fm_instances Action

Reload all FM Instances in HA Group

## Example Usage

```terraform
action "gigavuecore_reload_fm_instances" "example" {
  config {
    ha_group_name = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ha_group_name` (String, required) - HA group name
