action "gigavuecore_add_map_rule" "example" {
  config {
    alias      = "example"
    bidi       = true
    cluster_id = "example"
    comment    = "example"
    ip_rewrite = {
      dst_ip = "example"
      src_ip = "example"
    }
    matches = [ "example" ]
    rewrite = {
      dst_mac = "example"
      src_mac = "example"
    }
    rule_id   = 1
    rule_type = "example"
    vlan_tag = {
      tag_protocol_id = "0x8100"
      vlan_action     = "add"
      vlan_id         = 0
    }
  }
}
