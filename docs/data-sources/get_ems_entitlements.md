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

* `context` (Object({page_no, page_size, sort, total_items}), computed)
  * `page_no` (Number, computed)
  * `page_size` (Number, computed)
  * `sort` (List(String), computed)
  * `total_items` (Number, computed)
* `entitlements` (List(Object({available_licenses, description, eid, end_date, feature, gid, grace_days, license_status, license_type, num_licenses, sku, start_date})), computed)

