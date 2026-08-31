resource "gigavuecore_port_config" "example" {
  access_roles = {
    level1 = [ "example" ]
    level2 = [ "example" ]
    level3 = [ "example" ]
    level4 = [ "example" ]
  }
  alarm_thresholds = {
    alarm_buffer_threshold_rx = 0
    alarm_buffer_threshold_tx = 0
    alarm_threshold           = 0
    alarm_threshold_low       = 0
  }
  cluster_id       = "example"
  fec              = "cl91"
  gdp              = true
  header_strip     = "none"
  ingress_vlan_tag = 0
  l2_gre_id        = 1
  licensed         = true
  lock = {
    description  = "example"
    locking_user = "example"
    shared_with  = [ "example" ]
  }
  mpls_advanced = {
    mpls_advanced_opt = [{
      adv_opt = "ip4-wo-options"
    }]
  }
  neighbor_discovery = "none"
  port_id            = "example"
  ptp = {
    announce_interval      = 0
    delay_request_interval = 0
    enable                 = true
    local_priority         = 1
    role                   = "standard"
    role_alias             = "standard"
    sync_interval          = 0
    timestamp = {
      egress = {
        insert    = true
        source_id = 0
      }
      ingress = {
        insert    = true
        source_id = 0
      }
    }
    vlan = 4080
  }
  share = {
    tool_share_roles = [ "example" ]
  }
  tag_protocol_id = "0x8100"
  taptx           = "active"
  timestamp = {
    append_ingress = true
    source_id      = 1
    strip_egress   = true
  }
  vxlan_id = 1
}
