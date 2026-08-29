action "gigavuecore_update_inline_ssl_config" "example" {
  config {
    body_cluster_id = "example"
    caching = {
      persistence = {
        enable = true
      }
    }
    cluster_id     = "example"
    dhe_ciphersuit = "example"
    monitor = {
      enable = true
    }
    resumption = {
      client = {
        enable = true
      }
    }
    ssl_versions = {
      connection_reset_action_for_max_version = "example"
      connection_reset_action_for_min_version = "example"
      max_version                             = "example"
      min_version                             = "example"
    }
    start_tls = {
      enable = true
    }
  }
}
