action "gigavuecore_upgrade_cluster_nodes_image" "example" {
  config {
    async                      = true
    cluster_ids                = [ "example" ]
    image_file_specs           = "example"
    image_server               = "example"
    node_ids                   = [ "example" ]
    reboot                     = true
    skip_not_reachable_devices = true
  }
}
