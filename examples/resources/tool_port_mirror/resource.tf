resource "gigavuecore_tool_port_mirror" "example" {
  alias        = "example"
  comment      = "example"
  dst_ports    = [ "example" ]
  health_state = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  src_ports = [ "example" ]
}
