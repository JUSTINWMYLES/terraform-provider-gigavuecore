action "gigavuecore_upgrade_cluster_nodes_image" "example" {
  config {
    async       = true
    cluster_ids = ["example"]
    image_file_specs = [{
      device_model = "HD8"
      file_path    = "example"
      file_type    = "image"
      target_slot  = "2"
      update_uboot = true
    }]
    image_server               = "example"
    node_ids                   = ["example"]
    reboot                     = true
    skip_not_reachable_devices = true
  }
}
