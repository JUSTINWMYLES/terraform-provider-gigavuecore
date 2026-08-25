---
page_title: "gigavuecore_retry_device_upgrade Action - gigavuecore"
subcategory: ""
description: |-
  Retry Device Upgrade
---

# gigavuecore_retry_device_upgrade Action

Retry Device Upgrade

## Example Usage

```terraform
action "gigavuecore_retry_device_upgrade" "example" {
  config {
    device_upgrade_retry_specs = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `device_upgrade_retry_specs` (List of Dynamic, required)


