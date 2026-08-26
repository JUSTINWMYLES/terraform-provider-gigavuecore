resource "gigavuecore_map" "example" {
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
    svt_tag  = 1
    tag = {
      tag_protocol_id = "example"
      type            = "example"
      vlan_id         = 1
    }
  }
  flex_inline_failover = "example"
  flex_inline_vlan_id  = 1
  flow_rules = {
    drop_rules = [{
      gtp = {
        imei      = "example"
        imsi      = "example"
        interface = "example"
        msisdn    = "example"
        version   = "example"
      }
      rule_id = 1
    }]
    pass_rules = [{
      gtp = {
        imei      = "example"
        imsi      = "example"
        interface = "example"
        msisdn    = "example"
        version   = "example"
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
        nas_5_qi = "example"
        nci      = "example"
        nsiid    = "example"
        pei      = "example"
        plmn_id  = "example"
        supi     = "example"
        tac      = "example"
      }
      percentage = 1
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
        nsiid   = "example"
        pei     = "example"
        plmn_id = "example"
        supi    = "example"
        tac     = "example"
      }
      percentage = 1
      priority   = 1
      rule_id    = 1
    }]
  }
  flow_sample_diameter_rules = {
    pass_rules = [{
      diameter = {
        user_name = "example"
      }
      interface  = "example"
      percentage = 1
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
        interface = "example"
        msisdn    = "example"
        nas_5_qi  = "example"
        nci       = "example"
        plmn_id   = "example"
        qci       = 1
        snssai    = "example"
        tac       = "example"
        tac_5_g   = "example"
        version   = "example"
      }
      percentage      = 1
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
        interface = "example"
        msisdn    = "example"
        nas_5_qi  = "example"
        nci       = "example"
        plmn_id   = "example"
        qci       = 1
        snssai    = "example"
        tac       = "example"
        tac_5_g   = "example"
        version   = "example"
      }
      percentage      = 1
      periodic_recalc = true
      priority        = 1
      rule_id         = 1
    }]
  }
  flow_sample_sip_rules = {
    pass_rules = [{
      percentage = 1
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
        interface           = "example"
        type                = "example"
        version             = "example"
        whitelist_databases = [ "example" ]
      }
      rule_id = 1
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
      rule_id = 1
      sip = {
        type = "example"
      }
    }]
  }
  fstype = {
    offset = 1
    timer  = 1
    type   = "example"
  }
  gs_rules = {
    drop_rules = null
    pass_rules = null
  }
  gsop                = "example"
  inline_traffic_path = "example"
  inline_traffic_type = "example"
  ip_rewrite = {
    dst_ip = "example"
    src_ip = "example"
  }
  null_dst_port = true
  order         = 1
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
  rule_matching = "example"
  rules = {
    drop_rules = null
    pass_rules = null
  }
  src_ports    = [ "example" ]
  sub_type     = "example"
  traffic_type = "example"
  type         = "example"
  vlan_tag = {
    tag_protocol_id = "example"
    vlan_action     = "example"
    vlan_id         = 1
  }
}
