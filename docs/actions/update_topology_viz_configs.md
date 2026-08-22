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
    alias_config = null
    hierarchical_tags = [ "example" ]
    link_representation_configs = null
    network_devices_enabled = true
    placement_tags = "example"
    sub_group_keys = [ "example" ]
    switch = true
    tools_view_enabled = true
    topology_type = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias_config` (List(Dynamic), optional) - Node group titles can be set using this property. This is applicable only for smart sankey.
* `hierarchical_tags` (List(String), required) - List of topology hierarchical tag keys(Ids). This will allow the user to drill down from a top level hierarchy view to a lower level detailed view
* `link_representation_configs` (List(Dynamic), optional)
* `network_devices_enabled` (Bool, optional) - set this to 'true' in-order to view network devices information in topology view, false by default
* `placement_tags` (List(Dynamic), required)
* `sub_group_keys` (List(String), optional)
* `switch` (Bool, optional) - true/false, used for defining the update if for changing the topology type. true means topology type update otherwise some other properties update
* `tools_view_enabled` (Bool, optional) - set this to 'true' in-order to view tools information in topology view, false by default
* `topology_type` (String, optional) - Topology type. Decided based on the number of managed nodes and tags configuration. 'TAG\_BASED\_SANKEY' if the topology tags are configured,'SMART\_SANKEY' if the tags are not configured and number of managed nodes is lesser than or equal to 100, 'NOT\_DEFINED' if the tags are not configured and number of managed nodes is greater than 100
