action "gigavuecore_update_fabric_map_with_last_update_timestamp_validation" "example" {
  config {
    alias = "example"
    ap_rules = {
      drop_rules = [{
        application_profile = "example"
        rule_id             = 0
      }]
      pass_rules = [{
        application_profile = "example"
        rule_id             = 0
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
        type       = "example"
      }
      b_to_a = {
        ib_pathway = "example"
        tools      = [ "example" ]
        type       = "example"
      }
      oob_copy = [{
        direction = "example"
        dst_ports = [ "example" ]
        src_ports = [ "example" ]
        tag = {
          type = "example"
        }
      }]
      svt_mode = true
      svt_tag  = 0
      tag = {
        tag_protocol_id = "example"
        type            = "example"
        vlan_id         = 0
      }
    }
    flex_inline_failover = "example"
    flex_inline_vlan_id  = 0
    flow_rules = {
      drop_rules = [{
        gtp = {
          imei      = "example"
          imsi      = "example"
          interface = "example"
          msisdn    = "example"
          version   = "example"
        }
        rule_id = 0
      }]
      pass_rules = [{
        gtp = {
          imei      = "example"
          imsi      = "example"
          interface = "example"
          msisdn    = "example"
          version   = "example"
        }
        rule_id = 0
      }]
    }
    flow_sample_diameter_rules = {
      pass_rules = [{
        diameter = {
          user_name = "example"
        }
        interface  = "example"
        percentage = 0
        rule_id    = 0
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
          interface = "example"
          msisdn    = "example"
          nas_5_qi  = "example"
          nci       = "example"
          plmn_id   = "example"
          qci       = 0
          snssai    = "example"
          tac       = "example"
          tac_5_g   = "example"
          version   = "example"
        }
        percentage      = 0
        periodic_recalc = true
        priority        = 0
        rule_id         = 0
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
          interface = "example"
          msisdn    = "example"
          nas_5_qi  = "example"
          nci       = "example"
          plmn_id   = "example"
          qci       = 0
          snssai    = "example"
          tac       = "example"
          tac_5_g   = "example"
          version   = "example"
        }
        percentage      = 0
        periodic_recalc = true
        priority        = 0
        rule_id         = 0
      }]
    }
    flow_sample_sip_rules = {
      pass_rules = [{
        percentage = 0
        rule_id    = 0
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
    flow_whitelist_overlap_rules = {
      pass_rules = [{
        flow5_g = {
          dnn                 = "example"
          type                = "example"
          whitelist_databases = [ "example" ]
        }
        gtp = {
          apn                 = "example"
          interface           = "example"
          type                = "example"
          version             = "example"
          whitelist_databases = [ "example" ]
        }
        rule_id = 0
        sip = {
          type = "example"
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
          interface           = "example"
          type                = "example"
          version             = "example"
          whitelist_databases = [ "example" ]
        }
        rule_id = 0
        sip = {
          type = "example"
        }
      }]
    }
    gs_rules = {
      drop_rules = [{
        comment = "example"
        matches = [ "example" ]
        rule_id = 0
      }]
      pass_rules = [{
        comment = "example"
        matches = [ "example" ]
        rule_id = 0
      }]
    }
    gsop         = "example"
    health_state = "example"
    health_state_reasons = [{
      message                               = "example"
      severity                              = "example"
      traffic_health_state_computation_type = "example"
    }]
    inline_traffic_path = "example"
    inline_traffic_type = "example"
    mod_time            = 0
    order               = 0
    roles = {
      editors   = [ "example" ]
      listeners = [ "example" ]
      owners    = [ "example" ]
      viewers   = [ "example" ]
    }
    rule_matching = "example"
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
    rx_cluster_ports = [ "example" ]
    src_ports        = [ "example" ]
    sub_type         = "example"
    traffic_type     = "example"
    tx_cluster_ports = [ "example" ]
    type             = "example"
    updated_time     = 0
  }
}
