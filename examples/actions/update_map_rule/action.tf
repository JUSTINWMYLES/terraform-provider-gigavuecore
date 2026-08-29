action "gigavuecore_update_map_rule" "example" {
  config {
    alias        = "example"
    bidi         = true
    body_rule_id = 0
    cluster_id   = "example"
    comment      = "example"
    ip_rewrite = {
      dst_ip = "example"
      src_ip = "example"
    }
    matches = [ "example" ]
    rewrite = {
      dst_mac = "example"
      src_mac = "example"
    }
    rule_id   = "example"
    rule_type = "example"
    vlan_tag = {
      tag_protocol_id = "example"
      vlan_action     = "example"
      vlan_id         = 0
    }
  }
}
