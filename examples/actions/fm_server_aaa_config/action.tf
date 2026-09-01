action "gigavuecore_fm_server_aaa_config" "example" {
  config {
    accounting_method      = "example"
    auth_method            = "local"
    default_login_attempts = 3
    default_user_group     = "example"
    user_lock_enabled      = true
  }
}
