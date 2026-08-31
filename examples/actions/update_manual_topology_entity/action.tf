action "gigavuecore_update_manual_topology_entity" "example" {
  config {
    links = [{
      comment = "example"
      connections = [{
        port1 = "example"
        port2 = "example"
      }]
      endpoint1 = {
        comment         = "example"
        component_alias = "example"
        component_type  = "Port"
        port            = "example"
        topo_node_id    = "example"
        topo_node_type  = "manual"
      }
      endpoint2 = {
        comment         = "example"
        component_alias = "example"
        component_type  = "Port"
        port            = "example"
        topo_node_id    = "example"
        topo_node_type  = "manual"
      }
      topo_link_id = "example"
    }]
    nodes = [{
      comment      = "example"
      model        = "example"
      node_alias   = "example"
      topo_node_id = "example"
      type         = "example"
      vendor       = "example"
    }]
  }
}
