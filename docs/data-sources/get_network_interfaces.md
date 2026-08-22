---
page_title: "gigavuecore_get_network_interfaces Data Source - gigavuecore"
subcategory: ""
description: |-
  List all interfaces of unified deployment nodes via environment and connection id
---

# gigavuecore_get_network_interfaces Data Source

List all interfaces of unified deployment nodes via environment and connection id

## Example Usage

```terraform
data "gigavuecore_get_network_interfaces" "example" {
  env_id = null
  network_id = null
  unify_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `network_id` (String, required) - Network ID
* `unify_id` (String, required) - unified resource identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Set(Object({fabric_node_id, mgmt_ip, network_id, v_series_node_interface, virt_domain_id})), computed)

