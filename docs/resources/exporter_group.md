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

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_exporter_group.example {alias}/{cluster_id}
```
