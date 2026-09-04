action "gigavuecore_enable_alarm_auto_suppression" "example" {
  config {
    auto_suppression_rules = [{
      enable = true
      type   = "image_upgrade"
    }]
    enable = true
  }
}
