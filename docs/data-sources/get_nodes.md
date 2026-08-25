---
page_title: "gigavuecore_get_nodes Data Source - gigavuecore"
subcategory: ""
description: |-
  Get unified resource deployment nodes by environment and connection id
---

# gigavuecore_get_nodes Data Source

Get unified resource deployment nodes by environment and connection id

## Example Usage

```terraform
data "gigavuecore_get_nodes" "example" {
  env_id   = null
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `domain_id` (Attributes) - Unified Environment (see [below for nested schema](#nestedatt--items--domain_id))
* `env` (Attributes) - Unified Environment (see [below for nested schema](#nestedatt--items--env))
* `fabric_node_id` (String)
* `interfaces` (Attributes) - Unified Resource interfaces (see [below for nested schema](#nestedatt--items--interfaces))
* `mgmt_ip` (String)
* `unification_conn_id` (String)
<a id="nestedatt--items--domain_id"></a>
### Nested Schema for `items.domain_id`

Read-Only:

* `name` (String)
* `type` (String) - Unified Environment platform types
<a id="nestedatt--items--env"></a>
### Nested Schema for `items.env`

Read-Only:

* `name` (String)
* `type` (String) - Unified Environment platform types
<a id="nestedatt--items--interfaces"></a>
### Nested Schema for `items.interfaces`

Read-Only:

* `env` (Attributes Set) (see [below for nested schema](#nestedatt--items--interfaces--env))
<a id="nestedatt--items--interfaces--env"></a>
### Nested Schema for `items.interfaces.env`

Read-Only:

* `fabric_node_id` (String)
* `mgmt_ip` (String)
* `network_id` (String)
* `v_series_node_interface` (Attributes) - Cloud V Series node interface object (see [below for nested schema](#nestedatt--items--interfaces--env--v_series_node_interface))
* `virt_domain_id` (Attributes) - Virtual Domain Id for cloud resources (see [below for nested schema](#nestedatt--items--interfaces--env--virt_domain_id))
<a id="nestedatt--items--interfaces--env--v_series_node_interface"></a>
### Nested Schema for `items.interfaces.env.v_series_node_interface`

Read-Only:

* `ip_address` (String)
* `mac_address` (String)
* `name` (String)
* `public_ip_address` (String)
* `subnet_cidr` (String)
* `subnet_id` (String)
<a id="nestedatt--items--interfaces--env--virt_domain_id"></a>
### Nested Schema for `items.interfaces.env.virt_domain_id`

Read-Only:

* `name` (String)
* `type` (String)

