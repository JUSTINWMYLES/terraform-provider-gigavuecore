---
page_title: "gigavuecore_get_all_floating_feature_activations Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all the floating license feature activations, including the bindings associated with each activation
---

# gigavuecore_get_all_floating_feature_activations Data Source

Get all the floating license feature activations, including the bindings associated with each activation

## Example Usage

```terraform
data "gigavuecore_get_all_floating_feature_activations" "example" {
  activated = null
  aid = []
  eid = []
  end_date = null
  feature_description = []
  license_status = []
  license_type = []
  page = null
  sku = []
  sort = null
  start_date = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `activated` (Bool, optional) - Get Floating License activations by License Activation State
* `aid` (List(String), optional) - Get Floating License activations by Activation Id
* `eid` (List(String), optional) - Get Floating License activations by Entitlement Id
* `end_date` (String, optional) - Get Floating License activations by EndDate
* `feature_description` (List(String), optional) - Get Floating License activations by Feature Description
* `license_status` (List(String), optional) - Get Floating License activations by LicenseStatus
* `license_type` (List(String), optional) - Get Floating License activations by LicenseType
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sku` (List(String), optional) - Get Floating License activations by SKU
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `start_date` (String, optional) - Get Floating License activations by StartDate

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({aid, binding_errors, bindings, bundle, description, eid, end_date, grace_days, license_status, license_type, num_licenses, registration_date, revocation_code, revocation_date, revoked, sku, start_date, total_volume})), computed)

