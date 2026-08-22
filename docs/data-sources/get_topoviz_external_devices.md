---
page_title: "gigavuecore_get_topoviz_external_devices Data Source - gigavuecore"
subcategory: ""
description: |-
  Loads all external devices in the topology
---

# gigavuecore_get_topoviz_external_devices Data Source

Loads all external devices in the topology

## Example Usage

```terraform
data "gigavuecore_get_topoviz_external_devices" "example" {
  category = null
  discovered_type = null
  model = null
  node_alias = null
  page = null
  sort = null
  type = null
  vendor = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `category` (String, optional) - Category of devices
* `discovered_type` (String, optional)
* `model` (String, optional)
* `node_alias` (String, optional) - Aliases of the nodes
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `type` (String, optional) - Type of nodes
* `vendor` (String, optional)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({comment, model, node_alias, topo_node_id, topo_node_type, type, vendor})), computed)

