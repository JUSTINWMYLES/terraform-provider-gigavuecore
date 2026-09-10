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
  alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target TrafficPolicyGraph

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `monitoring_session_endpoint_iface_mappings` (Attributes) - Monitoring Session Endpoint Info (see [below for nested schema](#nestedatt--items--monitoring_session_endpoint_iface_mappings))

<a id="nestedatt--items--monitoring_session_endpoint_iface_mappings"></a>
### Nested Schema for `items.monitoring_session_endpoint_iface_mappings`

Read-Only:

* `monitoring_session_id` (String)
* `vseries_endpoint_iface_mappings` (Attributes List) (see [below for nested schema](#nestedatt--items--monitoring_session_endpoint_iface_mappings--vseries_endpoint_iface_mappings))

<a id="nestedatt--items--monitoring_session_endpoint_iface_mappings--vseries_endpoint_iface_mappings"></a>
### Nested Schema for `items.monitoring_session_endpoint_iface_mappings.vseries_endpoint_iface_mappings`

Read-Only:

* `endpoint_iface_mappings` (Attributes List) (see [below for nested schema](#nestedatt--items--monitoring_session_endpoint_iface_mappings--vseries_endpoint_iface_mappings--endpoint_iface_mappings))
* `vseries_node_ids` (List of String) - List of V Series Node Ids

<a id="nestedatt--items--monitoring_session_endpoint_iface_mappings--vseries_endpoint_iface_mappings--endpoint_iface_mappings"></a>
### Nested Schema for `items.monitoring_session_endpoint_iface_mappings.vseries_endpoint_iface_mappings.endpoint_iface_mappings`

Read-Only:

* `endpoint_id` (String)
* `iface` (String)

