---
page_title: "gigavuecore_connection Resource - gigavuecore"
subcategory: ""
description: |-
  Get unified resource connection by environment id
---

# gigavuecore_connection Resource

Get unified resource connection by environment id

## Example Usage

```terraform
resource "gigavuecore_connection" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `unified_resources_connection_get_query_response` (Dynamic, computed) - Unified Resource connection

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

