action "gigavuecore_update_topology_link" "example" {
  config {
    body_topo_link_id = "example"
    comment           = "example"
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
  }
}
