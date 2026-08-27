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
    destination_hierarchical_tags_filter = null
    destination_placement_tag_filter     = null
    destination_sub_group_filter         = null
    endpoint1_global_node_id             = "example"
    endpoint2_global_node_id             = "example"
    links_filter                         = null
    source_hierarchical_tags_filter      = null
    source_placement_tag_filter          = null
    source_sub_group_filter              = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `destination_hierarchical_tags_filter` (List of Dynamic, optional) - List of tag key and tag values which specifies the hierarchy of destination node group
* `destination_placement_tag_filter` (Dynamic, optional) - Tag key and tag values which specifies the placement of destination node group in sankey view
* `destination_sub_group_filter` (Dynamic, optional) - Property name and value based on which the nodes are subGrouped
* `endpoint1_global_node_id` (String, optional)
* `endpoint2_global_node_id` (String, optional)
* `links_filter` (Dynamic, optional) - Property name and values based on which the link are to be filtered
* `source_hierarchical_tags_filter` (List of Dynamic, optional) - List of tag key and tag values which specifies the hierarchy of source node group
* `source_placement_tag_filter` (Dynamic, optional) - Tag key and tag values which specifies the placement of source node group in sankey view
* `source_sub_group_filter` (Dynamic, optional) - Property name and value based on which the nodes are subGrouped


