resource "gigavuecore_network_group" "example" {
  alias        = "example"
  cluster_id   = "example"
  health_state = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  members = [{
    alias        = "example"
    cluster_name = "example"
    type         = "example"
  }]
}
