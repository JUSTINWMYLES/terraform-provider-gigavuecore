---
page_title: "gigavuecore_adv_hash Resource - gigavuecore"
subcategory: ""
description: |-
  Update Device card Gigastream Advanced Hash configuration
---

# gigavuecore_adv_hash Resource

Update Device card Gigastream Advanced Hash configuration

## Example Usage

```terraform
resource "gigavuecore_adv_hash" "example" {
  cluster_id = "example"
  fields     = ["ethertype"]
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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_adv_hash.example {slot_id}/{cluster_id}
```
