action "gigavuecore_update_aaa_auth_config" "example" {
  config {
    auth_sequence = ["local"]
    cluster_id    = "example"
    external_login_mapping = {
      default_local_user = "example"
      user_map_order     = "localOnly"
    }
    non_local_users = {
      hash_username       = true
      track_auth_failures = true
    }
    password_expiration = {
      duration = 1
      enabled  = true
    }
    sshd_max_sessions = 1
    user_lockout = {
      enable_admin_lockout = true
      enable_lockout       = true
      lock_time            = 0
      max_fail             = 0
      track_auth_failures  = true
      unlock_time          = 0
    }
  }
}
