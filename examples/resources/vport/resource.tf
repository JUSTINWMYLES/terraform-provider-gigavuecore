resource "gigavuecore_vport" "example" {
  alias            = "example"
  cluster_id       = "example"
  deferred_binding = true
  fail_over_action = "vport-bypass"
  gs_group         = "example"
  health_state     = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  inline_status      = "up"
  inner_traffic_path = "to-inline-tool"
  metadata_monitoring = {
    action    = "enable"
    exporters = [ "example" ]
  }
  mode               = "none"
  outer_traffic_path = "to-inline-tool"
  sa_apf_profile     = "example"
}
