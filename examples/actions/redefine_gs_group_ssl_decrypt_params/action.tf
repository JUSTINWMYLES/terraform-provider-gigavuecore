action "gigavuecore_redefine_gs_group_ssl_decrypt_params" "example" {
  config {
    alias               = "example"
    cluster_id          = "example"
    decrypt_fail_action = "example"
    enabled             = true
    hsm_pkcs11 = {
      debug_level    = 0
      dynamic_object = true
      load_sharing   = true
    }
    hsm_timeout             = 0
    key_cache_timeout       = 0
    key_map                 = "example"
    non_ssl_traffic         = "example"
    pending_session_timeout = 0
    session_timeout         = 0
    tcp_syn_timeout         = 0
    ticket_cache_timeout    = 0
  }
}
