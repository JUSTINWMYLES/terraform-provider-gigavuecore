action "gigavuecore_update_aaa_auth_config" "example" {
  config {
    auth_sequence          = [ "example" ]
    cluster_id             = "example"
    external_login_mapping = null
    non_local_users        = null
    password_expiration    = null
    sshd_max_sessions      = 1
    user_lockout           = null
  }
}
