action "gigavuecore_update_remote_auth_system_config" "example" {
  config {
    cluster_id = "example"
    ldap_config = {
      accept_user_roles = true
      map_enable        = true
      referrals         = true
      remote_map_table = [{
        local_account_name = "example"
        remote_base_dn     = "example"
      }]
      server_config_defaults = {
        base_dn         = "example"
        bind_dn         = "example"
        bind_pwd        = "example"
        bind_timeout    = 1
        group_attribute = "example"
        group_dn        = "example"
        login_attribute = "example"
        port            = 1
        search_scope    = "oneLevel"
        search_timeout  = 1
        ssl = {
          ca_list     = "none"
          cert_verify = true
          mode        = "none"
          server_port = 0
        }
        version = "v2"
      }
    }
    radius_config = {
      accept_user_roles = true
      server_config_defaults = {
        retries    = 0
        secret_key = "example"
        timeout    = 1
      }
    }
    tacacs_config = {
      accept_user_roles = true
      server_config_defaults = {
        retries    = 0
        secret_key = "example"
        service    = "gigamon"
        timeout    = 1
      }
    }
  }
}
