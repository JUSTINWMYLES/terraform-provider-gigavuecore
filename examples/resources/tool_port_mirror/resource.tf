resource "gigavuecore_tool_port_mirror" "example" {
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  dst_ports    = [ "example" ]
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  src_ports = [ "example" ]
}
