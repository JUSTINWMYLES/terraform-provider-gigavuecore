action "gigavuecore_cluster_config_add_member" "example" {
  config {
    bulk       = true
    cluster_id = "example"
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
        static_discovery_timeout  = 0
      }
      leader_preference = 0
      master_discovery = {
        auto_discovery            = true
        cluster_formation_timeout = 0
        primary_ip                = "example"
        primary_port              = 0
        secondary_ip              = "example"
        secondary_port            = 0
        static_discovery_timeout  = 0
      }
      master_preference = 0
      mgmt_address      = "example"
      vip_intf          = "example"
    }]
  }
}
