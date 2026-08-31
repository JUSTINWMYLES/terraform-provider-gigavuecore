---
page_title: "gigavuecore_sysdump Resource - gigavuecore"
subcategory: ""
description: |-
  Download Sysdump File
---

# gigavuecore_sysdump Resource

Download Sysdump File

## Example Usage

```terraform
resource "gigavuecore_sysdump" "example" {
}
```

## Schema

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

