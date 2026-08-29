---
page_title: "gigavuecore_update_topology_viz_configs Action - gigavuecore"
subcategory: ""
description: |-
  Configuring topology hierarchical and placement tags
---

# gigavuecore_update_topology_viz_configs Action

Configuring topology hierarchical and placement tags

## Example Usage

```terraform
action "gigavuecore_update_topology_viz_configs" "example" {
  config {
    alias_config = [{
      key   = "example"
      value = "example"
    }]
    hierarchical_tags = [ "example" ]
    link_representation_configs = [{
      color_code  = "example"
      link_format = "example"
      link_type   = "example"
    }]
    network_devices_enabled = true
    placement_tags = [{
      name       = "example"
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
    sub_group_keys     = [ "example" ]
    switch             = true
    tools_view_enabled = true
    topology_type      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias_config` (Attributes List, optional) - Node group titles can be set using this property. This is applicable only for smart sankey. (see [below for nested schema](#nestedatt--alias_config))
* `hierarchical_tags` (List of String, required) - List of topology hierarchical tag keys(Ids). This will allow the user to drill down from a top level hierarchy view to a lower level detailed view
* `link_representation_configs` (Attributes List, optional) (see [below for nested schema](#nestedatt--link_representation_configs))
* `network_devices_enabled` (Boolean, optional) - set this to 'true' in-order to view network devices information in topology view, false by default
* `placement_tags` (Attributes List, required) (see [below for nested schema](#nestedatt--placement_tags))
* `sub_group_keys` (List of String, optional)
* `switch` (Boolean, optional) - true/false, used for defining the update if for changing the topology type. true means topology type update otherwise some other properties update
* `tools_view_enabled` (Boolean, optional) - set this to 'true' in-order to view tools information in topology view, false by default
* `topology_type` (String, optional) - Topology type. Decided based on the number of managed nodes and tags configuration. 'TAG\_BASED\_SANKEY' if the topology tags are configured,'SMART\_SANKEY' if the tags are not configured and number of managed nodes is lesser than or equal to 100, 'NOT\_DEFINED' if the tags are not configured and number of managed nodes is greater than 100

<a id="nestedatt--alias_config"></a>
### Nested Schema for `alias_config`

Optional:

* `key` (String)
* `value` (String)

<a id="nestedatt--link_representation_configs"></a>
### Nested Schema for `link_representation_configs`

Optional:

* `color_code` (String) - The color code of the link
* `link_format` (String) - The format of the link
* `link_type` (String) - Represents the type of the link

<a id="nestedatt--placement_tags"></a>
### Nested Schema for `placement_tags`

Required:

* `name` (String) - Unique alias for the placement config. This will help the user to  give a meaning full name for placement tag config
* `tag_key` (String) - Unique placement tag key(Id)
* `tag_values` (List of String) - List of placement tag values associated with the specified tag key

