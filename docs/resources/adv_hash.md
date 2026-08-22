---
page_title: "gigavuecore_adv_hash Resource - gigavuecore"
subcategory: ""
description: |-
  Return Gigastream Advanced Hash configuration for the given Device Card
---

# gigavuecore_adv_hash Resource

Return Gigastream Advanced Hash configuration for the given Device Card

## Example Usage

```terraform
resource "gigavuecore_adv_hash" "example" {
  cluster_id = null
  fields = []
  slot = null
  slot_id = null
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - id of the defining cluster
* `fields` (Set(String), optional) - Required when 'type' == 'fields'. 'gtpteid' requires fields (portsrc and portdst) or (port6src and port6dst)
* `slot` (String, required) - device card slotId for the gigastream
* `slot_id` (String, required)
* `type` (String, required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `cluster_id` (String, computed) - id of the defining cluster
* `fields` (Set(String), computed) - Required when 'type' == 'fields'. 'gtpteid' requires fields (portsrc and portdst) or (port6src and port6dst)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_adv_hash.example {slot_id}
```
