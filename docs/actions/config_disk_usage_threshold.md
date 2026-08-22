---
page_title: "gigavuecore_config_disk_usage_threshold Action - gigavuecore"
subcategory: ""
description: |-
  Config Disk Usage Threshold
---

# gigavuecore_config_disk_usage_threshold Action

Config Disk Usage Threshold

## Example Usage

```terraform
action "gigavuecore_config_disk_usage_threshold" "example" {
  config {
    disk_threshold = 1
    user = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `disk_threshold` (Number, required)
* `user` (String, required) - User name
