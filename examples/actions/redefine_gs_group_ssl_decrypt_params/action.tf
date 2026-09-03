action "gigavuecore_redefine_gs_group_ssl_decrypt_params" "example" {
  config {
    alias               = "example"
    cluster_id          = "example"
    decrypt_fail_action = "drop"
    enabled             = true
    hsm_pkcs11 = {
      debug_level    = 0
      dynamic_object = true
      load_sharing   = true
    }
    hsm_timeout             = 2
    key_cache_timeout       = 1
    key_map                 = "example"
    non_ssl_traffic         = "drop"
    pending_session_timeout = 30
    session_timeout         = 30
    tcp_syn_timeout         = 20
    ticket_cache_timeout    = 1
  }
}
