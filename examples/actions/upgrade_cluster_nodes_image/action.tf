action "gigavuecore_upgrade_cluster_nodes_image" "example" {
  config {
    async       = true
    cluster_ids = [ "example" ]
    image_file_specs = [{
      device_model = "example"
      file_path    = "example"
      file_type    = "example"
      target_slot  = "example"
      update_uboot = true
    }]
    image_server               = "example"
    node_ids                   = [ "example" ]
    reboot                     = true
    skip_not_reachable_devices = true
  }
}
