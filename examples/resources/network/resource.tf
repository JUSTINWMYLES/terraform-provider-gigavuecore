resource "gigavuecore_network" "example" {
  alias = "example"
  comment = "example"
  forwarding_state = "example"
  health_state = "example"
  health_state_reasons = [{
    message = "example"
    severity = "example"
    traffic_health_state_computation_type = "example"
  }]
  heartbeat = {
    enabled = true
  }
  lfp = true
  physical_bypass = true
  port_a = "example"
  port_b = "example"
  redundancy_control_state = "example"
  redundancy_profile = "example"
  traffic_path = "example"
  type = "example"
}
