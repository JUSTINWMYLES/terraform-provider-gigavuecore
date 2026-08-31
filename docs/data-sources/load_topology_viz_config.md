---
page_title: "gigavuecore_load_topology_viz_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load FM global topology visualization configurations
---

# gigavuecore_load_topology_viz_config Data Source

Load FM global topology visualization configurations

## Example Usage

```terraform
data "gigavuecore_load_topology_viz_config" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias_config` (Attributes List, computed) - Title of node groups in smart sankey. (see [below for nested schema](#nestedatt--alias_config))
* `gdp_link_cache_ttl_in_secs` (Number, computed) - When a link is not reported by the GDP, how long the Topology has to maintain the unreported link. If the link reports within the configured TTL time, no changes to the existing link. But the link doesn't report then it will be deleted from Topology maintained links
* `hierarchical_tags` (List of String, computed) - List of topology hierarchical tag keys(Ids). This will allow the user to drill down from a top level hierarchy view to a lower level detailed view
* `link_representation_configs` (Attributes List, computed) (see [below for nested schema](#nestedatt--link_representation_configs))
* `placement_tags` (Attributes List, computed) (see [below for nested schema](#nestedatt--placement_tags))
* `sub_group_keys` (List of String, computed)
* `tools_view_enabled` (Boolean, computed) - shows tools information in topology view if it is set to 'true'. By default it will be false
* `topology_type` (String, computed) - Topology type. Decided based on the number of managed nodes and tags configuration. 'TAG\_BASED\_SANKEY' if the topology tags are configured,'SMART\_SANKEY' if the tags are not configured and number of managed nodes is lesser than or equal to 100, 'NOT\_DEFINED' if the tags are not configured and number of managed nodes is greater than 100
* `warn` (Boolean, computed) - 'true' if the tags are not configured and no. of managed nodes is more than the warning threshold(55 nodes),'false' otherwise

<a id="nestedatt--alias_config"></a>
### Nested Schema for `alias_config`

Read-Only:

* `key` (String)
* `value` (String)

<a id="nestedatt--link_representation_configs"></a>
### Nested Schema for `link_representation_configs`

Read-Only:

* `color_code` (String) - The color code of the link
* `link_format` (String) - The format of the link
* `link_type` (String) - Represents the type of the link

<a id="nestedatt--placement_tags"></a>
### Nested Schema for `placement_tags`

Read-Only:

* `name` (String) - Unique alias for the placement config. This will help the user to  give a meaning full name for placement tag config
* `tag_key` (String) - Unique placement tag key(Id)
* `tag_values` (List of String) - List of placement tag values associated with the specified tag key

