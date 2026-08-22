---
page_title: "gigavuecore_get_interfaces Data Source - gigavuecore"
subcategory: ""
description: |-
  Get interfaces of all unified resource deployment nodes via environment and connection id
---

# gigavuecore_get_interfaces Data Source

Get interfaces of all unified resource deployment nodes via environment and connection id

## Example Usage

```terraform
data "gigavuecore_get_interfaces" "example" {
  env_id = null
  unify_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Set(Object({fabric_node_id, mgmt_ip, network_id, v_series_node_interface, virt_domain_id})), computed)

