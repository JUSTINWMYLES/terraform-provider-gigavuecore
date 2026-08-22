---
page_title: "gigavuecore_remove_fm_instance Action - gigavuecore"
subcategory: ""
description: |-
  Remove FM Instance in HA Group
---

# gigavuecore_remove_fm_instance Action

Remove FM Instance in HA Group

## Example Usage

```terraform
action "gigavuecore_remove_fm_instance" "example" {
  config {
    ha_group_name = "example"
    hostname = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ha_group_name` (String, required) - HA group name
* `hostname` (String, required) - DNS name or IP Address
