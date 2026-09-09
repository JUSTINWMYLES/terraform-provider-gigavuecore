resource "gigavuecore_port_pair" "example" {
  alias           = "example"
  cluster_id      = "example"
  comment         = "example"
  config_mismatch = true
  health_state    = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  lfp_enabled = true
  port1       = "example"
  port2       = "example"
}
