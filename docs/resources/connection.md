---
page_title: "gigavuecore_connection Resource - gigavuecore"
subcategory: ""
description: |-
  Create unified resource connections by environment id
---

# gigavuecore_connection Resource

Create unified resource connections by environment id

~> **Note:** This resource is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

