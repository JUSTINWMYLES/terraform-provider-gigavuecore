resource "gigavuecore_port_group" "example" {
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  gigastreams  = ["example"]
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  port_weights = [1]
  ports        = ["example"]
  smart_lb     = true
  tunnel_lb_endpoints = [{
    tunnel_endpoint = "example"
    weight          = 1
  }]
}
