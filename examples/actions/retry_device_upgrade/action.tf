action "gigavuecore_retry_device_upgrade" "example" {
  config {
    device_upgrade_retry_specs = [{
      cluster_name = "example"
      option       = "example"
      task_id      = "example"
    }]
  }
}
