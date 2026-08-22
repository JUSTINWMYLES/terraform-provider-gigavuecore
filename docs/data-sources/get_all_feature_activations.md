---
page_title: "gigavuecore_get_all_feature_activations Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all the pooled license feature activations (floating and VBL), including the bindings associated with each floating activation
---

# gigavuecore_get_all_feature_activations Data Source

Get all the pooled license feature activations (floating and VBL), including the bindings associated with each floating activation

## Example Usage

```terraform
data "gigavuecore_get_all_feature_activations" "example" {
  sku = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `sku` (String, optional) - Get pooled License activations (floating and VBL) by SKU

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({aid, binding_errors, bindings, bundle, description, eid, end_date, grace_days, license_status, license_type, num_licenses, registration_date, revocation_code, revocation_date, revoked, sku, start_date, total_volume})), computed)

