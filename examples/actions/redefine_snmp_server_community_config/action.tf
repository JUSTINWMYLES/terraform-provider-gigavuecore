action "gigavuecore_redefine_snmp_server_community_config" "example" {
  config {
    cluster_id               = "example"
    community_strings        = [ "example" ]
    enable_community_auth    = true
    enable_community_auth_v1 = true
    enable_multi_community   = true
  }
}
