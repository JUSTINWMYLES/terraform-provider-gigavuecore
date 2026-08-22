---
page_title: "gigavuecore_get_ems_preferences Data Source - gigavuecore"
subcategory: ""
description: |-
  Get opt-in preferences for entitlements/automatic renewals, usage reporting and telemetry
---

# gigavuecore_get_ems_preferences Data Source

Get opt-in preferences for entitlements/automatic renewals, usage reporting and telemetry

## Example Usage

```terraform
data "gigavuecore_get_ems_preferences" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `ems_entitlements_enabled` (Bool, computed)
* `ems_telemetry_enabled` (Bool, computed)
* `ems_usage_enabled` (Bool, computed)

