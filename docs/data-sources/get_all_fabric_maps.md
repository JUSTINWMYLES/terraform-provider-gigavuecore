---
page_title: "gigavuecore_get_all_fabric_maps Data Source - gigavuecore"
subcategory: ""
description: |-
  Get all user-defined fabric maps
---

# gigavuecore_get_all_fabric_maps Data Source

Get all user-defined fabric maps

## Example Usage

```terraform
data "gigavuecore_get_all_fabric_maps" "example" {
  alias = null
  dst_ports = null
  health_state = null
  page = null
  sort = null
  src_ports = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - Alias of fabric map to filter by
* `dst_ports` (String, optional) - Destination port of fabric map to filter by
* `health_state` (String, optional) - HealthState of fabric map to filter by. Possible values are red,green,yellow
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC and default sort field is fabric map alias. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `src_ports` (String, optional) - Source port of fabric map to filter by

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `fabric_maps` (List(Object({afm_map, child_map_aliases, config_status, creation_time, decap_aliases, encap_aliases, error_message, health_state, health_state_reasons, parent_map_aliases, pending_reason, updated_time})), computed)

