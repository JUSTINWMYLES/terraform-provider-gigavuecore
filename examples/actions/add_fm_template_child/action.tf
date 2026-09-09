action "gigavuecore_add_fm_template_child" "example" {
  config {
    body_child_type  = "ldapServers"
    body_config_type = "LDAP_SERVERS_TEMPLATE"
    child_type       = "example"
    config = {
      ldap_servers = [{
        order          = "example"
        server_address = "example"
      }]
      remote_map_table = [{
        local_account_name = "example"
        remote_base_dn     = "example"
      }]
    }
    config_level       = "GLOBAL"
    config_level_value = ["example"]
    config_type        = "example"
  }
}
