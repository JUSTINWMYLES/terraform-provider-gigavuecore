resource "gigavuecore_sub_flow" "example" {
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
  rule_ids  = "example"
  rule_type = "example"
}
