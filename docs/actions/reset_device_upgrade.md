---
page_title: "gigavuecore_reset_device_upgrade Action - gigavuecore"
subcategory: ""
description: |-
  Reset Device Upgrade
---

# gigavuecore_reset_device_upgrade Action

Reset Device Upgrade

## Example Usage

```terraform
action "gigavuecore_reset_device_upgrade" "example" {
  config {
    device_upgrade_reset_specs = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `device_upgrade_reset_specs` (List(Dynamic), required)
