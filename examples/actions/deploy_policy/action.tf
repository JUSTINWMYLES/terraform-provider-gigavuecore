action "gigavuecore_deploy_policy" "example" {
  config {
    comment             = "example"
    deployed            = true
    deployment_error    = "example"
    deployment_percent  = "example"
    dest_port_timestamp = "example"
    health_state        = "example"
    health_state_reasons = [{
      message                               = "example"
      severity                              = "example"
      traffic_health_state_computation_type = "example"
    }]
    name             = "example"
    policy_id        = "example"
    policy_timestamp = "example"
    priority         = true
    rules = [{
      comment = "example"
      gs_operations = {
        gs_apps = {
          apf = {
            enabled = "example"
          }
          dedup = {
            enabled = "example"
          }
          diameter_whitelist = {
            enabled = "example"
          }
          flow_filter = {
            type = "example"
          }
          flow_sampling = {
            type = "example"
          }
          gseries_header_add = {
            types = [ "example" ]
          }
          gseries_header_remove = {
            enabled = "example"
          }
          gseries_load_balance = {
            fixed_offset = {
              hash   = "example"
              length = 0
              offset = 0
            }
            variable_offset = {
              end_delim   = "example"
              hash        = "example"
              start_delim = "example"
              start_field = "example"
            }
          }
          gseries_pattern_match = {
            fixed_offset = {
              length = 0
              offset = 0
            }
            variable_offset = {
              end_delim   = "example"
              start_delim = "example"
            }
          }
          gtp_whitelist = {
            enabled = "example"
          }
          header_add = {
            vlan = 0
          }
          header_remove = {
            ah1                = "example"
            ah2                = "example"
            custom_len         = 0
            erspan_flow_id     = 0
            fp_dst_switch_id   = 0
            fp_src_switch_id   = 0
            header_count       = 0
            offset             = "example"
            offset_range_value = 0
            protocol           = "example"
            timestamp_format   = "example"
            vlan_header        = "example"
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
              app_type               = "example"
              diameter_key_hash_type = "example"
              diameter_key_multi_hash_type = [{
                avp_codevalue = 0
                key           = "example"
              }]
              gtp_key_hash_type = "example"
              lb_type           = "example"
              sip_key_hash_type = "example"
            }
            stateless = {
              field_location = "example"
              hash_fields    = "example"
            }
          }
          masking = {
            content_type = "example"
            length       = 0
            offset       = 0
            pattern      = "example"
            protocol     = "example"
          }
          metadata_export = {
            cache = "example"
          }
          netflow = {
            enabled = "example"
          }
          sa_apf = {
            enabled = "example"
          }
          sip_whitelist = {
            enabled = "example"
          }
          slicing = {
            enhanced = "example"
            offset   = 0
            protocol = "example"
          }
          ssl_decrypt = {
            in_port  = 0
            out_port = 0
          }
          trailer_add = {
            types = [ "example" ]
          }
          trailer_remove = {
            enabled = "example"
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
            type = "example"
            vxlan = {
              port_dst = 0
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
              ttl        = 0
            }
            l2_gre_config = {
              dscp          = 0
              dst_ip        = "example"
              flow_label    = 0
              key           = 0
              pg_dst        = "example"
              prec          = 0
              session_field = "example"
              session_pos   = "example"
              ttl           = 0
            }
            tls_pcapng = {
              exporter       = "example"
              exporter_group = "example"
            }
            type = "example"
            vxlan_config = {
              dscp     = 0
              dst_ip   = "example"
              dst_port = 0
              src_port = 0
              ttl      = 0
              vni      = 0
            }
          }
        }
        gs_engines = [ "example" ]
        gs_group_params = {
          app_tcp = {
            application  = "example"
            load_balance = true
            tcp_control  = "example"
          }
          dedup = {
            action    = "example"
            ip_tclass = "example"
            ip_tos    = "example"
            tcp_seq   = "example"
            timer     = 0
            vlan      = "example"
          }
          diameter_packet = {
            timeout = 0
          }
          diameter_s6_a_session = {
            limit   = 0
            timeout = 0
          }
          diameter_whitelist = {
            whitelist = "example"
          }
          eflow = {
            enabled      = true
            interval     = 0
            log_enabled  = true
            packet_count = 0
            packet_ratio = 0
          }
          engine_watchdog_timer = {
            time = 0
          }
          erspan3 = {
            timestamp_format = "example"
          }
          flow_mask = {
            enabled = true
            length  = 0
            offset  = 0
          }
          flow_sampling = {
            ip_ranges = [ "example" ]
            rate      = 0
            timeout   = 0
            type      = "example"
          }
          generic_session_timeout = {
            time = 0
          }
          gpfcp_profiles = {
            g_pfcp_profiles = [ "example" ]
          }
          gs_group_system = {
            cpu_load_alarm_threshold = 0
          }
          gta_profiles = {
            gta_profiles = [ "example" ]
          }
          gtp_control_sampling = {
            enabled = true
          }
          gtp_flow = {
            timeout = 0
          }
          gtp_gpfcp_delay = {
            timeout = 0
          }
          gtp_persistence = {
            enabled          = true
            file_age_timeout = 0
            interval         = 0
            restart_age_time = 0
          }
          gtp_random_sampling = {
            enabled  = true
            interval = 0
          }
          gtp_whitelist = {
            multi_whitelists = [ "example" ]
            whitelist        = "example"
          }
          health_check = {
            action          = "example"
            dst_port        = 0
            enabled         = true
            interval        = 0
            protocol        = "example"
            rcv_port        = 0
            retries         = 0
            round_trip_time = 0
            src_port        = 0
          }
          hsm_group = {
            hsm_group = "example"
          }
          ip_frag = {
            forward              = true
            head_session_timeout = 0
            timeout              = 0
          }
          load_balance = {
            failover = {
              enabled               = true
              threshold_lt_bw       = 0
              threshold_lt_pkt_rate = 0
            }
            link_weight_type = "example"
            replicate_gtpc   = true
          }
          netflow = {
            monitor = "example"
          }
          node_role = {
            mob5_g_limit     = 0
            mob_lte_limit    = 0
            stand_alone_mode = true
            type             = "example"
          }
          port_throttle_sip = {
            port_throttle = "example"
          }
          resource = {
            buffer_asf_size = 0
            cpu = {
              overload_threshold = 0
            }
            hsm_ssl = {
              buffer        = 0
              packet_buffer = 0
              session_count = 0
            }
            inline_ssl = {
              standalone = true
            }
            metadata = 0
            packet_buffer = {
              overload_threshold = 0
            }
            session_overload = {
              overload_threshold = 0
            }
            tunnel_overload = {
              overload_threshold = 0
            }
            xpkt_match = {
              flows = 0
            }
          }
          rtp_ports = {
            range = {
              port     = 0
              port_max = 0
            }
          }
          sa_apf = {
            buffer_size = 0
          }
          session_logging = {
            interface          = "example"
            log_level          = "example"
            remote_syslog_ip   = "example"
            remote_syslog_port = 0
          }
          sffp_profiles = {
            sffp_profiles = [ "example" ]
          }
          sip_media = {
            timeout = 0
          }
          sip_ports = {
            ports = [ 0 ]
          }
          sip_session = {
            timeout = 0
          }
          sip_tcp_idle_timeout = {
            time = 0
          }
          sip_whitelist = {
            whitelist = "example"
          }
          ssl_decrypt = {
            decrypt_fail_action = "example"
            enabled             = true
            hsm_pkcs11 = {
              debug_level    = 0
              dynamic_object = true
              load_sharing   = true
            }
            hsm_timeout             = 0
            key_cache_timeout       = 0
            key_map                 = "example"
            non_ssl_traffic         = "example"
            pending_session_timeout = 0
            session_timeout         = 0
            tcp_syn_timeout         = 0
            ticket_cache_timeout    = 0
          }
          xpkt_match = {
            enabled = true
          }
        }
        templates = [{
          gs_apps = {
            apf = {
              enabled = "example"
            }
            dedup = {
              enabled = "example"
            }
            diameter_whitelist = {
              enabled = "example"
            }
            flow_filter = {
              type = "example"
            }
            flow_sampling = {
              type = "example"
            }
            gseries_header_add = {
              types = [ "example" ]
            }
            gseries_header_remove = {
              enabled = "example"
            }
            gseries_load_balance = {
              fixed_offset = {
                hash   = "example"
                length = 0
                offset = 0
              }
              variable_offset = {
                end_delim   = "example"
                hash        = "example"
                start_delim = "example"
                start_field = "example"
              }
            }
            gseries_pattern_match = {
              fixed_offset = {
                length = 0
                offset = 0
              }
              variable_offset = {
                end_delim   = "example"
                start_delim = "example"
              }
            }
            gtp_whitelist = {
              enabled = "example"
            }
            header_add = {
              vlan = 0
            }
            header_remove = {
              ah1                = "example"
              ah2                = "example"
              custom_len         = 0
              erspan_flow_id     = 0
              fp_dst_switch_id   = 0
              fp_src_switch_id   = 0
              header_count       = 0
              offset             = "example"
              offset_range_value = 0
              protocol           = "example"
              timestamp_format   = "example"
              vlan_header        = "example"
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
                app_type               = "example"
                diameter_key_hash_type = "example"
                diameter_key_multi_hash_type = [{
                  avp_codevalue = 0
                  key           = "example"
                }]
                gtp_key_hash_type = "example"
                lb_type           = "example"
                sip_key_hash_type = "example"
              }
              stateless = {
                field_location = "example"
                hash_fields    = "example"
              }
            }
            masking = {
              content_type = "example"
              length       = 0
              offset       = 0
              pattern      = "example"
              protocol     = "example"
            }
            metadata_export = {
              cache = "example"
            }
            netflow = {
              enabled = "example"
            }
            sa_apf = {
              enabled = "example"
            }
            sip_whitelist = {
              enabled = "example"
            }
            slicing = {
              enhanced = "example"
              offset   = 0
              protocol = "example"
            }
            ssl_decrypt = {
              in_port  = 0
              out_port = 0
            }
            trailer_add = {
              types = [ "example" ]
            }
            trailer_remove = {
              enabled = "example"
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
              type = "example"
              vxlan = {
                port_dst = 0
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
                ttl        = 0
              }
              l2_gre_config = {
                dscp          = 0
                dst_ip        = "example"
                flow_label    = 0
                key           = 0
                pg_dst        = "example"
                prec          = 0
                session_field = "example"
                session_pos   = "example"
                ttl           = 0
              }
              tls_pcapng = {
                exporter       = "example"
                exporter_group = "example"
              }
              type = "example"
              vxlan_config = {
                dscp     = 0
                dst_ip   = "example"
                dst_port = 0
                src_port = 0
                ttl      = 0
                vni      = 0
              }
            }
          }
          gs_group_params = {
            app_tcp = {
              application  = "example"
              load_balance = true
              tcp_control  = "example"
            }
            dedup = {
              action    = "example"
              ip_tclass = "example"
              ip_tos    = "example"
              tcp_seq   = "example"
              timer     = 0
              vlan      = "example"
            }
            diameter_packet = {
              timeout = 0
            }
            diameter_s6_a_session = {
              limit   = 0
              timeout = 0
            }
            diameter_whitelist = {
              whitelist = "example"
            }
            eflow = {
              enabled      = true
              interval     = 0
              log_enabled  = true
              packet_count = 0
              packet_ratio = 0
            }
            engine_watchdog_timer = {
              time = 0
            }
            erspan3 = {
              timestamp_format = "example"
            }
            flow_mask = {
              enabled = true
              length  = 0
              offset  = 0
            }
            flow_sampling = {
              ip_ranges = [ "example" ]
              rate      = 0
              timeout   = 0
              type      = "example"
            }
            generic_session_timeout = {
              time = 0
            }
            gpfcp_profiles = {
              g_pfcp_profiles = [ "example" ]
            }
            gs_group_system = {
              cpu_load_alarm_threshold = 0
            }
            gta_profiles = {
              gta_profiles = [ "example" ]
            }
            gtp_control_sampling = {
              enabled = true
            }
            gtp_flow = {
              timeout = 0
            }
            gtp_gpfcp_delay = {
              timeout = 0
            }
            gtp_persistence = {
              enabled          = true
              file_age_timeout = 0
              interval         = 0
              restart_age_time = 0
            }
            gtp_random_sampling = {
              enabled  = true
              interval = 0
            }
            gtp_whitelist = {
              multi_whitelists = [ "example" ]
              whitelist        = "example"
            }
            health_check = {
              action          = "example"
              dst_port        = 0
              enabled         = true
              interval        = 0
              protocol        = "example"
              rcv_port        = 0
              retries         = 0
              round_trip_time = 0
              src_port        = 0
            }
            hsm_group = {
              hsm_group = "example"
            }
            ip_frag = {
              forward              = true
              head_session_timeout = 0
              timeout              = 0
            }
            load_balance = {
              failover = {
                enabled               = true
                threshold_lt_bw       = 0
                threshold_lt_pkt_rate = 0
              }
              link_weight_type = "example"
              replicate_gtpc   = true
            }
            netflow = {
              monitor = "example"
            }
            node_role = {
              mob5_g_limit     = 0
              mob_lte_limit    = 0
              stand_alone_mode = true
              type             = "example"
            }
            port_throttle_sip = {
              port_throttle = "example"
            }
            resource = {
              buffer_asf_size = 0
              cpu = {
                overload_threshold = 0
              }
              hsm_ssl = {
                buffer        = 0
                packet_buffer = 0
                session_count = 0
              }
              inline_ssl = {
                standalone = true
              }
              metadata = 0
              packet_buffer = {
                overload_threshold = 0
              }
              session_overload = {
                overload_threshold = 0
              }
              tunnel_overload = {
                overload_threshold = 0
              }
              xpkt_match = {
                flows = 0
              }
            }
            rtp_ports = {
              range = {
                port     = 0
                port_max = 0
              }
            }
            sa_apf = {
              buffer_size = 0
            }
            session_logging = {
              interface          = "example"
              log_level          = "example"
              remote_syslog_ip   = "example"
              remote_syslog_port = 0
            }
            sffp_profiles = {
              sffp_profiles = [ "example" ]
            }
            sip_media = {
              timeout = 0
            }
            sip_ports = {
              ports = [ 0 ]
            }
            sip_session = {
              timeout = 0
            }
            sip_tcp_idle_timeout = {
              time = 0
            }
            sip_whitelist = {
              whitelist = "example"
            }
            ssl_decrypt = {
              decrypt_fail_action = "example"
              enabled             = true
              hsm_pkcs11 = {
                debug_level    = 0
                dynamic_object = true
                load_sharing   = true
              }
              hsm_timeout             = 0
              key_cache_timeout       = 0
              key_map                 = "example"
              non_ssl_traffic         = "example"
              pending_session_timeout = 0
              session_timeout         = 0
              tcp_syn_timeout         = 0
              ticket_cache_timeout    = 0
            }
            xpkt_match = {
              enabled = true
            }
          }
          template_id = "example"
        }]
      }
      health_state = "example"
      health_state_reasons = [{
        message                               = "example"
        severity                              = "example"
        traffic_health_state_computation_type = "example"
      }]
      high_priority_drop = true
      matches = [{
        aggregate_filters = [{
          name           = "example"
          operation_type = "example"
          simple_filters = {
            advanced_filter_properties = {
              key    = "example"
              values = [ "example" ]
            }
            type = "example"
            values = {
              mask      = "example"
              offset    = "example"
              subset    = "example"
              value     = "example"
              value_max = "example"
            }
          }
          template_id = "example"
        }]
        operation_type   = "example"
        rule_criteria_id = "example"
      }]
      no_expansion_tags = [ "example" ]
      rule_id           = "example"
      rule_name         = "example"
      tools = [{
        alias          = "example"
        is_drop        = true
        is_giga_stream = true
        is_port        = true
        via_gsop       = true
      }]
      type = "example"
    }]
    src_port_timestamp = "example"
    src_ports_info = {
      comment      = "example"
      ports        = [ "example" ]
      template_ids = [ "example" ]
    }
    tags = [{
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
  }
}
