action "gigavuecore_modify_tunnel_logical_group" "example" {
  config {
    tunnel_logical_groups = [{
      alert_policy_name = "example"
      alias             = "example"
      custom_alias      = "example"
      decap_tunnel_infos = [{
        tunnel_operation_type = "example"
        tunnel_type           = "example"
      }]
      encap_tunnel_infos = [{
        tunnel_operation_type = "example"
        tunnel_type           = "example"
      }]
      health_state = "example"
      health_state_reasons = [{
        message                               = "example"
        severity                              = "example"
        traffic_health_state_computation_type = "example"
      }]
      is_active   = true
      key         = "example"
      state       = "example"
      tunnel_id   = "example"
      tunnel_type = "example"
    }]
  }
}
