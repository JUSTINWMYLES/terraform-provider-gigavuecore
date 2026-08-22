---
page_title: "gigavuecore_get_node_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Obtain node status of all unified deployment nodes via environment and connection id
---

# gigavuecore_get_node_status Data Source

Obtain node status of all unified deployment nodes via environment and connection id

## Example Usage

```terraform
data "gigavuecore_get_node_status" "example" {
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

* `items` (List(Object({fabric_node_id, status, virt_domain_id})), computed)

