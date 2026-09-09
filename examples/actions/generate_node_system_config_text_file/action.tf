action "gigavuecore_generate_node_system_config_text_file" "example" {
  config {
    cluster_id = "example"
    destination = {
      hostname = "example"
      password = "example"
      path     = "example"
      username = "example"
    }
    filename     = "example"
    only_traffic = true
    protocol     = "scp"
    remote       = true
  }
}
