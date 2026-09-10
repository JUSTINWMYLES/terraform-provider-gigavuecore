resource "gigavuecore_alert_policy" "example" {
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
  description   = "example"
  enabled       = true
  metric        = "TUNNEL_TRAFFIC_MONITORING"
  policy_name   = "example"
  resource_type = "tunnelMonitoring"
  resources     = "example"
}
