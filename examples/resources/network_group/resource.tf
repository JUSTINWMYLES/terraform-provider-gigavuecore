resource "gigavuecore_network_group" "example" {
  alias        = "example"
  cluster_id   = "example"
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  members = [{
    alias        = "example"
    cluster_name = "example"
    type         = "IN"
  }]
}
