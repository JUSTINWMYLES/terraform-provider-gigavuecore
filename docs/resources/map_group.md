---
page_title: "gigavuecore_map_group Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Map Group
---

# gigavuecore_map_group Resource

Create a new Map Group

## Example Usage

```terraform
resource "gigavuecore_map_group" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  maps       = ["example"]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Map Group alias
* `cluster_id` (String, required) - id of the defining cluster
* `comment` (String, optional)
* `maps` (List of String, required) - List of maps.SecondLevel maps with flowSample-overlap or flowWhitelist-overlap subtype

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
terraform import gigavuecore_map_group.example {alias}/{cluster_id}
```
