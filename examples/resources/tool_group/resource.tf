resource "gigavuecore_tool_group" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  current_state = {
    inline_tools         = [ "example" ]
    spare_tool           = "example"
    spare_tool_status    = "example"
    switched_inline_tool = "example"
  }
  enabled           = true
  failover_action   = "example"
  failover_mode     = "example"
  flex_status       = "example"
  flex_traffic_path = "example"
  hash              = "example"
  health_state      = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  inline_tools              = [ "example" ]
  min_group_size            = 0
  oper_status               = "example"
  release_spare_if_possible = true
  spare_inline_tool         = "example"
}
