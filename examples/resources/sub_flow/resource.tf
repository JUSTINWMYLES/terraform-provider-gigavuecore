resource "gigavuecore_sub_flow" "example" {
  flow_rules = {
    drop_rules = [{
      gtp = {
        imei      = "*"
        imsi      = "*"
        interface = "Gn"
        msisdn    = "*"
        version   = "any"
      }
      rule_id = 1
    }]
    pass_rules = [{
      gtp = {
        imei      = "*"
        imsi      = "*"
        interface = "Gn"
        msisdn    = "*"
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
        gpsi     = "*"
        nas_5_qi = "0"
        nci      = "*"
        nsiid    = "0"
        pei      = "*"
        plmn_id  = "*"
        supi     = "*"
        tac      = "*"
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
        gpsi    = "*"
        nci     = "a1b2c3d4e"
        nsiid   = "0"
        pei     = "*"
        plmn_id = "123.45"
        supi    = "*"
        tac     = "*"
      }
      percentage = 0
      priority   = 1
      rule_id    = 1
    }]
  }
  flow_sample_diameter_rules = {
    pass_rules = [{
      diameter = {
        user_name = "*"
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
        eci       = "a1b2c3d4"
        imei      = "*"
        imsi      = "*"
        interface = "Gn"
        msisdn    = "*"
        nas_5_qi  = "0"
        nci       = "*"
        plmn_id   = "123.45"
        qci       = 0
        snssai    = "0"
        tac       = "abc1"
        tac_5_g   = "*"
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
        eci       = "a1b2c3d4"
        imei      = "*"
        imsi      = "*"
        interface = "Gn"
        msisdn    = "*"
        nas_5_qi  = "0"
        nci       = "*"
        plmn_id   = "123.45"
        qci       = 0
        snssai    = "0"
        tac       = "abc1"
        tac_5_g   = "*"
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
    whitelist_databases = ["example"]
  }
  flow_whitelist_overlap_rules = {
    pass_rules = [{
      flow5_g = {
        dnn                 = "example"
        type                = "example"
        whitelist_databases = ["example"]
      }
      gtp = {
        apn                 = "example"
        interface           = "Gn"
        type                = "example"
        version             = "v1"
        whitelist_databases = ["example"]
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
        whitelist_databases = ["example"]
      }
      gtp = {
        apn                 = "example"
        interface           = "Gn"
        type                = "example"
        version             = "v1"
        whitelist_databases = ["example"]
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
      matches = ["example"]
      rule_id = 1
    }]
    pass_rules = [{
      comment = "example"
      matches = ["example"]
      rule_id = 1
    }]
  }
  rule_ids  = "example"
  rule_type = "example"
}
