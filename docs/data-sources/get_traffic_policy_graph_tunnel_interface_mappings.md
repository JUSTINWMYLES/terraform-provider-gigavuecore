---
page_title: "gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings Data Source - gigavuecore"
subcategory: ""
description: |-
  Find TrafficPolicyGraph Tunnel Interface Mappings by alias
---

# gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings Data Source

Find TrafficPolicyGraph Tunnel Interface Mappings by alias

## Example Usage

```terraform
data "gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target TrafficPolicyGraph

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `traffic_policy_graph_endpoint_interface_mappings` (List(Object({monitoring_session_endpoint_iface_mappings})), computed)

