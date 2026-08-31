---
page_title: "gigavuecore_get_port_neighbors Data Source - gigavuecore"
subcategory: ""
description: |-
  Get Port Neighbors
---

# gigavuecore_get_port_neighbors Data Source

Get Port Neighbors

## Example Usage

```terraform
data "gigavuecore_get_port_neighbors" "example" {
  cluster_id = "example"
  page       = "example"
  port_id    = "example"
  slot_id    = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `port_id` (String, optional) - target device Port Id. Ignored if 'slotId' parameter is provided.
* `slot_id` (String, optional) - target device Slot Id. Takes precedence over 'portId' parameter
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `neighbors` (List of Dynamic)
* `port_id` (String) - id of the port this neighbor details are for

