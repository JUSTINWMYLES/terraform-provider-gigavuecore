resource "gigavuecore_inline_network_group" "example" {
  alias = "example"
  bundled = true
  comment = "example"
  health_state = "example"
  health_state_reasons = [{
    message = "example"
    severity = "example"
    traffic_health_state_computation_type = "example"
  }]
  inline_networks = [ "example" ]
}
