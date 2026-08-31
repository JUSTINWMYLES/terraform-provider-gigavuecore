action "gigavuecore_update_topology_viz_configs" "example" {
  config {
    alias_config = [{
      key   = "example"
      value = "example"
    }]
    hierarchical_tags = [ "example" ]
    link_representation_configs = [{
      color_code  = "example"
      link_format = "SOLID"
      link_type   = "cascade"
    }]
    network_devices_enabled = true
    placement_tags = [{
      name       = "example"
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
    sub_group_keys     = [ "clusterId" ]
    switch             = true
    tools_view_enabled = true
    topology_type      = "TAG_BASED_SANKEY"
  }
}
