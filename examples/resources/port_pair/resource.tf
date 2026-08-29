resource "gigavuecore_port_pair" "example" {
  alias           = "example"
  cluster_id      = "example"
  comment         = "example"
  config_mismatch = true
  health_state    = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  lfp_enabled = true
  port1       = "example"
  port2       = "example"
}
