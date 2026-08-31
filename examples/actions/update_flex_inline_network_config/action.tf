action "gigavuecore_update_flex_inline_network_config" "example" {
  config {
    alias                 = "example"
    cluster_id            = "example"
    config_data           = ["example"]
    config_status         = "OPEN"
    config_status_reasons = ["example"]
    config_type           = "FLEXINLINE_MAP"
    health_state          = "green"
    health_state_reasons = [{
      message                               = "example"
      severity                              = "green"
      traffic_health_state_computation_type = "PORT_LOW_UTIL"
    }]
    solution_alias = "example"
  }
}
