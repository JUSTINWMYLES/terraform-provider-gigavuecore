---
page_title: "gigavuecore_set_ems_preferences Action - gigavuecore"
subcategory: ""
description: |-
  Set opt-in preferences for entitlements/automatic renewals, usage reporting and telemetry
---

# gigavuecore_set_ems_preferences Action

Set opt-in preferences for entitlements/automatic renewals, usage reporting and telemetry

## Example Usage

```terraform
action "gigavuecore_set_ems_preferences" "example" {
  config {
    ems_entitlements_enabled = true
    ems_telemetry_enabled = true
    ems_usage_enabled = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `ems_entitlements_enabled` (Bool, optional)
* `ems_telemetry_enabled` (Bool, optional)
* `ems_usage_enabled` (Bool, optional)
