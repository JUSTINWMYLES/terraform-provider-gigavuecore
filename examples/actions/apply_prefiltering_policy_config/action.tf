action "gigavuecore_apply_prefiltering_policy_config" "example" {
  config {
    id   = "example"
    name = "example"
    rules = [{
      action    = "pass"
      direction = "bidi"
      filters   = ["example"]
      priority  = "example"
      rule_name = "example"
    }]
  }
}
