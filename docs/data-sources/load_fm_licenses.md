---
page_title: "gigavuecore_load_fm_licenses Data Source - gigavuecore"
subcategory: ""
description: |-
  Load FM Licenses
---

# gigavuecore_load_fm_licenses Data Source

Load FM Licenses

## Example Usage

```terraform
data "gigavuecore_load_fm_licenses" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `licenses` (List(Object({active, description, end_date, inactive_reason, license_key, revoked, start_date, target_fm_id, valid, well_formed})), computed)

