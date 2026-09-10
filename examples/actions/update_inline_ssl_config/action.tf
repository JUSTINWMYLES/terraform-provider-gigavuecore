action "gigavuecore_update_inline_ssl_config" "example" {
  config {
    body_cluster_id = "example"
    caching = {
      persistence = {
        enable = true
      }
    }
    cluster_id     = "example"
    dhe_ciphersuit = "disable"
    monitor = {
      enable = true
    }
    resumption = {
      client = {
        enable = true
      }
    }
    ssl_versions = {
      connection_reset_action_for_max_version = "no-decrypt"
      connection_reset_action_for_min_version = "no-decrypt"
      max_version                             = "sslv3"
      min_version                             = "sslv3"
    }
    start_tls = {
      enable = true
    }
  }
}
