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
  cluster_id  = null
  description = null
  map_roles   = null
  name        = null
  page        = null
  sort        = null
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `description` (String, computed)
* `id` (String, computed)


