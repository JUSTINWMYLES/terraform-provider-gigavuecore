---
page_title: "gigavuecore_node Resource - gigavuecore"
subcategory: ""
description: |-
  Get tool configuration for a GigaInsight Node
---

# gigavuecore_node Resource

Get tool configuration for a GigaInsight Node

## Example Usage

```terraform
resource "gigavuecore_node" "example" {
  env     = null
  node_id = null
  yaml    = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `env` (Dynamic, required)
* `node_id` (String, required)
* `yaml` (Dynamic, required)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_node.example {node_id}
```
