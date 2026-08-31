---
page_title: "gigavuecore_map_group Resource - gigavuecore"
subcategory: ""
description: |-
  Find Map Group by alias
---

# gigavuecore_map_group Resource

Find Map Group by alias

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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_map_group.example {alias}/{cluster_id}
```
