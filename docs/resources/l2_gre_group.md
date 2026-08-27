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
  alias      = null
  box_id     = null
  cluster_id = null
  comment    = null
  l2_gre_ids = []
  page       = null
  sort       = null
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `box_id` (String, computed) - device box id. valid range 1 - 64.
* `comment` (String, computed)
* `id` (String, computed)
* `l2_gre_ids` (List of Number, computed)


