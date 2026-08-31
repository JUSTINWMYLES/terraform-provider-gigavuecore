resource "gigavuecore_network" "example" {
  alias            = "example"
  cluster_id       = "example"
  comment          = "example"
  forwarding_state = "physicalBypass"
  health_state     = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  heartbeat = {
    enabled = true
  }
  lfp                      = true
  physical_bypass          = true
  port_a                   = "example"
  port_b                   = "example"
  redundancy_control_state = "neutral"
  redundancy_profile       = "example"
  traffic_path             = "drop"
  type                     = "protected"
}
