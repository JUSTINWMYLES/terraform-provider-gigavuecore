list "gigavuecore_alert_policy" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    enabled       = true
    page          = "example"
    policy_name   = "example"
    resource_type = "example"
    sort          = "example"
  }
}
