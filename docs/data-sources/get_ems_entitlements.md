---
page_title: "gigavuecore_get_ems_entitlements Data Source - gigavuecore"
subcategory: ""
description: |-
  Get entitlements
---

# gigavuecore_get_ems_entitlements Data Source

Get entitlements

## Example Usage

```terraform
data "gigavuecore_get_ems_entitlements" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed pair of values in a (fieldName:sortOrder) format. The default sort order is DESC. Example: sort=(sku:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) (see [below for nested schema](#nestedatt--context))
* `entitlements` (Attributes List, computed) (see [below for nested schema](#nestedatt--entitlements))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number)
* `page_size` (Number)
* `sort` (List of String)
* `total_items` (Number)
<a id="nestedatt--entitlements"></a>
### Nested Schema for `entitlements`

Read-Only:

* `available_licenses` (Number)
* `description` (String)
* `eid` (String)
* `end_date` (String)
* `feature` (String)
* `gid` (String)
* `grace_days` (Number)
* `license_status` (String)
* `license_type` (String)
* `num_licenses` (Number)
* `sku` (String)
* `start_date` (String)

