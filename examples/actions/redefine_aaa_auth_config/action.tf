action "gigavuecore_redefine_aaa_auth_config" "example" {
  config {
    auth_sequence = [ "example" ]
    cluster_id    = "example"
    external_login_mapping = {
      default_local_user = "example"
      user_map_order     = "example"
    }
    non_local_users = {
      hash_username       = true
      track_auth_failures = true
    }
    password_expiration = {
      duration = 0
      enabled  = true
    }
    sshd_max_sessions = 0
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
