---
page_title: "gigavuecore_redefine_system_ndp Action - gigavuecore"
subcategory: ""
description: |-
  Redefine System Ndp's Config
---

# gigavuecore_redefine_system_ndp Action

Redefine System Ndp's Config

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_system_ndp" "example" {
  config {
    ndp_refresh_interval = 3
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `ndp_refresh_interval` (Number, optional) - System Ndp Refresh Interval in Seconds


