resource "gigavuecore_port_config" "example" {
  access_roles = {
    level1 = [ "example" ]
    level2 = [ "example" ]
    level3 = [ "example" ]
    level4 = [ "example" ]
  }
  alarm_thresholds = {
    alarm_buffer_threshold_rx = 1
    alarm_buffer_threshold_tx = 1
    alarm_threshold = 1
    alarm_threshold_low = 1
  }
  fec = "example"
  gdp = true
  header_strip = "example"
  id = "example"
  ingress_vlan_tag = 1
  l2_gre_id = 1
  licensed = true
  lock = {
    description = "example"
    locking_user = "example"
    shared_with = [ "example" ]
  }
  mpls_advanced = {
    mpls_advanced_opt = [{
      adv_opt = "example"
    }]
  }
  neighbor_discovery = "example"
  port_id = "example"
  ptp = {
    announce_interval = 1
    delay_request_interval = 1
    enable = true
    local_priority = 1
    role = "example"
    role_alias = "example"
    sync_interval = 1
    timestamp = {
      egress = {
        insert = true
        source_id = 1
      }
      ingress = {
        insert = true
        source_id = 1
      }
    }
    vlan = 1
  }
  share = {
    tool_share_roles = [ "example" ]
  }
  tag_protocol_id = "example"
  taptx = "example"
  timestamp = {
    append_ingress = true
    source_id = 1
    strip_egress = true
  }
  vxlan_id = 1
}
