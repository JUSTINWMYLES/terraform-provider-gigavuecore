---
page_title: "gigavuecore_port_filter Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new PortFilter
---

# gigavuecore_port_filter Resource

Create a new PortFilter

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration attribute triggers a resource replacement (all config-settable attributes carry a RequiresReplace plan modifier); create, read, and delete remain functional.

## Example Usage

```terraform
resource "gigavuecore_port_filter" "example" {
  cluster_id = "example"
  port       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `port` (String, required)

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
terraform import gigavuecore_port_filter.example {port}/{cluster_id}
```
