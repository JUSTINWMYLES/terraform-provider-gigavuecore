---
page_title: "gigavuecore_leave_ha_instance Action - gigavuecore"
subcategory: ""
description: |-
  Join FM Instance to HA Group
---

# gigavuecore_leave_ha_instance Action

Join FM Instance to HA Group

## Example Usage

```terraform
action "gigavuecore_leave_ha_instance" "example" {
  config {
    ha_group_name = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ha_group_name` (String, required) - HA group name
