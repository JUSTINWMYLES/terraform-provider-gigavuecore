resource "gigavuecore_vport" "example" {
  alias            = "example"
  deferred_binding = true
  fail_over_action = "example"
  gs_group         = "example"
  health_state     = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  inline_status      = "example"
  inner_traffic_path = "example"
  metadata_monitoring = {
    action    = "example"
    exporters = [ "example" ]
  }
  mode               = "example"
  outer_traffic_path = "example"
  sa_apf_profile     = "example"
}
