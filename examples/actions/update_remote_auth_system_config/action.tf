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
        bind_timeout    = 0
        group_attribute = "example"
        group_dn        = "example"
        login_attribute = "example"
        port            = 0
        search_scope    = "example"
        search_timeout  = 0
        ssl = {
          ca_list     = "example"
          cert_verify = true
          mode        = "example"
          server_port = 0
        }
        version = "example"
      }
    }
    radius_config = {
      accept_user_roles = true
      server_config_defaults = {
        retries    = 0
        secret_key = "example"
        timeout    = 0
      }
    }
    tacacs_config = {
      accept_user_roles = true
      server_config_defaults = {
        retries    = 0
        secret_key = "example"
        service    = "example"
        timeout    = 0
      }
    }
  }
}
