---
page_title: "gigavuecore_get_topoviz_links Data Source - gigavuecore"
subcategory: ""
description: |-
  Loads all topology links
---

# gigavuecore_get_topoviz_links Data Source

Loads all topology links

## Example Usage

```terraform
data "gigavuecore_get_topoviz_links" "example" {
  endpoint1_alias = null
  endpoint1_host_name = null
  endpoint2_alias = null
  endpoint2_host_name = null
  link_source = null
  link_speed = null
  link_type = null
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `endpoint1_alias` (String, optional) - Aliases of endpoint port/gigastream
* `endpoint1_host_name` (String, optional) - Hostname of endpoint device
* `endpoint2_alias` (String, optional) - Aliases of endpoint port/gigastream
* `endpoint2_host_name` (String, optional) - Hostname of endpoint devices
* `link_source` (String, optional) - Link sources
* `link_speed` (String, optional) - Link speeds
* `link_type` (String, optional) - Types of links
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({connections, endpoint1, endpoint1_cluster_name, endpoint1_global_node_id, endpoint2, endpoint2_cluster_name, endpoint2_global_node_id, link_id, link_speed, link_type})), computed)

