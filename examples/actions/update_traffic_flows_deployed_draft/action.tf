action "gigavuecore_update_traffic_flows_deployed_draft" "example" {
  config {
    alias           = "example"
    body_alias      = "example"
    comment         = "example"
    deployment_type = "example"
    enable          = true
    flows           = [ "example" ]
    has_draft       = true
    priority_type   = "example"
    sources_and_rules = [{
      alias = "example"
      components = [{
        cluster_id = "example"
        components = [{
          ids  = [ "example" ]
          type = "example"
        }]
      }]
      inline_traffic_path = "example"
      inline_traffic_type = "example"
      ip_rewrite = {
        dst_ip = "example"
        src_ip = "example"
      }
      rewrite = {
        dst_mac = "example"
        src_mac = "example"
      }
      rule_matching = "example"
      rule_type     = "example"
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
      tags = [{
        tag_key    = "example"
        tag_values = [ "example" ]
      }]
      traffic_type = "example"
      vlan_tag = {
        tag_protocol_id = "example"
        vlan_action     = "example"
        vlan_id         = 0
      }
    }]
    tags = [{
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
  }
}
