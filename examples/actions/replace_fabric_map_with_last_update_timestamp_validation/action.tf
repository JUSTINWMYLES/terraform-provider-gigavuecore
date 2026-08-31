action "gigavuecore_replace_fabric_map_with_last_update_timestamp_validation" "example" {
  config {
    alias = "example"
    ap_rules = {
      drop_rules = [{
        application_profile = "example"
        rule_id             = 1
      }]
      pass_rules = [{
        application_profile = "example"
        rule_id             = 1
      }]
    }
    body_alias        = "example"
    cluster_id        = "example"
    comment           = "example"
    dst_ports         = [ "example" ]
    egress_gigastream = [ "example" ]
    enable            = true
    encap_tunnel      = "example"
    flex_inline = {
      a_to_b = {
        ib_pathway = "example"
        tools      = [ "example" ]
        type       = "bypass"
      }
      b_to_a = {
        ib_pathway = "example"
        tools      = [ "example" ]
        type       = "bypass"
      }
      oob_copy = [{
        direction = "aToB"
        dst_ports = [ "example" ]
        src_ports = [ "example" ]
        tag = {
          type = "none"
        }
      }]
      svt_mode = true
      svt_tag  = 0
      tag = {
        tag_protocol_id = "0x8100"
        type            = "auto"
        vlan_id         = 0
      }
    }
    flex_inline_failover = "bypass"
    flex_inline_vlan_id  = 1
    flow_rules = {
      drop_rules = [{
        gtp = {
          imei      = "example"
          imsi      = "example"
          interface = "Gn"
          msisdn    = "example"
          version   = "any"
        }
        rule_id = 1
      }]
      pass_rules = [{
        gtp = {
          imei      = "example"
          imsi      = "example"
          interface = "Gn"
          msisdn    = "example"
          version   = "any"
        }
        rule_id = 1
      }]
    }
    flow_sample5_g_overlap_rules = {
      pass_rules = [{
        comment = "example"
        flow5_g = {
          dnn      = "example"
          gpsi     = "example"
          nas_5_qi = "0"
          nci      = "example"
          nsiid    = "0"
          pei      = "example"
          plmn_id  = "example"
          supi     = "example"
          tac      = "example"
        }
        percentage = 0
        rule_id    = 1
      }]
    }
    flow_sample5_g_rules = {
      pass_rules = [{
        comment = "example"
        flow5_g = {
          dnn     = "example"
          gpsi    = "example"
          nci     = "example"
          nsiid   = "0"
          pei     = "example"
          plmn_id = "example"
          supi    = "example"
          tac     = "example"
        }
        percentage = 0
        priority   = 1
        rule_id    = 1
      }]
    }
    flow_sample_diameter_rules = {
      pass_rules = [{
        diameter = {
          user_name = "example"
        }
        interface  = "s6a"
        percentage = 0
        rule_id    = 1
      }]
    }
    flow_sample_overlap_rules = {
      pass_rules = [{
        comment = "example"
        gtp = {
          apn       = "example"
          eci       = "example"
          imei      = "example"
          imsi      = "example"
          interface = "Gn"
          msisdn    = "example"
          nas_5_qi  = "0"
          nci       = "example"
          plmn_id   = "example"
          qci       = 0
          snssai    = "0"
          tac       = "abc1"
          tac_5_g   = "example"
          version   = "any"
        }
        percentage      = 0
        periodic_recalc = true
        priority        = 1
        rule_id         = 1
      }]
    }
    flow_sample_rules = {
      pass_rules = [{
        comment = "example"
        gtp = {
          apn       = "example"
          eci       = "example"
          imei      = "example"
          imsi      = "example"
          interface = "Gn"
          msisdn    = "example"
          nas_5_qi  = "0"
          nci       = "example"
          plmn_id   = "example"
          qci       = 0
          snssai    = "0"
          tac       = "abc1"
          tac_5_g   = "example"
          version   = "any"
        }
        percentage      = 0
        periodic_recalc = true
        priority        = 1
        rule_id         = 1
      }]
    }
    flow_sample_sip_rules = {
      pass_rules = [{
        percentage = 0
        rule_id    = 1
        sip = {
          callee_id = "example"
          callee_id_range = {
            max_value = "example"
            value     = "example"
          }
          caller_id = "example"
          caller_id_range = {
            max_value = "example"
            value     = "example"
          }
          id_range = {
            max_value = "example"
            value     = "example"
          }
        }
      }]
    }
    flow_whitelist5_g_overlap_rules = {
      dnn  = "example"
      type = "example"
    }
    flow_whitelist5_g_rules = {
      dnn                 = "example"
      type                = "example"
      whitelist_databases = [ "example" ]
    }
    flow_whitelist_overlap_rules = {
      pass_rules = [{
        flow5_g = {
          dnn                 = "example"
          type                = "example"
          whitelist_databases = [ "example" ]
        }
        gtp = {
          apn                 = "example"
          interface           = "Gn"
          type                = "example"
          version             = "v1"
          whitelist_databases = [ "example" ]
        }
        rule_id = 1
        sip = {
          type = "all"
        }
      }]
    }
    flow_whitelist_rules = {
      pass_rules = [{
        flow5_g = {
          dnn                 = "example"
          type                = "example"
          whitelist_databases = [ "example" ]
        }
        gtp = {
          apn                 = "example"
          interface           = "Gn"
          type                = "example"
          version             = "v1"
          whitelist_databases = [ "example" ]
        }
        rule_id = 1
        sip = {
          type = "all"
        }
      }]
    }
    fstype = {
      offset = 1
      timer  = 15
      type   = "_default"
    }
    gs_rules = {
      drop_rules = [{
        comment = "example"
        matches = [ "example" ]
        rule_id = 1
      }]
      pass_rules = [{
        comment = "example"
        matches = [ "example" ]
        rule_id = 1
      }]
    }
    gsop         = "example"
    health_state = "green"
    health_state_reasons = [{
      message                               = "example"
      severity                              = "green"
      traffic_health_state_computation_type = "PORT_LOW_UTIL"
    }]
    inline_traffic_path = "normal"
    inline_traffic_type = "symmetric"
    ip_rewrite = {
      dst_ip = "example"
      src_ip = "example"
    }
    mod_time      = 0
    null_dst_port = true
    order         = 0
    rewrite = {
      dst_mac = "example"
      src_mac = "example"
    }
    roles = {
      editors   = [ "example" ]
      listeners = [ "example" ]
      owners    = [ "example" ]
      viewers   = [ "example" ]
    }
    rule_matching = "normal"
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
    rx_cluster_ports = [ "example" ]
    src_ports        = [ "example" ]
    sub_type         = "byRule"
    traffic_type     = "control"
    tx_cluster_ports = [ "example" ]
    type             = "regular"
    updated_time     = 1.0
    vlan_tag = {
      tag_protocol_id = "0x8100"
      vlan_action     = "add"
      vlan_id         = 0
    }
  }
}
