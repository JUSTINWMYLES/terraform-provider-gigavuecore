action "gigavuecore_fetch_node_system_config_file" "example" {
  config {
    cluster_id = "example"
    protocol   = "scp"
    source = {
      hostname = "example"
      password = "example"
      path     = "example"
      username = "example"
    }
  }
}
