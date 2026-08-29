action "gigavuecore_restore_node_system_config_text_file" "example" {
  config {
    clear_config = true
    cluster_id   = "example"
    destination = {
      hostname = "example"
      password = "example"
      path     = "example"
      username = "example"
    }
    fail_continue = true
    filename      = "example"
    protocol      = "example"
    remote        = true
  }
}
