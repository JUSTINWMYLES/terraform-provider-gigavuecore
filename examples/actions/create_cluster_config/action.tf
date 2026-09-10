action "gigavuecore_create_cluster_config" "example" {
  config {
    cluster_member_specs = [{
      box_id       = 0
      cluster_intf = "example"
      cluster_port = 0
      leader_discovery = {
        auto_discovery            = true
        cluster_formation_timeout = 0
        primary_ip                = "example"
        primary_port              = 0
        secondary_ip              = "example"
        secondary_port            = 0
        static_discovery_timeout  = 10
      }
      leader_preference = 0
      master_discovery = {
        auto_discovery            = true
        cluster_formation_timeout = 0
        primary_ip                = "example"
        primary_port              = 0
        secondary_ip              = "example"
        secondary_port            = 0
        static_discovery_timeout  = 10
      }
      master_preference = 0
      mgmt_address      = "example"
      vip_intf          = "example"
    }]
    cluster_params = {
      cluster_id           = "example"
      cluster_vip          = "example"
      cluster_vip_mask_len = 0
      ip_protocol          = "ipv4"
      shared_secret        = "example"
      stacking_mode        = "example"
    }
    cluster_type           = "example"
    seed_node_mgmt_address = "example"
  }
}
