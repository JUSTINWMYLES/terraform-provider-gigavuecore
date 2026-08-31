---
page_title: "gigavuecore_header_strip Resource - gigavuecore"
subcategory: ""
description: |-
  Load header strip for target box
---

# gigavuecore_header_strip Resource

Load header strip for target box

~> **Note:** The update operation for this resource is not wired to a remote API endpoint: the API spec exposes no usable update mapping. Changing any configuration fails at apply time with an explicit "not wired" diagnostic. To change this resource, replace it (for example `terraform apply -replace=...`); create, read, and delete remain functional.

## Example Usage

```terraform
resource "gigavuecore_header_strip" "example" {
  box_id      = "example"
  cluster_id  = "example"
  mpls_labels = ["example"]
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, required) - device box id. valid range 1 - 64.
* `cluster_id` (String, required) - Target cluster ID.
* `mpls_labels` (List of String, optional) - mpls ids, valid and required. Range can be specified. Example:1..200

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
terraform import gigavuecore_header_strip.example {box_id}/{cluster_id}
```
