resource "gigavuecore_tool_group" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  current_state = {
    inline_tools         = ["example"]
    spare_tool           = "example"
    spare_tool_status    = "down"
    switched_inline_tool = "example"
  }
  enabled           = true
  failover_action   = "toolBypass"
  failover_mode     = "disabled"
  flex_status       = "forwarding"
  flex_traffic_path = "drop"
  hash              = "advanced"
  health_state      = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  inline_tools              = ["example"]
  min_group_size            = 1
  oper_status               = "up"
  release_spare_if_possible = true
  spare_inline_tool         = "example"
}
