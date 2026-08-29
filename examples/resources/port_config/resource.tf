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
  fec              = "example"
  gdp              = true
  header_strip     = "example"
  ingress_vlan_tag = 0
  l2_gre_id        = 0
  licensed         = true
  lock = {
    description  = "example"
    locking_user = "example"
    shared_with  = [ "example" ]
  }
  mpls_advanced = {
    mpls_advanced_opt = [{
      adv_opt = "example"
    }]
  }
  neighbor_discovery = "example"
  port_id            = "example"
  ptp = {
    announce_interval      = 0
    delay_request_interval = 0
    enable                 = true
    local_priority         = 0
    role                   = "example"
    role_alias             = "example"
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
    vlan = 0
  }
  share = {
    tool_share_roles = [ "example" ]
  }
  tag_protocol_id = "example"
  taptx           = "example"
  timestamp = {
    append_ingress = true
    source_id      = 0
    strip_egress   = true
  }
  vxlan_id = 0
}
