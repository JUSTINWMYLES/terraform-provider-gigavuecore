---
page_title: "gigavuecore_get_node_interfaces Data Source - gigavuecore"
subcategory: ""
description: |-
  List interfaces of an unified resource deployment nodes via environment and connection id
---

# gigavuecore_get_node_interfaces Data Source

List interfaces of an unified resource deployment nodes via environment and connection id

## Example Usage

```terraform
data "gigavuecore_get_node_interfaces" "example" {
  env_id   = "example"
  node_id  = "example"
  unify_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `node_id` (String, required) - unified resource node identifier
* `unify_id` (String, required) - unified resource identifier

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes Set, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `fabric_node_id` (String)
* `mgmt_ip` (String)
* `network_id` (String)
* `v_series_node_interface` (Attributes) - Cloud V Series node interface object (see [below for nested schema](#nestedatt--items--v_series_node_interface))
* `virt_domain_id` (Attributes) - Virtual Domain Id for cloud resources (see [below for nested schema](#nestedatt--items--virt_domain_id))

<a id="nestedatt--items--v_series_node_interface"></a>
### Nested Schema for `items.v_series_node_interface`

Read-Only:

* `ip_address` (String)
* `mac_address` (String)
* `name` (String)
* `public_ip_address` (String)
* `subnet_cidr` (String)
* `subnet_id` (String)

<a id="nestedatt--items--virt_domain_id"></a>
### Nested Schema for `items.virt_domain_id`

Read-Only:

* `name` (String)
* `type` (String)

