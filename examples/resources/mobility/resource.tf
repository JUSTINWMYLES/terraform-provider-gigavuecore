resource "gigavuecore_mobility" "example" {
  sites = "example"
  solution_alias = "example"
  solution_type = "example"
  tags = [{
    tag_key = "example"
    tag_values = [ "example" ]
  }]
  traffic_policies = {
    for5_g = {
      gtp_flow_timeout = 1
      gtp_persistence = {
        enabled = true
        file_age_timeout = 1
        interval = 1
        restart_age_time = 1
      }
      load_balancing = {
        app_type = "example"
        hashing_key = "example"
      }
      overlap_mode = true
      sampling = {
        flow_maps = [{
          alias = "example"
          comment = "example"
          rules = [{
            comment = "example"
            control_plane_percentage = 1
            dnn = "example"
            gpsi = "example"
            nas_5_qi = "example"
            nci = "example"
            nsiid = "example"
            pei = "example"
            plmn_id = "example"
            supi = "example"
            tac = "example"
            user_plane_percentage = 1
          }]
          source_group_id = "example"
          tags = [{
            tag_key = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
      }
      whitelisting = {
        flow_maps = [{
          alias = "example"
          comment = "example"
          rules = [{
            dnn = "example"
            type = "example"
            whitelist_databases = [ "example" ]
          }]
          source_group_id = "example"
          tags = [{
            tag_key = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
        multi_whitelists = [ "example" ]
        white_list_alias = "example"
      }
    }
    for_lte = {
      gtp_flow_timeout = 1
      gtp_persistence = {
        enabled = true
        file_age_timeout = 1
        interval = 1
        restart_age_time = 1
      }
      overlap_mode = true
      sampling = {
        flow_maps = [{
          alias = "example"
          comment = "example"
          rules = [{
            apn = "example"
            comment = "example"
            control_plane_percentage = 1
            eci = "example"
            imei = "example"
            imsi = "example"
            interface = "example"
            msisdn = "example"
            nas_5_qi = "example"
            nci = "example"
            periodic_recalc = true
            plmn_id = "example"
            qci = 1
            snssai = "example"
            tac = "example"
            tac_5_g = "example"
            user_plane_percentage = 1
            version = "example"
          }]
          source_group_id = "example"
          tags = [{
            tag_key = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
      }
      whitelisting = {
        flow_maps = [{
          alias = "example"
          comment = "example"
          rules = [{
            apn = "example"
            interface = "example"
            type = "example"
            version = "example"
            whitelist_databases = [ "example" ]
          }]
          source_group_id = "example"
          tool = "example"
        }]
        multi_whitelists = [ "example" ]
        white_list_alias = "example"
      }
    }
    for_non_cups_lte = {
      flowfiltering = {
        flow_maps = [{
          alias = "example"
          comment = "example"
          drop_rules = [{
            imei = "example"
            imsi = "example"
            interface = "example"
            msisdn = "example"
            version = "example"
          }]
          pass_rules = [{
            imei = "example"
            imsi = "example"
            interface = "example"
            msisdn = "example"
            version = "example"
          }]
          source_group_id = "example"
          tags = [{
            tag_key = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
      }
      gtp_flow_timeout = 1
      gtp_persistence = {
        enabled = true
        file_age_timeout = 1
        interval = 1
        restart_age_time = 1
      }
      load_balancing = {
        app_type = "example"
        hashing_key = "example"
      }
      overlap_mode = true
      sampling = {
        flow_maps = [{
          alias = "example"
          comment = "example"
          rules = [{
            apn = "example"
            comment = "example"
            eci = "example"
            gtp_sample_percentage = 1
            imei = "example"
            imsi = "example"
            interface = "example"
            msisdn = "example"
            periodic_recalc = true
            plmn_id = "example"
            qci = 1
            tac = "example"
            version = "example"
          }]
          source_group_id = "example"
          tags = [{
            tag_key = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
      }
      whitelisting = {
        flow_maps = [{
          alias = "example"
          comment = "example"
          rules = [{
            apn = "example"
            interface = "example"
            type = "example"
            version = "example"
            whitelist_databases = [ "example" ]
          }]
          source_group_id = "example"
          tags = [{
            tag_key = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
        multi_whitelists = [ "example" ]
        white_list_alias = "example"
      }
    }
  }
}
