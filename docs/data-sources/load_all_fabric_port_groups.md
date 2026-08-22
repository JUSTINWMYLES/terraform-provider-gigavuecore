---
page_title: "gigavuecore_load_all_fabric_port_groups Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all the Fabric Port Groups
---

# gigavuecore_load_all_fabric_port_groups Data Source

Get all the Fabric Port Groups

## Example Usage

```terraform
data "gigavuecore_load_all_fabric_port_groups" "example" {
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
* `giga_fabric_port_groups` (List(Object({alias, comment, health_state, health_state_reasons, port_list, port_weights, smart_lb, tags, traffic_health_state, traffic_health_state_reasons})), computed)

