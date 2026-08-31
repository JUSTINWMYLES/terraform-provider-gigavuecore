action "gigavuecore_update_traffic_flows_draft" "example" {
  config {
    alias           = "example"
    body_alias      = "example"
    comment         = "example"
    deployment_type = "REGULAR"
    enable          = true
    flows           = [ "example" ]
    has_draft       = true
    priority_type   = "LOWEST"
    sources_and_rules = [{
      alias = "example"
      components = [{
        cluster_id = "example"
        components = [{
          ids  = [ "example" ]
          type = "PORT"
        }]
      }]
      inline_traffic_path = "normal"
      inline_traffic_type = "symmetric"
      ip_rewrite = {
        dst_ip = "example"
        src_ip = "example"
      }
      rewrite = {
        dst_mac = "example"
        src_mac = "example"
      }
      rule_matching = "normal"
      rule_type     = "byRule"
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
          matches = [ "example" ]
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
      }
      tags = [{
        tag_key    = "example"
        tag_values = [ "example" ]
      }]
      traffic_type = "control"
      vlan_tag = {
        tag_protocol_id = "0x8100"
        vlan_action     = "add"
        vlan_id         = 0
      }
    }]
    tags = [{
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
  }
}
