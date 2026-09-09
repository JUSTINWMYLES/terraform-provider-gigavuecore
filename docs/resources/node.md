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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_node.example {node_id}
```
