---
page_title: "gigavuecore_get_all_tunnel_logical_groups Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All Tunnel Logical Groups
---

# gigavuecore_get_all_tunnel_logical_groups Data Source

Get All Tunnel Logical Groups

## Example Usage

```terraform
data "gigavuecore_get_all_tunnel_logical_groups" "example" {
  alias = null
  destination_id = null
  is_active = null
  page = null
  sort = null
  source_id = null
  state = null
  tunnel_id = null
  tunnel_type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional)
* `destination_id` (String, optional)
* `is_active` (String, optional)
* `page` (String, optional)
* `sort` (String, optional)
* `source_id` (String, optional)
* `state` (String, optional)
* `tunnel_id` (String, optional)
* `tunnel_type` (String, optional)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `tunnel_logical_groups` (List(Object({alert_policy_name, alias, custom_alias, decap_tunnel_infos, encap_tunnel_infos, health_state, health_state_reasons, is_active, key, state, tunnel_id, tunnel_type})), computed)

