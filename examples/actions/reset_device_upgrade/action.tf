action "gigavuecore_reset_device_upgrade" "example" {
  config {
    device_upgrade_reset_specs = [{
      cluster_name = "example"
      task_id      = "example"
    }]
  }
}
