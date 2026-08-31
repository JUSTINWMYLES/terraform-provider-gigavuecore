action "gigavuecore_redefine_ldap_system_config" "example" {
  config {
    accept_user_roles = true
    cluster_id        = "example"
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
}
