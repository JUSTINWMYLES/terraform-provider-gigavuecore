resource "gigavuecore_serial_tool_group" "example" {
  alias           = "example"
  cluster_id      = "example"
  comment         = "example"
  enabled         = true
  failover_action = "toolBypass"
  health_state    = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  inline_tools        = ["example"]
  per_direction_order = "reverse"
}
