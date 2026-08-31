action "gigavuecore_resume_device_upgrade" "example" {
  config {
    device_upgrade_resume_specs = [{
      activate_stage   = true
      cluster_name     = "example"
      config_backup    = true
      device_ip        = "example"
      fetch_stage      = true
      install_stage    = true
      post_check_stage = true
      task_id          = "example"
    }]
  }
}
