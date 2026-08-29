---
page_title: "gigavuecore_reload_fm_instance Action - gigavuecore"
subcategory: ""
description: |-
  Reload FM Instance in HA Group
---

# gigavuecore_reload_fm_instance Action

Reload FM Instance in HA Group

## Example Usage

```terraform
action "gigavuecore_reload_fm_instance" "example" {
  config {
    ha_group_name = "example"
    hostname      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `ha_group_name` (String, required) - HA group name
* `hostname` (String, required) - DNS name or IP Address


