action "gigavuecore_update_alert_policies" "example" {
  config {
    context = {
      page_no     = 1
      page_size   = 0
      sort        = ["example"]
      total_items = 0
    }
    policies = [{
      clear_condition = {
        interval = {
          duration = 0
          unit     = "MINUTES"
        }
        threshold = {
          severity  = "Info"
          threshold = 1.0
        }
        type = "TIME_BASED"
      }
      condition = {
        interval = {
          duration = 0
          unit     = "MINUTES"
        }
        thresholds = [{
          severity  = "Info"
          threshold = 1.0
        }]
      }
      description     = "example"
      enabled         = true
      metric          = "TUNNEL_TRAFFIC_MONITORING"
      new_policy_name = "example"
      policy_name     = "example"
      resource_type   = "tunnelMonitoring"
      resources       = "example"
    }]
  }
}
