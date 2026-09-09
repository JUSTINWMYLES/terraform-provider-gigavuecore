action "gigavuecore_update_device_port" "example" {
  config {
    admin_status  = "up"
    alias         = "example"
    auto_neg      = true
    body_port_id  = "example"
    breakout_mode = "none"
    cable_length  = "example"
    cluster_id    = "example"
    comment       = "example"
    config_speed  = "10M"
    duplex        = "full"
    force_link_up = true
    mtu           = 1500
    port_id       = "example"
    port_type     = "network"
    ude = {
      enabled = true
    }
  }
}
