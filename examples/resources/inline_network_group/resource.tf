resource "gigavuecore_inline_network_group" "example" {
  alias        = "example"
  bundled      = true
  cluster_id   = "example"
  comment      = "example"
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  inline_networks = [ "example" ]
}
