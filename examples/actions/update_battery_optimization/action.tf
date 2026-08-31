action "gigavuecore_update_battery_optimization" "example" {
  config {
    cluster_id = "example"
    cpu_hibernation = {
      enabled         = true
      sleep_in_mins   = 6
      threshold_level = 25
    }
    monitor_port = [{
      enabled       = true
      level         = "alarm"
      port_group_id = "example"
    }]
    unused_port = [{
      enabled       = true
      level         = "alarm"
      port_group_id = "example"
    }]
  }
}
