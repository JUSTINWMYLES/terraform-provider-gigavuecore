action "gigavuecore_update_traffic_flows_global_settings" "example" {
  config {
    body = {
      auto_migrate = true
      fabric_resource = {
        mode  = "example"
        scope = "example"
        type  = "example"
      }
      l2_circuit = {
        vlan_ids = "example"
      }
    }
  }
}
