---
page_title: "gigavuecore_get_eligible_renewers Data Source - gigavuecore"
subcategory: ""
description: |-
  Get eligible renewable licenses
---

# gigavuecore_get_eligible_renewers Data Source

Get eligible renewable licenses

## Example Usage

```terraform
data "gigavuecore_get_eligible_renewers" "example" {
  renewee_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `renewee_id` (String, required) - Renewee ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({aid, binding_errors, bindings, bundle, description, eid, end_date, grace_days, license_status, license_type, num_licenses, registration_date, revocation_code, revocation_date, revoked, sku, start_date, total_volume})), computed)

