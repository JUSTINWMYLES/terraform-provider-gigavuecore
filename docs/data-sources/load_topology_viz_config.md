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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias_config` (List(Object({key, value})), computed) - Title of node groups in smart sankey.
* `gdp_link_cache_ttl_in_secs` (Number, computed) - When a link is not reported by the GDP, how long the Topology has to maintain the unreported link. If the link reports within the configured TTL time, no changes to the existing link. But the link doesn't report then it will be deleted from Topology maintained links
* `hierarchical_tags` (List(String), computed) - List of topology hierarchical tag keys(Ids). This will allow the user to drill down from a top level hierarchy view to a lower level detailed view
* `link_representation_configs` (List(Object({color_code, link_format, link_type})), computed)
* `placement_tags` (List(Object({name, tag_key, tag_values})), computed)
* `sub_group_keys` (List(String), computed)
* `tools_view_enabled` (Bool, computed) - shows tools information in topology view if it is set to 'true'. By default it will be false
* `topology_type` (String, computed) - Topology type. Decided based on the number of managed nodes and tags configuration. 'TAG\_BASED\_SANKEY' if the topology tags are configured,'SMART\_SANKEY' if the tags are not configured and number of managed nodes is lesser than or equal to 100, 'NOT\_DEFINED' if the tags are not configured and number of managed nodes is greater than 100
* `warn` (Bool, computed) - 'true' if the tags are not configured and no. of managed nodes is more than the warning threshold(55 nodes),'false' otherwise

