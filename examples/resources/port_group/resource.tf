resource "gigavuecore_port_group" "example" {
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  gigastreams  = [ "example" ]
  health_state = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  port_weights = [ 0 ]
  ports        = [ "example" ]
  smart_lb     = true
  tunnel_lb_endpoints = [{
    tunnel_endpoint = "example"
    weight          = 0
  }]
}
