action "gigavuecore_enable_alarm_auto_suppression" "example" {
  config {
    auto_suppression_rules = [{
      enable = true
      type   = "example"
    }]
    enable = true
  }
}
