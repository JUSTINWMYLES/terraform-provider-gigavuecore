---
page_title: "gigavuecore_get_l2_circuit_resource_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Get current L2Circuit resource pool configuration
---

# gigavuecore_get_l2_circuit_resource_config Data Source

Get current L2Circuit resource pool configuration

## Example Usage

```terraform
data "gigavuecore_get_l2_circuit_resource_config" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `vlan_ids` (String, computed) - VLAN id ranges used for L2CIRCUIT resource type.

