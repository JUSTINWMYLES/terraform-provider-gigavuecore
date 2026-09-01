---
page_title: "gigavuecore_node Resource - gigavuecore"
subcategory: ""
description: |-
  Update tool configuration for a GigaInsight Node
---

# gigavuecore_node Resource

Update tool configuration for a GigaInsight Node

## Example Usage

```terraform
resource "gigavuecore_node" "example" {
  env     = "example"
  node_id = "example"
  yaml    = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `env` (Dynamic, required)
* `node_id` (String, required) - The unique identifier of the GigaInsight Node
* `yaml` (Dynamic, required)

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
terraform import gigavuecore_node.example {node_id}
```
