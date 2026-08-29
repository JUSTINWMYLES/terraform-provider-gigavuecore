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
}
