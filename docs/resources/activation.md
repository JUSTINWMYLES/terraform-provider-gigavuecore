---
page_title: "gigavuecore_activation Resource - gigavuecore"
subcategory: ""
description: |-
  Get license activations(for a specific entitlement,or for all entitlements)
---

# gigavuecore_activation Resource

Get license activations(for a specific entitlement,or for all entitlements)

## Example Usage

```terraform
resource "gigavuecore_activation" "example" {
  eli_id   = null
  page     = null
  quantity = null
  sort     = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `eli_id` (String, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `quantity` (Number, optional)
* `sort` (String, optional) - parentheses-enclosed pair of values in a (fieldName:sortOrder) format. The default sort order is DESC. Example: sort=(numLicenses:DESC)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `eid` (String, computed)
* `entl_item_id` (String, computed)
* `gid` (String, computed)
* `id` (String, computed)
* `imported` (Boolean, computed)
* `num_licenses` (Number, computed)
* `owner_fm` (Boolean, computed)
* `vmac` (String, computed)


