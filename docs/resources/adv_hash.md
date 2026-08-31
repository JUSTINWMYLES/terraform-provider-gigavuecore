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
  cluster_id = "example"
  fields     = [ "ethertype" ]
  slot       = "example"
  slot_id    = "example"
  type       = "all"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster
* `fields` (Set of String, optional) - Required when 'type' == 'fields'. 'gtpteid' requires fields (portsrc and portdst) or (port6src and port6dst)
* `slot` (String, required) - device card slotId for the gigastream
* `slot_id` (String, required) - target device card slot Id
* `type` (String, required)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_adv_hash.example {slot_id}/{cluster_id}
```
