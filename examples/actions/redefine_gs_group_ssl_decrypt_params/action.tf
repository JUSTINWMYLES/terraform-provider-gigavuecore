action "gigavuecore_redefine_gs_group_ssl_decrypt_params" "example" {
  config {
    alias                   = "example"
    cluster_id              = "example"
    decrypt_fail_action     = "example"
    enabled                 = true
    hsm_pkcs11              = null
    hsm_timeout             = 1
    key_cache_timeout       = 1
    key_map                 = "example"
    non_ssl_traffic         = "example"
    pending_session_timeout = 1
    session_timeout         = 1
    tcp_syn_timeout         = 1
    ticket_cache_timeout    = 1
  }
}
