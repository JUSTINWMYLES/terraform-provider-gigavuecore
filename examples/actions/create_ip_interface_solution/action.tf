action "gigavuecore_create_ip_interface_solution" "example" {
  config {
    ip_interface_configs = [{
      alias                 = "example"
      applications          = [ "example" ]
      cluster_name          = "example"
      config_status         = "example"
      config_status_reasons = "example"
      gateway               = "example"
      interfaces            = [ "example" ]
      ip_address            = "example"
      ip_mask               = "example"
      managed_status        = "example"
      mtu                   = 0
      ref_count             = 0
    }]
    tags = [{
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
  }
}
