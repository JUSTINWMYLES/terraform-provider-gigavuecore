resource "gigavuecore_network_lag" "example" {
  alias            = "example"
  cdp              = true
  cluster_id       = "example"
  comment          = "example"
  forwarding_state = "physicalBypass"
  health_state     = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  inline_networks          = ["example"]
  lacp                     = true
  lfp                      = true
  physical_bypass          = true
  redundancy_control_state = "neutral"
  redundancy_profile       = "example"
  traffic_path             = "drop"
}
