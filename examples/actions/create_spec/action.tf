action "gigavuecore_create_spec" "example" {
  config {
    activate_stage       = true
    cluster_ids          = ["example"]
    config_backup        = true
    exec_cluster_counter = 0
    fetch_stage          = true
    image_file_specs = [{
      device_model           = "HC2"
      device_model_qualifier = "HDCCV1"
      file_path              = "example"
      file_type              = "example"
      target_slot            = "example"
      update_uboot           = "example"
    }]
    image_server               = "example"
    install_stage              = true
    node_ids                   = ["example"]
    post_check_stage           = true
    reboot                     = true
    skip_not_reachable_devices = true
    tags = [{
      tag_key    = "example"
      tag_values = ["example"]
    }]
    task_id       = "example"
    task_name     = "example"
    upgrade_state = "ACTIVE"
    version       = "example"
  }
}
