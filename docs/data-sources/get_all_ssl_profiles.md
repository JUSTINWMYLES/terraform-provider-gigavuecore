---
page_title: "gigavuecore_get_all_ssl_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Apps Ssl Profile
---

# gigavuecore_get_all_ssl_profiles Data Source

Get all Apps Ssl Profile

## Example Usage

```terraform
data "gigavuecore_get_all_ssl_profiles" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `apps_ssl_profiles` (List(Object({alias, cipher, mtls, version})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

