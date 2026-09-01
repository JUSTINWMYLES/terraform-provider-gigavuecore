---
page_title: "gigavuecore_set_ems_preferences Action - gigavuecore"
subcategory: ""
description: |-
  Set opt-in preferences for entitlements/automatic renewals, usage reporting and telemetry
---

# gigavuecore_set_ems_preferences Action

Set opt-in preferences for entitlements/automatic renewals, usage reporting and telemetry

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_set_ems_preferences" "example" {
  config {
    ems_entitlements_enabled = true
    ems_telemetry_enabled    = true
    ems_usage_enabled        = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `ems_entitlements_enabled` (Boolean, optional)
* `ems_telemetry_enabled` (Boolean, optional)
* `ems_usage_enabled` (Boolean, optional)


