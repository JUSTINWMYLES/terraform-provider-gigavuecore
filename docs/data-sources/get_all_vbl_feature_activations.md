---
page_title: "gigavuecore_get_all_vbl_feature_activations Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all the VBL license feature activations
---

# gigavuecore_get_all_vbl_feature_activations Data Source

Get all the VBL license feature activations

## Example Usage

```terraform
data "gigavuecore_get_all_vbl_feature_activations" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({aid, binding_errors, bindings, bundle, description, eid, end_date, feature, grace_days, license_status, license_type, num_licenses, registration_date, revocation_code, revocation_date, revoked, sku, start_date, total_volume})), computed)

