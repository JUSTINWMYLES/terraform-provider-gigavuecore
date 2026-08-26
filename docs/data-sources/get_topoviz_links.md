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
  endpoint1_alias     = null
  endpoint1_host_name = null
  endpoint2_alias     = null
  endpoint2_host_name = null
  link_source         = null
  link_speed          = null
  link_type           = null
  page                = null
  sort                = null
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `connections` (Attributes List) (see [below for nested schema](#nestedatt--items--connections))
* `endpoint1` (Attributes List) (see [below for nested schema](#nestedatt--items--endpoint1))
* `endpoint1_cluster_name` (String)
* `endpoint1_global_node_id` (String)
* `endpoint2` (Attributes List) (see [below for nested schema](#nestedatt--items--endpoint2))
* `endpoint2_cluster_name` (String)
* `endpoint2_global_node_id` (String)
* `link_id` (String)
* `link_speed` (String)
* `link_type` (String)
<a id="nestedatt--items--connections"></a>
### Nested Schema for `items.connections`

Read-Only:

* `port1` (String)
* `port1_reported` (Boolean)
* `port2` (String)
* `port2_reported` (Boolean)
<a id="nestedatt--items--endpoint1"></a>
### Nested Schema for `items.endpoint1`

Read-Only:

* `alias` (String) - PortId if endpointType is port, alias if it is a gigastream
* `endpoint_type` (String)
* `port_type` (String)
* `ports` (List of String) - List of portIds in the gigastream
<a id="nestedatt--items--endpoint2"></a>
### Nested Schema for `items.endpoint2`

Read-Only:

* `alias` (String) - PortId if endpointType is port, alias if it is a gigastream
* `endpoint_type` (String)
* `port_type` (String)
* `ports` (List of String) - List of portIds in the gigastream

