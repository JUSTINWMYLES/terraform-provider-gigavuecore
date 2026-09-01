action "gigavuecore_add_source_rules" "example" {
  config {
    drop_rules = [{
      bidi    = true
      comment = "example"
      ip_rewrite = {
        dst_ip = "example"
        src_ip = "example"
      }
      matches = ["example"]
      rewrite = {
        dst_mac = "example"
        src_mac = "example"
      }
      rule_id = 1
      vlan_tag = {
        tag_protocol_id = "0x8100"
        vlan_action     = "add"
        vlan_id         = 0
      }
    }]
    pass_rules = [{
      bidi    = true
      comment = "example"
      ip_rewrite = {
        dst_ip = "example"
        src_ip = "example"
      }
      matches = ["example"]
      rewrite = {
        dst_mac = "example"
        src_mac = "example"
      }
      rule_id = 1
      vlan_tag = {
        tag_protocol_id = "0x8100"
        vlan_action     = "add"
        vlan_id         = 0
      }
    }]
    policy_alias       = "example"
    source_rules_alias = "example"
  }
}
