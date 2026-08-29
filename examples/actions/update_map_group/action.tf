action "gigavuecore_update_map_group" "example" {
  config {
    alias = "example"
    map_group_config = {
      cluster_name = "example"
      comment      = "example"
      components = [{
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
      map_group_alias = "example"
    }
    policy_alias       = "example"
    source_alias       = "example"
    source_rules_alias = "example"
  }
}
