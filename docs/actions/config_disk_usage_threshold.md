---
page_title: "gigavuecore_config_disk_usage_threshold Action - gigavuecore"
subcategory: ""
description: |-
  Config Disk Usage Threshold
---

# gigavuecore_config_disk_usage_threshold Action

Config Disk Usage Threshold

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_config_disk_usage_threshold" "example" {
  config {
    disk_threshold = 60
    user           = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `disk_threshold` (Number, required)
* `user` (String, required) - User name


