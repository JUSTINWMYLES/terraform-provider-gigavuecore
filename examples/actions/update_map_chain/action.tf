action "gigavuecore_update_map_chain" "example" {
  config {
    alias        = "example"
    policy_alias = "example"
    priority_configs = [{
      cluster_name        = "example"
      collector_map_alias = "example"
      map_chain_id        = "example"
      ordered_components = [{
        alias        = "example"
        health_state = "example"
        health_state_reasons = [{
          message                               = "example"
          severity                              = "example"
          traffic_health_state_computation_type = "example"
        }]
        source_alias         = "example"
        traffic_health_state = "example"
        traffic_health_state_reasons = [{
          message                               = "example"
          severity                              = "example"
          traffic_health_state_computation_type = "example"
        }]
        type = "example"
      }]
      src_ports_as_id = "example"
    }]
    source_alias       = "example"
    source_rules_alias = "example"
  }
}
