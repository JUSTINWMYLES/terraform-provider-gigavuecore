resource "gigavuecore_stack_link" "example" {
  alias        = "example"
  cluster_id   = "example"
  comment      = "example"
  endpoint1    = "example"
  endpoint2    = "example"
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  port_list1 = ["example"]
  port_list2 = ["example"]
  type       = "gigastream"
}
