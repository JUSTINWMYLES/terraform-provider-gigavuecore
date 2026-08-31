---
page_title: "gigavuecore_user_role Resource - gigavuecore"
subcategory: ""
description: |-
  Load all User Roles
---

# gigavuecore_user_role Resource

Load all User Roles

## Example Usage

```terraform
resource "gigavuecore_user_role" "example" {
  cluster_id  = "example"
  description = "example"
  map_roles   = true
  name        = "example"
  page        = "example"
  sort        = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `description` (String, optional)
* `map_roles` (Boolean, optional) - Retrieves all roles that can be accessed in maps
* `name` (String, required)
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
terraform import gigavuecore_user_role.example {name}:{cluster_id}
```
