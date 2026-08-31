resource "gigavuecore_flex_inline" "example" {
  alias = "example"
  cluster_configs = [{
    cluster_id = "example"
    export_criteria = {
      lsb = 0
    }
    export_type = "ipBased"
    ib_pathway  = "example"
    source = {
      alias = "example"
      type  = "IN"
    }
  }]
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  keep_on_device = true
  resilient_config = {
    side_a = "source"
    side_b = "source"
  }
  target_traffic_path = "drop"
}
