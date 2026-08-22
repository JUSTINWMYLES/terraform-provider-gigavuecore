---
page_title: "gigavuecore_get_all_tcp_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all Apps TCP Profile
---

# gigavuecore_get_all_tcp_profiles Data Source

Get all Apps TCP Profile

## Example Usage

```terraform
data "gigavuecore_get_all_tcp_profiles" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `apps_tcp_profiles` (List(Object({alias, keep_alive_timer, selective_ack, syn_retries})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

