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

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `traffic_policy_graph_endpoint_interface_mappings` (Attributes List, computed) (see [below for nested schema](#nestedatt--traffic_policy_graph_endpoint_interface_mappings))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--traffic_policy_graph_endpoint_interface_mappings"></a>
### Nested Schema for `traffic_policy_graph_endpoint_interface_mappings`

Read-Only:

* `monitoring_session_endpoint_iface_mappings` (Attributes) - Monitoring Session Endpoint Info (see [below for nested schema](#nestedatt--traffic_policy_graph_endpoint_interface_mappings--monitoring_session_endpoint_iface_mappings))
<a id="nestedatt--traffic_policy_graph_endpoint_interface_mappings--monitoring_session_endpoint_iface_mappings"></a>
### Nested Schema for `traffic_policy_graph_endpoint_interface_mappings.monitoring_session_endpoint_iface_mappings`

Read-Only:

* `monitoring_session_id` (String)
* `vseries_endpoint_iface_mappings` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policy_graph_endpoint_interface_mappings--monitoring_session_endpoint_iface_mappings--vseries_endpoint_iface_mappings))
<a id="nestedatt--traffic_policy_graph_endpoint_interface_mappings--monitoring_session_endpoint_iface_mappings--vseries_endpoint_iface_mappings"></a>
### Nested Schema for `traffic_policy_graph_endpoint_interface_mappings.monitoring_session_endpoint_iface_mappings.vseries_endpoint_iface_mappings`

Read-Only:

* `endpoint_iface_mappings` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policy_graph_endpoint_interface_mappings--monitoring_session_endpoint_iface_mappings--vseries_endpoint_iface_mappings--endpoint_iface_mappings))
* `vseries_node_ids` (List of String) - List of V Series Node Ids
<a id="nestedatt--traffic_policy_graph_endpoint_interface_mappings--monitoring_session_endpoint_iface_mappings--vseries_endpoint_iface_mappings--endpoint_iface_mappings"></a>
### Nested Schema for `traffic_policy_graph_endpoint_interface_mappings.monitoring_session_endpoint_iface_mappings.vseries_endpoint_iface_mappings.endpoint_iface_mappings`

Read-Only:

* `endpoint_id` (String)
* `iface` (String)

