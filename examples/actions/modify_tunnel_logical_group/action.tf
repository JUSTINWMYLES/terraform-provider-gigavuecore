action "gigavuecore_modify_tunnel_logical_group" "example" {
  config {
    tunnel_logical_groups = [{
      alert_policy_name = "example"
      alias             = "example"
      custom_alias      = "example"
      decap_tunnel_infos = [{
        tunnel_operation_type = "Encap"
        tunnel_type           = "Circuit"
      }]
      encap_tunnel_infos = [{
        tunnel_operation_type = "Encap"
        tunnel_type           = "Circuit"
      }]
      health_state = "green"
      health_state_reasons = [{
        message                               = "example"
        severity                              = "green"
        traffic_health_state_computation_type = "PORT_LOW_UTIL"
      }]
      is_active   = true
      key         = "example"
      state       = "Complete"
      tunnel_id   = 0
      tunnel_type = "Embedded Circuit"
    }]
  }
}
