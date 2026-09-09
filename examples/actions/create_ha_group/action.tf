action "gigavuecore_create_ha_group" "example" {
  config {
    fm_ha_tunnel = {
      tunnel_auth_mode = "PSK"
    }
    name = "example"
    nodes = {
      cluster_ip_address    = "example"
      entity_id             = "example"
      hostname              = "example"
      idp_meta_data_url     = "example"
      management_ip_address = "example"
      password              = "example"
      public_ip_address     = "example"
      reachable             = true
      seed_node             = true
      username              = "example"
    }
  }
}
