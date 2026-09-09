---
page_title: "gigavuecore_exporter_group Resource - gigavuecore"
subcategory: ""
description: |-
  Create Apps Exporter Group
---

# gigavuecore_exporter_group Resource

Create Apps Exporter Group

## Example Usage

```terraform
resource "gigavuecore_exporter_group" "example" {
  alias       = "example"
  cluster_id  = "example"
  description = "example"
  exporters   = ["example"]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - Target Cluster ID
* `description` (String, optional)
* `exporters` (List of String, optional)

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
terraform import gigavuecore_exporter_group.example {alias}/{cluster_id}
```
