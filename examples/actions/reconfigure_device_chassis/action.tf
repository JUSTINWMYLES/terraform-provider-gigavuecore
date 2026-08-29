action "gigavuecore_reconfigure_device_chassis" "example" {
  config {
    box_id       = 0
    chassis_type = "example"
    cluster_id   = "example"
    gdp          = true
    l2_gre_id    = 0
    leaf_config = {
      mode = "example"
    }
    mode          = "example"
    node_id       = "example"
    serial_number = "example"
    vxlan_id      = 0
  }
}
