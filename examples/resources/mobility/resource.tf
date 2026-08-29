resource "gigavuecore_mobility" "example" {
  deployed     = true
  health_state = "example"
  site_name    = "example"
  site_tag     = "example"
  sites = [{
    alias = "example"
    cp_nodes = [{
      additional_tool_ports = [ "example" ]
      alias                 = "example"
      app5_g_http2_ports    = [ 0 ]
      app_tcp = {
        application  = "example"
        load_balance = true
        tcp_control  = "example"
      }
      collector_tools                     = [ "example" ]
      control_metadata_ip_interface_alias = "example"
      gtp_control_sample                  = true
      gtp_random_sampling = {
        enabled  = true
        interval = 0
      }
      ip_interface_alias = "example"
      location = {
        cluster_id   = "example"
        engine_ports = [ "example" ]
      }
      node_override_network_ports = [ "example" ]
      node_type                   = "example"
      number_of5_g_sessions       = 0
      number_of_lte_sessions      = 0
      tags = [{
        tag_key    = "example"
        tag_values = [ "example" ]
      }]
      traffic_sources = [{
        comment                = "example"
        expand_port_identifier = true
        group_interfaces       = true
        ip4_frag_rule_type     = "example"
        network_function_interfaces = [{
          comment                       = "example"
          interface_type                = "example"
          source_override_dst_ports     = [ "example" ]
          source_override_network_ports = [ "example" ]
          tunnel_identifiers            = [ "example" ]
        }]
        network_function_name = "example"
        network_function_type = "example"
        source_group_id       = "example"
        tags = [{
          tag_key    = "example"
          tag_values = [ "example" ]
        }]
      }]
    }]
    gtp_nodes = [{
      additional_tool_ports = [ "example" ]
      alias                 = "example"
      collector_tools       = [ "example" ]
      gtp_random_sampling = {
        enabled  = true
        interval = 0
      }
      location = {
        cluster_id   = "example"
        engine_ports = [ "example" ]
      }
      node_override_network_ports = [ "example" ]
      node_type                   = "example"
      tags = [{
        tag_key    = "example"
        tag_values = [ "example" ]
      }]
      traffic_sources = [{
        comment                = "example"
        expand_port_identifier = true
        group_interfaces       = true
        ip4_frag_rule_type     = "example"
        network_function_interfaces = [{
          comment                       = "example"
          interface_type                = "example"
          source_override_dst_ports     = [ "example" ]
          source_override_network_ports = [ "example" ]
          tunnel_identifiers            = [ "example" ]
        }]
        network_function_name = "example"
        network_function_type = "example"
        source_group_id       = "example"
        tags = [{
          tag_key    = "example"
          tag_values = [ "example" ]
        }]
      }]
    }]
    network_ports = [ "example" ]
    sam_exporter_nodes = [{
      smaf_details = [{
        control_application = {
          interface_address = "example"
          port              = 0
          protocol          = "example"
        }
        management_address = "example"
        user_application = {
          interface_address = "example"
          port              = 0
          protocol          = "example"
        }
      }]
      alias = "example"
      app_profile_config = {
        application_id = true
        applications = {
          attributes = [{
            name  = "example"
            value = "example"
          }]
          is_user_defined = true
          name            = "example"
        }
        counter = {
          bytes        = true
          bytes_long   = true
          packets      = true
          packets_long = true
        }
        flow = {
          end_reason = true
        }
        gtpu = {
          qfi  = true
          teid = true
        }
        inner_ipv4 = {
          destination = true
          protocol    = true
          source      = true
        }
        inner_ipv6 = {
          destination = true
          next_header = true
          source      = true
        }
        outer_ipv4 = {
          destination = true
          source      = true
        }
        outer_ipv6 = {
          destination = true
          source      = true
        }
        timestamp = {
          flow_end_msec   = true
          flow_endsec     = true
          flow_start_msec = true
          flow_startsec   = true
        }
        transport = {
          dst_port = true
          src_port = true
        }
      }
      control_plane_setting = {
        encoding        = "example"
        encoding_format = "example"
        event_enable = {
          modify = true
          update = true
        }
        trigger = "example"
      }
      engine_meta_data_cache_configs = [{
        engine_port   = "example"
        event         = "example"
        flow_behavior = "example"
        flows_size    = 0
        idle_timeout  = 0
        match = {
          ipv4 = {
            destination = {
              prefix_min_mask = "example"
            }
            protocol = true
            source = {
              prefix_min_mask = "example"
            }
          }
          ipv6 = {
            destination = {
              prefix_min_mask = "example"
            }
            next_header = true
            source = {
              prefix_min_mask = "example"
            }
          }
          transport = {
            dst_port = true
            src_port = true
          }
        }
        observation_domain_id = 0
      }]
      engine_source_mappings = [{
        engine_port    = "example"
        network_source = "example"
      }]
      exporter_config = {
        active_timeout   = "example"
        inactive_timeout = 0
        record_type      = "example"
      }
      ip_interface_alias = "example"
      location = {
        cluster_id   = "example"
        engine_ports = [ "example" ]
      }
      node_override_network_ports = [ "example" ]
      node_type                   = "example"
      param_configs = [{
        engine_port       = "example"
        resource_metadata = 0
      }]
      tags = [{
        tag_key    = "example"
        tag_values = [ "example" ]
      }]
      traffic_sources = [{
        comment                = "example"
        expand_port_identifier = true
        group_interfaces       = true
        ip4_frag_rule_type     = "example"
        network_function_interfaces = [{
          comment                       = "example"
          interface_type                = "example"
          source_override_dst_ports     = [ "example" ]
          source_override_network_ports = [ "example" ]
          tunnel_identifiers            = [ "example" ]
        }]
        network_function_name = "example"
        network_function_type = "example"
        source_group_id       = "example"
        tags = [{
          tag_key    = "example"
          tag_values = [ "example" ]
        }]
      }]
    }]
    site_override_of_policy_arrangements = {
      for5_g = {
        gtp_flow_timeout = 0
        gtp_persistence = {
          enabled          = true
          file_age_timeout = 0
          interval         = 0
          restart_age_time = 0
        }
        load_balancing = {
          app_type    = "example"
          hashing_key = "example"
        }
        overlap_mode = true
        sampling = {
          flow_maps = [{
            alias   = "example"
            comment = "example"
            rules = [{
              comment                  = "example"
              control_plane_percentage = 0
              dnn                      = "example"
              gpsi                     = "example"
              nas_5_qi                 = "example"
              nci                      = "example"
              nsiid                    = "example"
              pei                      = "example"
              plmn_id                  = "example"
              supi                     = "example"
              tac                      = "example"
              user_plane_percentage    = 0
            }]
            source_group_id = "example"
            tags = [{
              tag_key    = "example"
              tag_values = [ "example" ]
            }]
            tool = "example"
          }]
        }
        whitelisting = {
          flow_maps = [{
            alias   = "example"
            comment = "example"
            rules = [{
              dnn                 = "example"
              type                = "example"
              whitelist_databases = [ "example" ]
            }]
            source_group_id = "example"
            tags = [{
              tag_key    = "example"
              tag_values = [ "example" ]
            }]
            tool = "example"
          }]
          multi_whitelists = [ "example" ]
          white_list_alias = "example"
        }
      }
      for_lte = {
        gtp_flow_timeout = 0
        gtp_persistence = {
          enabled          = true
          file_age_timeout = 0
          interval         = 0
          restart_age_time = 0
        }
        overlap_mode = true
        sampling = {
          flow_maps = [{
            alias   = "example"
            comment = "example"
            rules = [{
              apn                      = "example"
              comment                  = "example"
              control_plane_percentage = 0
              eci                      = "example"
              imei                     = "example"
              imsi                     = "example"
              interface                = "example"
              msisdn                   = "example"
              nas_5_qi                 = "example"
              nci                      = "example"
              periodic_recalc          = true
              plmn_id                  = "example"
              qci                      = 0
              snssai                   = "example"
              tac                      = "example"
              tac_5_g                  = "example"
              user_plane_percentage    = 0
              version                  = "example"
            }]
            source_group_id = "example"
            tags = [{
              tag_key    = "example"
              tag_values = [ "example" ]
            }]
            tool = "example"
          }]
        }
        whitelisting = {
          flow_maps = [{
            alias   = "example"
            comment = "example"
            rules = [{
              apn                 = "example"
              interface           = "example"
              type                = "example"
              version             = "example"
              whitelist_databases = [ "example" ]
            }]
            source_group_id = "example"
            tool            = "example"
          }]
          multi_whitelists = [ "example" ]
          white_list_alias = "example"
        }
      }
      for_non_cups_lte = {
        flowfiltering = {
          flow_maps = [{
            alias   = "example"
            comment = "example"
            drop_rules = [{
              imei      = "example"
              imsi      = "example"
              interface = "example"
              msisdn    = "example"
              version   = "example"
            }]
            pass_rules = [{
              imei      = "example"
              imsi      = "example"
              interface = "example"
              msisdn    = "example"
              version   = "example"
            }]
            source_group_id = "example"
            tags = [{
              tag_key    = "example"
              tag_values = [ "example" ]
            }]
            tool = "example"
          }]
        }
        gtp_flow_timeout = 0
        gtp_persistence = {
          enabled          = true
          file_age_timeout = 0
          interval         = 0
          restart_age_time = 0
        }
        load_balancing = {
          app_type    = "example"
          hashing_key = "example"
        }
        overlap_mode = true
        sampling = {
          flow_maps = [{
            alias   = "example"
            comment = "example"
            rules = [{
              apn                   = "example"
              comment               = "example"
              eci                   = "example"
              gtp_sample_percentage = 0
              imei                  = "example"
              imsi                  = "example"
              interface             = "example"
              msisdn                = "example"
              periodic_recalc       = true
              plmn_id               = "example"
              qci                   = 0
              tac                   = "example"
              version               = "example"
            }]
            source_group_id = "example"
            tags = [{
              tag_key    = "example"
              tag_values = [ "example" ]
            }]
            tool = "example"
          }]
        }
        whitelisting = {
          flow_maps = [{
            alias   = "example"
            comment = "example"
            rules = [{
              apn                 = "example"
              interface           = "example"
              type                = "example"
              version             = "example"
              whitelist_databases = [ "example" ]
            }]
            source_group_id = "example"
            tags = [{
              tag_key    = "example"
              tag_values = [ "example" ]
            }]
            tool = "example"
          }]
          multi_whitelists = [ "example" ]
          white_list_alias = "example"
        }
      }
    }
    skip_deployment = true
    tags = [{
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
    tool_bindings = [{
      alias                    = "example"
      meta_data_exporter_alias = "example"
      tool_cluster_id          = "example"
      tool_resource_id         = "example"
      tool_resource_type       = "example"
    }]
    up_nodes = [{
      additional_tool_ports = [ "example" ]
      alias                 = "example"
      collector_tools       = [ "example" ]
      gtp_control_sample    = true
      gtp_random_sampling = {
        enabled  = true
        interval = 0
      }
      ip_interface_alias = "example"
      location = {
        cluster_id   = "example"
        engine_ports = [ "example" ]
      }
      node_override_network_ports = [ "example" ]
      node_type                   = "example"
      number_of_lte_sessions      = 0
      stand_alone_mode            = true
      tags = [{
        tag_key    = "example"
        tag_values = [ "example" ]
      }]
      traffic_sources = [{
        comment                = "example"
        expand_port_identifier = true
        group_interfaces       = true
        ip4_frag_rule_type     = "example"
        network_function_interfaces = [{
          comment                       = "example"
          interface_type                = "example"
          source_override_dst_ports     = [ "example" ]
          source_override_network_ports = [ "example" ]
          tunnel_identifiers            = [ "example" ]
        }]
        network_function_name = "example"
        network_function_type = "example"
        source_group_id       = "example"
        tags = [{
          tag_key    = "example"
          tag_values = [ "example" ]
        }]
      }]
    }]
  }]
  solution_alias = "example"
  solution_type  = "example"
  tags = [{
    tag_key    = "example"
    tag_values = [ "example" ]
  }]
  traffic_policies = {
    for5_g = {
      gtp_flow_timeout = 0
      gtp_persistence = {
        enabled          = true
        file_age_timeout = 0
        interval         = 0
        restart_age_time = 0
      }
      load_balancing = {
        app_type    = "example"
        hashing_key = "example"
      }
      overlap_mode = true
      sampling = {
        flow_maps = [{
          alias   = "example"
          comment = "example"
          rules = [{
            comment                  = "example"
            control_plane_percentage = 0
            dnn                      = "example"
            gpsi                     = "example"
            nas_5_qi                 = "example"
            nci                      = "example"
            nsiid                    = "example"
            pei                      = "example"
            plmn_id                  = "example"
            supi                     = "example"
            tac                      = "example"
            user_plane_percentage    = 0
          }]
          source_group_id = "example"
          tags = [{
            tag_key    = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
      }
      whitelisting = {
        flow_maps = [{
          alias   = "example"
          comment = "example"
          rules = [{
            dnn                 = "example"
            type                = "example"
            whitelist_databases = [ "example" ]
          }]
          source_group_id = "example"
          tags = [{
            tag_key    = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
        multi_whitelists = [ "example" ]
        white_list_alias = "example"
      }
    }
    for_lte = {
      gtp_flow_timeout = 0
      gtp_persistence = {
        enabled          = true
        file_age_timeout = 0
        interval         = 0
        restart_age_time = 0
      }
      overlap_mode = true
      sampling = {
        flow_maps = [{
          alias   = "example"
          comment = "example"
          rules = [{
            apn                      = "example"
            comment                  = "example"
            control_plane_percentage = 0
            eci                      = "example"
            imei                     = "example"
            imsi                     = "example"
            interface                = "example"
            msisdn                   = "example"
            nas_5_qi                 = "example"
            nci                      = "example"
            periodic_recalc          = true
            plmn_id                  = "example"
            qci                      = 0
            snssai                   = "example"
            tac                      = "example"
            tac_5_g                  = "example"
            user_plane_percentage    = 0
            version                  = "example"
          }]
          source_group_id = "example"
          tags = [{
            tag_key    = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
      }
      whitelisting = {
        flow_maps = [{
          alias   = "example"
          comment = "example"
          rules = [{
            apn                 = "example"
            interface           = "example"
            type                = "example"
            version             = "example"
            whitelist_databases = [ "example" ]
          }]
          source_group_id = "example"
          tool            = "example"
        }]
        multi_whitelists = [ "example" ]
        white_list_alias = "example"
      }
    }
    for_non_cups_lte = {
      flowfiltering = {
        flow_maps = [{
          alias   = "example"
          comment = "example"
          drop_rules = [{
            imei      = "example"
            imsi      = "example"
            interface = "example"
            msisdn    = "example"
            version   = "example"
          }]
          pass_rules = [{
            imei      = "example"
            imsi      = "example"
            interface = "example"
            msisdn    = "example"
            version   = "example"
          }]
          source_group_id = "example"
          tags = [{
            tag_key    = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
      }
      gtp_flow_timeout = 0
      gtp_persistence = {
        enabled          = true
        file_age_timeout = 0
        interval         = 0
        restart_age_time = 0
      }
      load_balancing = {
        app_type    = "example"
        hashing_key = "example"
      }
      overlap_mode = true
      sampling = {
        flow_maps = [{
          alias   = "example"
          comment = "example"
          rules = [{
            apn                   = "example"
            comment               = "example"
            eci                   = "example"
            gtp_sample_percentage = 0
            imei                  = "example"
            imsi                  = "example"
            interface             = "example"
            msisdn                = "example"
            periodic_recalc       = true
            plmn_id               = "example"
            qci                   = 0
            tac                   = "example"
            version               = "example"
          }]
          source_group_id = "example"
          tags = [{
            tag_key    = "example"
            tag_values = [ "example" ]
          }]
          tool = "example"
        }]
      }
      whitelisting = {
        flow_maps = [{
          alias   = "example"
          comment = "example"
          rules = [{
            apn                 = "example"
            interface           = "example"
            type                = "example"
            version             = "example"
            whitelist_databases = [ "example" ]
          }]
          source_group_id = "example"
          tags = [{
            tag_key    = "example"
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
