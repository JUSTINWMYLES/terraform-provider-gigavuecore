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

* `fabric_node_id` (String)
* `status` (String)
* `virt_domain_id` (Attributes) - Virtual Domain Id for cloud resources (see [below for nested schema](#nestedatt--items--virt_domain_id))
<a id="nestedatt--items--virt_domain_id"></a>
### Nested Schema for `items.virt_domain_id`

Read-Only:

* `name` (String)
* `type` (String)

