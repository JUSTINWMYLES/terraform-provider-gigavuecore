---
page_title: "gigavuecore_load_ip_destination_statuses Data Source - gigavuecore"
subcategory: ""
description: |-
  Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup
---

# gigavuecore_load_ip_destination_statuses Data Source

Load IP destination status of all or those matching the IP Interface alias or IP type or GsGroup

## Example Usage

```terraform
data "gigavuecore_load_ip_destination_statuses" "example" {
  cluster_id = null
  gs_group_alias = null
  ip_interface_alias = null
  ip_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `gs_group_alias` (String, optional) - GsGroup alias
* `ip_interface_alias` (String, optional) - IP Interface alias
* `ip_type` (String, optional) - IP type

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `ip_interfaces` (List(Object({destination, gs_group, interface_alias, status, te_id})), computed)

