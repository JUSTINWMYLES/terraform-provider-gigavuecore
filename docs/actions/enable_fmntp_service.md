---
page_title: "gigavuecore_enable_fmntp_service Action - gigavuecore"
subcategory: ""
description: |-
  Enable/Disable NTP Service in FM
---

# gigavuecore_enable_fmntp_service Action

Enable/Disable NTP Service in FM

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_enable_fmntp_service" "example" {
  config {
    enabled = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `enabled` (Boolean, required) - NTP service enable status


