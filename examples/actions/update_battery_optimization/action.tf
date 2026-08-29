action "gigavuecore_update_battery_optimization" "example" {
  config {
    cluster_id = "example"
    cpu_hibernation = {
      enabled         = true
      sleep_in_mins   = 0
      threshold_level = 0
    }
    monitor_port = [{
      enabled       = true
      level         = "example"
      port_group_id = "example"
    }]
    unused_port = [{
      enabled       = true
      level         = "example"
      port_group_id = "example"
    }]
  }
}
