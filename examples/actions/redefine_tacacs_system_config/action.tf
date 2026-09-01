action "gigavuecore_redefine_tacacs_system_config" "example" {
  config {
    accept_user_roles = true
    cluster_id        = "example"
    server_config_defaults = {
      retries    = 0
      secret_key = "example"
      service    = "gigamon"
      timeout    = 1
    }
  }
}
