resource "gigavuecore_map_template" "example" {
  alias      = "example"
  cluster_id = "example"
  comment    = "example"
  roles = {
    editors = [ "example" ]
    owners  = [ "example" ]
    viewers = [ "example" ]
  }
  rules = {
    drop_rules = [{
      bidi    = true
      comment = "example"
      ip_rewrite = {
        dst_ip = "example"
        src_ip = "example"
      }
      matches = [ "example" ]
      rewrite = {
        dst_mac = "example"
        src_mac = "example"
      }
      rule_id = 0
      vlan_tag = {
        tag_protocol_id = "example"
        vlan_action     = "example"
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
      matches = [ "example" ]
      rewrite = {
        dst_mac = "example"
        src_mac = "example"
      }
      rule_id = 0
      vlan_tag = {
        tag_protocol_id = "example"
        vlan_action     = "example"
        vlan_id         = 0
      }
    }]
  }
}
