action "gigavuecore_add_flow_rules" "example" {
  config {
    alias      = "example"
    flow_alias = "example"
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
        percentage = 0
        rule_id    = 0
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
        percentage = 0
        priority   = 0
        rule_id    = 0
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
    rule_type      = "example"
    sub_flow_alias = "example"
  }
}
