resource "gigavuecore_tunnel_application" "example" {
  alias                 = "example"
  cluster_id            = "example"
  config_status         = "example"
  config_status_reasons = [ "example" ]
  decap = [{
    additonalgsops = {
      apf = {
        enabled = "enabled"
      }
      dedup = {
        enabled = "enabled"
      }
      diameter_whitelist = {
        enabled = "enabled"
      }
      flow_filter = {
        type = "gtp"
      }
      flow_sampling = {
        type = "ip"
      }
      gseries_header_add = {
        types = [ "srcid" ]
      }
      gseries_header_remove = {
        enabled = "enabled"
      }
      gseries_load_balance = {
        fixed_offset = {
          hash   = "checksum"
          length = 1
          offset = 0
        }
        variable_offset = {
          end_delim   = "example"
          hash        = "checksum"
          start_delim = "example"
          start_field = "example"
        }
      }
      gseries_pattern_match = {
        fixed_offset = {
          length = 1
          offset = 0
        }
        variable_offset = {
          end_delim   = "example"
          start_delim = "example"
        }
      }
      gtp_whitelist = {
        enabled = "enabled"
      }
      header_add = {
        vlan = 0
      }
      header_remove = {
        ah1                = "none"
        ah2                = "none"
        custom_len         = 1
        erspan_flow_id     = 0
        fp_dst_switch_id   = 0
        fp_src_switch_id   = 0
        header_count       = 1
        offset             = "start"
        offset_range_value = 0
        protocol           = "gtp"
        timestamp_format   = "gigasmart"
        vlan_header        = "all"
        vxlan_id           = 0
      }
      icap = {
        icap_profile = "example"
      }
      inline_ssl = {
        inline_ssl_profile = "example"
      }
      load_balance = {
        enhanced = {
          elb_alias = "example"
        }
        stateful = {
          app_type               = "gtp"
          diameter_key_hash_type = "sessionId"
          diameter_key_multi_hash_type = [{
            avp_codevalue = 0
            key           = "sessionId"
          }]
          gtp_key_hash_type = "imsi"
          lb_type           = "leastBw"
          sip_key_hash_type = "callerId"
        }
        stateless = {
          field_location = "inner"
          hash_fields    = "ipOnly"
        }
      }
      masking = {
        content_type = "message_cpim"
        length       = 1
        offset       = 0
        pattern      = "a1"
        protocol     = "none"
      }
      metadata_export = {
        cache = "example"
      }
      netflow = {
        enabled = "enabled"
      }
      sa_apf = {
        enabled = "enabled"
      }
      sip_whitelist = {
        enabled = "enabled"
      }
      slicing = {
        enhanced = "example"
        offset   = 4
        protocol = "none"
      }
      ssl_decrypt = {
        in_port  = 0
        out_port = 0
      }
      trailer_add = {
        types = [ "crc" ]
      }
      trailer_remove = {
        enabled = "enabled"
      }
      tunnel_decap = {
        custom = {
          port_dst = 0
          port_src = 0
        }
        erspan_flow_id = 0
        gmip_port      = 0
        l2_gre_key     = 0
        tls_pcapng = {
          decap_key = "example"
          listener  = "example"
        }
        type = "gmip"
        vxlan = {
          port_dst = 1
          port_src = 0
          vni      = 0
        }
      }
      tunnel_encap = {
        gmip_config = {
          dscp       = 0
          dst_ip     = "example"
          dst_port   = 0
          flow_label = 0
          prec       = 0
          src_port   = 0
          ttl        = 1
        }
        l2_gre_config = {
          dscp          = 0
          dst_ip        = "example"
          flow_label    = 0
          key           = 0
          pg_dst        = "example"
          prec          = 0
          session_field = "fiveTupleIpv4"
          session_pos   = "inner"
          ttl           = 1
        }
        tls_pcapng = {
          exporter       = "example"
          exporter_group = "example"
        }
        type = "gmip"
        vxlan_config = {
          dscp     = 0
          dst_ip   = "example"
          dst_port = 4789
          src_port = 0
          ttl      = 1
          vni      = 1
        }
      }
    }
    application_ports = [ 0 ]
    destination_port  = [ "example" ]
    gsop_alias        = "example"
    listener_alias    = "example"
    map_alias         = "example"
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
  }]
  description = "example"
  dscp        = 0
  encap = {
    additonalgsops = {
      apf = {
        enabled = "enabled"
      }
      dedup = {
        enabled = "enabled"
      }
      diameter_whitelist = {
        enabled = "enabled"
      }
      flow_filter = {
        type = "gtp"
      }
      flow_sampling = {
        type = "ip"
      }
      gseries_header_add = {
        types = [ "srcid" ]
      }
      gseries_header_remove = {
        enabled = "enabled"
      }
      gseries_load_balance = {
        fixed_offset = {
          hash   = "checksum"
          length = 1
          offset = 0
        }
        variable_offset = {
          end_delim   = "example"
          hash        = "checksum"
          start_delim = "example"
          start_field = "example"
        }
      }
      gseries_pattern_match = {
        fixed_offset = {
          length = 1
          offset = 0
        }
        variable_offset = {
          end_delim   = "example"
          start_delim = "example"
        }
      }
      gtp_whitelist = {
        enabled = "enabled"
      }
      header_add = {
        vlan = 0
      }
      header_remove = {
        ah1                = "none"
        ah2                = "none"
        custom_len         = 1
        erspan_flow_id     = 0
        fp_dst_switch_id   = 0
        fp_src_switch_id   = 0
        header_count       = 1
        offset             = "start"
        offset_range_value = 0
        protocol           = "gtp"
        timestamp_format   = "gigasmart"
        vlan_header        = "all"
        vxlan_id           = 0
      }
      icap = {
        icap_profile = "example"
      }
      inline_ssl = {
        inline_ssl_profile = "example"
      }
      load_balance = {
        enhanced = {
          elb_alias = "example"
        }
        stateful = {
          app_type               = "gtp"
          diameter_key_hash_type = "sessionId"
          diameter_key_multi_hash_type = [{
            avp_codevalue = 0
            key           = "sessionId"
          }]
          gtp_key_hash_type = "imsi"
          lb_type           = "leastBw"
          sip_key_hash_type = "callerId"
        }
        stateless = {
          field_location = "inner"
          hash_fields    = "ipOnly"
        }
      }
      masking = {
        content_type = "message_cpim"
        length       = 1
        offset       = 0
        pattern      = "a1"
        protocol     = "none"
      }
      metadata_export = {
        cache = "example"
      }
      netflow = {
        enabled = "enabled"
      }
      sa_apf = {
        enabled = "enabled"
      }
      sip_whitelist = {
        enabled = "enabled"
      }
      slicing = {
        enhanced = "example"
        offset   = 4
        protocol = "none"
      }
      ssl_decrypt = {
        in_port  = 0
        out_port = 0
      }
      trailer_add = {
        types = [ "crc" ]
      }
      trailer_remove = {
        enabled = "enabled"
      }
      tunnel_decap = {
        custom = {
          port_dst = 0
          port_src = 0
        }
        erspan_flow_id = 0
        gmip_port      = 0
        l2_gre_key     = 0
        tls_pcapng = {
          decap_key = "example"
          listener  = "example"
        }
        type = "gmip"
        vxlan = {
          port_dst = 1
          port_src = 0
          vni      = 0
        }
      }
      tunnel_encap = {
        gmip_config = {
          dscp       = 0
          dst_ip     = "example"
          dst_port   = 0
          flow_label = 0
          prec       = 0
          src_port   = 0
          ttl        = 1
        }
        l2_gre_config = {
          dscp          = 0
          dst_ip        = "example"
          flow_label    = 0
          key           = 0
          pg_dst        = "example"
          prec          = 0
          session_field = "fiveTupleIpv4"
          session_pos   = "inner"
          ttl           = 1
        }
        tls_pcapng = {
          exporter       = "example"
          exporter_group = "example"
        }
        type = "gmip"
        vxlan_config = {
          dscp     = 0
          dst_ip   = "example"
          dst_port = 4789
          src_port = 0
          ttl      = 1
          vni      = 1
        }
      }
    }
    export_configs = [{
      exporter_alias          = "example"
      remote_application_port = 0
      remote_ip               = "example"
      source_application_port = 0
    }]
    exporter_group_alias = "example"
    gsop_alias           = "example"
    map_alias            = "example"
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
    source_port = [ "example" ]
  }
  gsgroup          = "example"
  ip_interface_in  = "example"
  ip_interface_out = "example"
  local_key_alias  = "example"
  remote_key_alias = "example"
  ssl_profile = {
    cipher  = "example"
    mtls    = "enable"
    version = "example"
  }
  ssl_profile_alias = "example"
  tcp_profile = {
    keep_alive_timer = 30
    selective_ack    = "enable"
    syn_retries      = 1
  }
  tcp_profile_alias = "example"
  traffic_dir       = "IN"
  ttl               = 1
  tunnel_type       = "example"
}
