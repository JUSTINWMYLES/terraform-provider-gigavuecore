---
page_title: "gigavuecore_resume_device_upgrade Action - gigavuecore"
subcategory: ""
description: |-
  Resume Device Upgrade
---

# gigavuecore_resume_device_upgrade Action

Resume Device Upgrade

## Example Usage

```terraform
action "gigavuecore_resume_device_upgrade" "example" {
  config {
    device_upgrade_resume_specs = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `device_upgrade_resume_specs` (List(Dynamic), required)
