action "gigavuecore_reboot_cluster_nodes" "example" {
  config {
    cluster_ids                = [ "example" ]
    node_ids                   = [ "example" ]
    skip_not_reachable_devices = true
  }
}
