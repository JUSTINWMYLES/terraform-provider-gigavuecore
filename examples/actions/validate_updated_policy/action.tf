action "gigavuecore_validate_updated_policy" "example" {
  config {
    body_name = "example"
    comment   = "example"
    dest_ports = [{
      operation = "add"
      port_ids  = [ "example" ]
      rule_id   = "example"
      tools = [{
        alias          = "example"
        is_drop        = true
        is_giga_stream = true
        is_port        = true
        via_gsop       = true
      }]
    }]
    gigasmart_info = [{
      gs_operations = {
        gs_apps = {
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
        gs_engines = [ "example" ]
        gs_group_params = {
          app_tcp = {
            application  = "broadcast"
            load_balance = true
            tcp_control  = "broadcast"
          }
          dedup = {
            action    = "count"
            ip_tclass = "include"
            ip_tos    = "include"
            tcp_seq   = "include"
            timer     = 10
            vlan      = "include"
          }
          diameter_packet = {
            timeout = 1
          }
          diameter_s6_a_session = {
            limit   = 1
            timeout = 30
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
            timestamp_format = "gigasmart"
          }
          flow_mask = {
            enabled = true
            length  = 1
            offset  = 0
          }
          flow_sampling = {
            ip_ranges = [ "example" ]
            rate      = 5
            timeout   = 1
            type      = "deviceIp"
          }
          generic_session_timeout = {
            time = 5
          }
          gpfcp_profiles = {
            g_pfcp_profiles = [ "example" ]
          }
          gs_group_system = {
            cpu_load_alarm_threshold = 20
          }
          gta_profiles = {
            gta_profiles = [ "example" ]
          }
          gtp_control_sampling = {
            enabled = true
          }
          gtp_flow = {
            timeout = 1
          }
          gtp_gpfcp_delay = {
            timeout = 0
          }
          gtp_persistence = {
            enabled          = true
            file_age_timeout = 10
            interval         = 10
            restart_age_time = 10
          }
          gtp_random_sampling = {
            enabled  = true
            interval = 12
          }
          gtp_whitelist = {
            multi_whitelists = [ "example" ]
            whitelist        = "example"
          }
          health_check = {
            action          = "pass"
            dst_port        = 1
            enabled         = true
            interval        = 5
            protocol        = "icmp"
            rcv_port        = 1
            retries         = 1
            round_trip_time = 1
            src_port        = 1
          }
          hsm_group = {
            hsm_group = "example"
          }
          ip_frag = {
            forward              = true
            head_session_timeout = 15
            timeout              = 5
          }
          load_balance = {
            failover = {
              enabled               = true
              threshold_lt_bw       = 50
              threshold_lt_pkt_rate = 500
            }
            link_weight_type = "speed"
            replicate_gtpc   = true
          }
          netflow = {
            monitor = "example"
          }
          node_role = {
            mob5_g_limit     = 1
            mob_lte_limit    = 1
            stand_alone_mode = true
            type             = "control"
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
              packet_buffer = 20
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
            log_level          = "err"
            remote_syslog_ip   = "example"
            remote_syslog_port = 0
          }
          sffp_profiles = {
            sffp_profiles = [ "example" ]
          }
          sip_media = {
            timeout = 30
          }
          sip_ports = {
            ports = [ 1 ]
          }
          sip_session = {
            timeout = 30
          }
          sip_tcp_idle_timeout = {
            time = 20
          }
          sip_whitelist = {
            whitelist = "example"
          }
          ssl_decrypt = {
            decrypt_fail_action = "drop"
            enabled             = true
            hsm_pkcs11 = {
              debug_level    = 0
              dynamic_object = true
              load_sharing   = true
            }
            hsm_timeout             = 2
            key_cache_timeout       = 1
            key_map                 = "example"
            non_ssl_traffic         = "drop"
            pending_session_timeout = 30
            session_timeout         = 30
            tcp_syn_timeout         = 20
            ticket_cache_timeout    = 1
          }
          xpkt_match = {
            enabled = true
          }
        }
        templates = [{
          gs_apps = {
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
          gs_group_params = {
            app_tcp = {
              application  = "broadcast"
              load_balance = true
              tcp_control  = "broadcast"
            }
            dedup = {
              action    = "count"
              ip_tclass = "include"
              ip_tos    = "include"
              tcp_seq   = "include"
              timer     = 10
              vlan      = "include"
            }
            diameter_packet = {
              timeout = 1
            }
            diameter_s6_a_session = {
              limit   = 1
              timeout = 30
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
              timestamp_format = "gigasmart"
            }
            flow_mask = {
              enabled = true
              length  = 1
              offset  = 0
            }
            flow_sampling = {
              ip_ranges = [ "example" ]
              rate      = 5
              timeout   = 1
              type      = "deviceIp"
            }
            generic_session_timeout = {
              time = 5
            }
            gpfcp_profiles = {
              g_pfcp_profiles = [ "example" ]
            }
            gs_group_system = {
              cpu_load_alarm_threshold = 20
            }
            gta_profiles = {
              gta_profiles = [ "example" ]
            }
            gtp_control_sampling = {
              enabled = true
            }
            gtp_flow = {
              timeout = 1
            }
            gtp_gpfcp_delay = {
              timeout = 0
            }
            gtp_persistence = {
              enabled          = true
              file_age_timeout = 10
              interval         = 10
              restart_age_time = 10
            }
            gtp_random_sampling = {
              enabled  = true
              interval = 12
            }
            gtp_whitelist = {
              multi_whitelists = [ "example" ]
              whitelist        = "example"
            }
            health_check = {
              action          = "pass"
              dst_port        = 1
              enabled         = true
              interval        = 5
              protocol        = "icmp"
              rcv_port        = 1
              retries         = 1
              round_trip_time = 1
              src_port        = 1
            }
            hsm_group = {
              hsm_group = "example"
            }
            ip_frag = {
              forward              = true
              head_session_timeout = 15
              timeout              = 5
            }
            load_balance = {
              failover = {
                enabled               = true
                threshold_lt_bw       = 50
                threshold_lt_pkt_rate = 500
              }
              link_weight_type = "speed"
              replicate_gtpc   = true
            }
            netflow = {
              monitor = "example"
            }
            node_role = {
              mob5_g_limit     = 1
              mob_lte_limit    = 1
              stand_alone_mode = true
              type             = "control"
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
                packet_buffer = 20
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
              log_level          = "err"
              remote_syslog_ip   = "example"
              remote_syslog_port = 0
            }
            sffp_profiles = {
              sffp_profiles = [ "example" ]
            }
            sip_media = {
              timeout = 30
            }
            sip_ports = {
              ports = [ 1 ]
            }
            sip_session = {
              timeout = 30
            }
            sip_tcp_idle_timeout = {
              time = 20
            }
            sip_whitelist = {
              whitelist = "example"
            }
            ssl_decrypt = {
              decrypt_fail_action = "drop"
              enabled             = true
              hsm_pkcs11 = {
                debug_level    = 0
                dynamic_object = true
                load_sharing   = true
              }
              hsm_timeout             = 2
              key_cache_timeout       = 1
              key_map                 = "example"
              non_ssl_traffic         = "drop"
              pending_session_timeout = 30
              session_timeout         = 30
              tcp_syn_timeout         = 20
              ticket_cache_timeout    = 1
            }
            xpkt_match = {
              enabled = true
            }
          }
          template_id = "example"
        }]
      }
      operation = "add"
      rule_id   = "example"
    }]
    name      = "example"
    policy_id = "example"
    priority  = "example"
    rule_criteria = [{
      matches = {
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
      }
      operation        = "add"
      rule_criteria_id = "example"
      rule_id          = "example"
    }]
    rules_info = [{
      operation = "add"
      rule = {
        comment = "example"
        gs_operations = {
          gs_apps = {
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
          gs_engines = [ "example" ]
          gs_group_params = {
            app_tcp = {
              application  = "broadcast"
              load_balance = true
              tcp_control  = "broadcast"
            }
            dedup = {
              action    = "count"
              ip_tclass = "include"
              ip_tos    = "include"
              tcp_seq   = "include"
              timer     = 10
              vlan      = "include"
            }
            diameter_packet = {
              timeout = 1
            }
            diameter_s6_a_session = {
              limit   = 1
              timeout = 30
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
              timestamp_format = "gigasmart"
            }
            flow_mask = {
              enabled = true
              length  = 1
              offset  = 0
            }
            flow_sampling = {
              ip_ranges = [ "example" ]
              rate      = 5
              timeout   = 1
              type      = "deviceIp"
            }
            generic_session_timeout = {
              time = 5
            }
            gpfcp_profiles = {
              g_pfcp_profiles = [ "example" ]
            }
            gs_group_system = {
              cpu_load_alarm_threshold = 20
            }
            gta_profiles = {
              gta_profiles = [ "example" ]
            }
            gtp_control_sampling = {
              enabled = true
            }
            gtp_flow = {
              timeout = 1
            }
            gtp_gpfcp_delay = {
              timeout = 0
            }
            gtp_persistence = {
              enabled          = true
              file_age_timeout = 10
              interval         = 10
              restart_age_time = 10
            }
            gtp_random_sampling = {
              enabled  = true
              interval = 12
            }
            gtp_whitelist = {
              multi_whitelists = [ "example" ]
              whitelist        = "example"
            }
            health_check = {
              action          = "pass"
              dst_port        = 1
              enabled         = true
              interval        = 5
              protocol        = "icmp"
              rcv_port        = 1
              retries         = 1
              round_trip_time = 1
              src_port        = 1
            }
            hsm_group = {
              hsm_group = "example"
            }
            ip_frag = {
              forward              = true
              head_session_timeout = 15
              timeout              = 5
            }
            load_balance = {
              failover = {
                enabled               = true
                threshold_lt_bw       = 50
                threshold_lt_pkt_rate = 500
              }
              link_weight_type = "speed"
              replicate_gtpc   = true
            }
            netflow = {
              monitor = "example"
            }
            node_role = {
              mob5_g_limit     = 1
              mob_lte_limit    = 1
              stand_alone_mode = true
              type             = "control"
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
                packet_buffer = 20
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
              log_level          = "err"
              remote_syslog_ip   = "example"
              remote_syslog_port = 0
            }
            sffp_profiles = {
              sffp_profiles = [ "example" ]
            }
            sip_media = {
              timeout = 30
            }
            sip_ports = {
              ports = [ 1 ]
            }
            sip_session = {
              timeout = 30
            }
            sip_tcp_idle_timeout = {
              time = 20
            }
            sip_whitelist = {
              whitelist = "example"
            }
            ssl_decrypt = {
              decrypt_fail_action = "drop"
              enabled             = true
              hsm_pkcs11 = {
                debug_level    = 0
                dynamic_object = true
                load_sharing   = true
              }
              hsm_timeout             = 2
              key_cache_timeout       = 1
              key_map                 = "example"
              non_ssl_traffic         = "drop"
              pending_session_timeout = 30
              session_timeout         = 30
              tcp_syn_timeout         = 20
              ticket_cache_timeout    = 1
            }
            xpkt_match = {
              enabled = true
            }
          }
          templates = [{
            gs_apps = {
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
            gs_group_params = {
              app_tcp = {
                application  = "broadcast"
                load_balance = true
                tcp_control  = "broadcast"
              }
              dedup = {
                action    = "count"
                ip_tclass = "include"
                ip_tos    = "include"
                tcp_seq   = "include"
                timer     = 10
                vlan      = "include"
              }
              diameter_packet = {
                timeout = 1
              }
              diameter_s6_a_session = {
                limit   = 1
                timeout = 30
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
                timestamp_format = "gigasmart"
              }
              flow_mask = {
                enabled = true
                length  = 1
                offset  = 0
              }
              flow_sampling = {
                ip_ranges = [ "example" ]
                rate      = 5
                timeout   = 1
                type      = "deviceIp"
              }
              generic_session_timeout = {
                time = 5
              }
              gpfcp_profiles = {
                g_pfcp_profiles = [ "example" ]
              }
              gs_group_system = {
                cpu_load_alarm_threshold = 20
              }
              gta_profiles = {
                gta_profiles = [ "example" ]
              }
              gtp_control_sampling = {
                enabled = true
              }
              gtp_flow = {
                timeout = 1
              }
              gtp_gpfcp_delay = {
                timeout = 0
              }
              gtp_persistence = {
                enabled          = true
                file_age_timeout = 10
                interval         = 10
                restart_age_time = 10
              }
              gtp_random_sampling = {
                enabled  = true
                interval = 12
              }
              gtp_whitelist = {
                multi_whitelists = [ "example" ]
                whitelist        = "example"
              }
              health_check = {
                action          = "pass"
                dst_port        = 1
                enabled         = true
                interval        = 5
                protocol        = "icmp"
                rcv_port        = 1
                retries         = 1
                round_trip_time = 1
                src_port        = 1
              }
              hsm_group = {
                hsm_group = "example"
              }
              ip_frag = {
                forward              = true
                head_session_timeout = 15
                timeout              = 5
              }
              load_balance = {
                failover = {
                  enabled               = true
                  threshold_lt_bw       = 50
                  threshold_lt_pkt_rate = 500
                }
                link_weight_type = "speed"
                replicate_gtpc   = true
              }
              netflow = {
                monitor = "example"
              }
              node_role = {
                mob5_g_limit     = 1
                mob_lte_limit    = 1
                stand_alone_mode = true
                type             = "control"
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
                  packet_buffer = 20
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
                log_level          = "err"
                remote_syslog_ip   = "example"
                remote_syslog_port = 0
              }
              sffp_profiles = {
                sffp_profiles = [ "example" ]
              }
              sip_media = {
                timeout = 30
              }
              sip_ports = {
                ports = [ 1 ]
              }
              sip_session = {
                timeout = 30
              }
              sip_tcp_idle_timeout = {
                time = 20
              }
              sip_whitelist = {
                whitelist = "example"
              }
              ssl_decrypt = {
                decrypt_fail_action = "drop"
                enabled             = true
                hsm_pkcs11 = {
                  debug_level    = 0
                  dynamic_object = true
                  load_sharing   = true
                }
                hsm_timeout             = 2
                key_cache_timeout       = 1
                key_map                 = "example"
                non_ssl_traffic         = "drop"
                pending_session_timeout = 30
                session_timeout         = 30
                tcp_syn_timeout         = 20
                ticket_cache_timeout    = 1
              }
              xpkt_match = {
                enabled = true
              }
            }
            template_id = "example"
          }]
        }
        health_state = "green"
        health_state_reasons = [{
          message                               = "example"
          severity                              = "green"
          traffic_health_state_computation_type = "PORT_LOW_UTIL"
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
      }
      rule_id   = "example"
      rule_name = "example"
    }]
    src_ports = [{
      operation    = "add"
      port_ids     = [ "example" ]
      template_ids = [ "example" ]
    }]
    tags = [{
      tag_key    = "example"
      tag_values = [ "example" ]
    }]
  }
}
