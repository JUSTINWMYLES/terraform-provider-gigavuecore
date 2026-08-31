---
page_title: "gigavuecore_load_topoviz_links Action - gigavuecore"
subcategory: ""
description: |-
  load links from given nodes or nodeGroups
---

# gigavuecore_load_topoviz_links Action

load links from given nodes or nodeGroups

## Example Usage

```terraform
action "gigavuecore_load_topoviz_links" "example" {
  config {
    destination_hierarchical_tags_filter = [{
      tag_key   = "example"
      tag_value = "example"
    }]
    destination_placement_tag_filter = {
      tag_key   = "example"
      tag_value = "example"
    }
    destination_sub_group_filter = {
      tag_key   = "example"
      tag_value = "example"
    }
    endpoint1_global_node_id = "example"
    endpoint2_global_node_id = "example"
    links_filter = {
      link_type = [ "example" ]
    }
    source_hierarchical_tags_filter = [{
      tag_key   = "example"
      tag_value = "example"
    }]
    source_placement_tag_filter = {
      tag_key   = "example"
      tag_value = "example"
    }
    source_sub_group_filter = {
      tag_key   = "example"
      tag_value = "example"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `destination_hierarchical_tags_filter` (Attributes List, optional) - List of tag key and tag values which specifies the hierarchy of destination node group (see [below for nested schema](#nestedatt--destination_hierarchical_tags_filter))
* `destination_placement_tag_filter` (Attributes, optional) - Tag key and tag values which specifies the placement of destination node group in sankey view (see [below for nested schema](#nestedatt--destination_placement_tag_filter))
* `destination_sub_group_filter` (Attributes, optional) - Property name and value based on which the nodes are subGrouped (see [below for nested schema](#nestedatt--destination_sub_group_filter))
* `endpoint1_global_node_id` (String, optional)
* `endpoint2_global_node_id` (String, optional)
* `links_filter` (Attributes, optional) - Property name and values based on which the link are to be filtered (see [below for nested schema](#nestedatt--links_filter))
* `source_hierarchical_tags_filter` (Attributes List, optional) - List of tag key and tag values which specifies the hierarchy of source node group (see [below for nested schema](#nestedatt--source_hierarchical_tags_filter))
* `source_placement_tag_filter` (Attributes, optional) - Tag key and tag values which specifies the placement of source node group in sankey view (see [below for nested schema](#nestedatt--source_placement_tag_filter))
* `source_sub_group_filter` (Attributes, optional) - Property name and value based on which the nodes are subGrouped (see [below for nested schema](#nestedatt--source_sub_group_filter))

<a id="nestedatt--destination_hierarchical_tags_filter"></a>
### Nested Schema for `destination_hierarchical_tags_filter`

Optional:

* `tag_key` (String)
* `tag_value` (String)

<a id="nestedatt--destination_placement_tag_filter"></a>
### Nested Schema for `destination_placement_tag_filter`

Optional:

* `tag_key` (String)
* `tag_value` (String)

<a id="nestedatt--destination_sub_group_filter"></a>
### Nested Schema for `destination_sub_group_filter`

Optional:

* `tag_key` (String)
* `tag_value` (String)

<a id="nestedatt--links_filter"></a>
### Nested Schema for `links_filter`

Optional:

* `link_type` (List of String) - Type of topology links - cascade, circuit, stack, tool, network

<a id="nestedatt--source_hierarchical_tags_filter"></a>
### Nested Schema for `source_hierarchical_tags_filter`

Optional:

* `tag_key` (String)
* `tag_value` (String)

<a id="nestedatt--source_placement_tag_filter"></a>
### Nested Schema for `source_placement_tag_filter`

Optional:

* `tag_key` (String)
* `tag_value` (String)

<a id="nestedatt--source_sub_group_filter"></a>
### Nested Schema for `source_sub_group_filter`

Optional:

* `tag_key` (String)
* `tag_value` (String)

