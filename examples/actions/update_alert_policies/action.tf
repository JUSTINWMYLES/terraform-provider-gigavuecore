action "gigavuecore_update_alert_policies" "example" {
  config {
    context = {
      page_no     = 0
      page_size   = 0
      sort        = [ "example" ]
      total_items = 0
    }
    policies = [{
      clear_condition = {
        interval = {
          duration = 0
          unit     = "example"
        }
        threshold = {
          severity  = "example"
          threshold = 1.0
        }
        type = "example"
      }
      condition = {
        interval = {
          duration = 0
          unit     = "example"
        }
        thresholds = [{
          severity  = "example"
          threshold = 1.0
        }]
      }
      description     = "example"
      enabled         = true
      metric          = "example"
      new_policy_name = "example"
      policy_name     = "example"
      resource_type   = "example"
      resources       = "example"
    }]
  }
}
