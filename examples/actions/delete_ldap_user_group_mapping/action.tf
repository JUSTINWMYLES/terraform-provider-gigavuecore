action "gigavuecore_delete_ldap_user_group_mapping" "example" {
  config {
    cluster_id = "example"
    remote_map_table = [{
      local_account_name = "example"
      remote_base_dn     = "example"
    }]
  }
}
