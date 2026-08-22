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

* `items` (List(Object({domain_id, env, fabric_node_id, interfaces, mgmt_ip, unification_conn_id})), computed)

