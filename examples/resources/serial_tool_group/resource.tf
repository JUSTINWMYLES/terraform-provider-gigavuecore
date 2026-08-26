resource "gigavuecore_serial_tool_group" "example" {
  alias           = "example"
  cluster_id      = "example"
  comment         = "example"
  enabled         = true
  failover_action = "example"
  health_state    = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  inline_tools        = [ "example" ]
  per_direction_order = "example"
}
