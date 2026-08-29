---
page_title: "gigavuecore_l2_gre_group Resource - gigavuecore"
subcategory: ""
description: |-
  Load all Circuit Tunnel L2Gre Groups
---

# gigavuecore_l2_gre_group Resource

Load all Circuit Tunnel L2Gre Groups

## Example Usage

```terraform
resource "gigavuecore_l2_gre_group" "example" {
  alias      = "example"
  box_id     = "example"
  cluster_id = "example"
  comment    = "example"
  l2_gre_ids = [ 0 ]
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `box_id` (String, optional) - device box id. valid range 1 - 64.
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional)
* `l2_gre_ids` (List of Number, optional)
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_l2_gre_group.example {alias}:{cluster_id}
```
