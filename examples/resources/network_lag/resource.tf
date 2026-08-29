resource "gigavuecore_network_lag" "example" {
  alias            = "example"
  cdp              = true
  cluster_id       = "example"
  comment          = "example"
  forwarding_state = "example"
  health_state     = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  inline_networks          = [ "example" ]
  lacp                     = true
  lfp                      = true
  physical_bypass          = true
  redundancy_control_state = "example"
  redundancy_profile       = "example"
  traffic_path             = "example"
}
