---
page_title: "gigavuecore_redefine_system_ndp Action - gigavuecore"
subcategory: ""
description: |-
  Redefine System Ndp's Config
---

# gigavuecore_redefine_system_ndp Action

Redefine System Ndp's Config

## Example Usage

```terraform
action "gigavuecore_redefine_system_ndp" "example" {
  config {
    ndp_refresh_interval = 0
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `ndp_refresh_interval` (Number, optional) - System Ndp Refresh Interval in Seconds


