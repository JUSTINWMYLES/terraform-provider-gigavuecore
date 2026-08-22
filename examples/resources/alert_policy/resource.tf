resource "gigavuecore_alert_policy" "example" {
  clear_condition = {
    interval = {
      duration = 1
      unit = "example"
    }
    threshold = {
      severity = "example"
      threshold = 1.0
    }
    type = "example"
  }
  condition = {
    interval = {
      duration = 1
      unit = "example"
    }
    thresholds = [{
      severity = "example"
      threshold = 1.0
    }]
  }
  description = "example"
  enabled = true
  metric = "example"
  policy_name = "example"
  resource_type = "example"
  resources = "example"
}
