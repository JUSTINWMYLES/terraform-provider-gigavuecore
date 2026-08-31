resource "gigavuecore_solution" "example" {
  app_export_config = {
    cache_config = {
      advance_hash     = true
      alias            = "example"
      description      = "example"
      dpi_inject_limit = 0
      event            = "txnEnd"
      exporters        = [ "example" ]
      flow_behavior    = "unidir"
      match = {
        datalink = {
          mac_dst = true
          mac_src = true
          vlan    = true
        }
        interface = {
          in_name_width     = 1
          in_physical_width = 2
        }
        ip = {
          version = true
        }
        ipv4 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp = true
          fragmentation = {
            flags  = true
            offset = true
          }
          header_len = true
          option_map = true
          precedence = true
          protocol   = true
          section = {
            header_size  = 1
            payload_size = 1
          }
          source = {
            prefix_min_mask = "example"
          }
          tos          = true
          total_length = true
          ttl          = true
        }
        ipv6 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp          = true
          extension_map = true
          flow_label    = true
          fragmentation = {
            flags  = true
            offset = true
          }
          hop_limit = true
          length = {
            header  = true
            payload = true
            total   = true
          }
          next_header = true
          precedence  = true
          section = {
            header_size  = 1
            payload_size = 1
          }
          source = {
            prefix_min_mask = "example"
          }
          traffic_class = true
        }
        transport = {
          dst_port = true
          icmp = {
            ipv4_code = true
            ipv4_type = true
            ipv6_code = true
            ipv6_type = true
          }
          src_port = true
          tcp = {
            ack_number  = true
            dst_port    = true
            flags       = true
            header_len  = true
            seq_number  = true
            src_port    = true
            urgent_ptr  = true
            window_size = true
          }
          udp = {
            dst_port = true
            msg_len  = true
            src_port = true
          }
        }
      }
      multi_collect         = true
      network_profiles      = [ "example" ]
      observation_domain_id = 0
      sampling = {
        mode                 = "multiRate"
        single_sampling_rate = 10
      }
      size = {
        flows = 1
      }
      timeout = {
        idle = 1
      }
    }
    cache_config_alias = "example"
    destination_configs = [{
      application_names = [{
        attributes = [{
          name  = "example"
          value = "example"
        }]
        is_user_defined = true
        name            = "example"
      }]
      destination_name    = "example"
      export_ip_interface = "example"
      export_meta_app_profile = {
        alias          = "example"
        application_id = true
        applications = [{
          attributes = [{
            name  = "example"
            value = "example"
          }]
          is_user_defined = true
          name            = "example"
        }]
        counter = {
          bytes           = true
          bytes_long      = true
          inner_byte      = true
          inner_byte_long = true
          packets         = true
          packets_long    = true
        }
        datalink = {
          mac_dst = true
          mac_src = true
          vlan    = true
        }
        description = "example"
        flow = {
          end_reason = true
        }
        gtpu = {
          qfi  = true
          teid = true
        }
        interface = {
          in_name_width      = 1
          in_physical_width  = 2
          out_physical_width = 2
        }
        ip = {
          version = true
        }
        ipv4 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp = true
          fragmentation = {
            flags  = true
            offset = true
          }
          header_len = true
          option_map = true
          precedence = true
          protocol   = true
          section = {
            header_size  = 1
            payload_size = 1
          }
          source = {
            prefix_min_mask = "example"
          }
          tos          = true
          total_length = true
          ttl          = true
        }
        ipv6 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp          = true
          extension_map = true
          flow_label    = true
          fragmentation = {
            flags  = true
            offset = true
          }
          hop_limit = true
          length = {
            header  = true
            payload = true
            total   = true
          }
          next_header = true
          precedence  = true
          section = {
            header_size  = 1
            payload_size = 1
          }
          source = {
            prefix_min_mask = "example"
          }
          traffic_class = true
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
          flow_end_msec     = true
          flow_endsec       = true
          flow_start_msec   = true
          flow_startsec     = true
          sys_up_time_first = true
          sys_up_time_last  = true
        }
        transport = {
          dst_port = true
          icmp = {
            ipv4_code = true
            ipv4_type = true
            ipv6_code = true
            ipv6_type = true
          }
          src_port = true
          tcp = {
            ack_number  = true
            dst_port    = true
            flags       = true
            header_len  = true
            seq_number  = true
            src_port    = true
            urgent_ptr  = true
            window_size = true
          }
          udp = {
            dst_port = true
            msg_len  = true
            src_port = true
          }
        }
        type = "export"
      }
      export_meta_app_profile_alias = "example"
      exporter_alias                = "example"
      exporter_config = {
        alias                = "example"
        application_profiles = [ "example" ]
        cef = {
          active_timeout   = 1
          inactive_timeout = 1
        }
        description = "example"
        destination = {
          dscp         = 0
          ipv4_address = "example"
          l4_port_dst  = 1
          l4_port_src  = 1
          l4_protocol  = "udp"
          ttl          = 1
        }
        max_pkt_size = 0
        mobility_sam = {
          encoding        = "example"
          encoding_format = "hierarchy"
          event_enable = {
            modify = true
            update = true
          }
          trigger = "example"
        }
        monitor = {
          timeout = 60
        }
        netflow = {
          active_timeout   = 1
          inactive_timeout = 1
          template_refresh = 1
          template_type    = "cohesive"
          version          = "v5"
        }
        snmp = {
          enabled = true
        }
        source = {
          ip_interface = "example"
        }
        type = "cef"
      }
    }]
    gsop_alias = "example"
    gsop_config = {
      alias      = "example"
      cluster_id = "example"
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
      gs_group     = "example"
      health_state = "green"
      health_state_reasons = [{
        message                               = "example"
        severity                              = "green"
        traffic_health_state_computation_type = "PORT_LOW_UTIL"
      }]
    }
  }
  app_filter_config = {
    egress_traffic_config = [{
      drop_app_profile_alias = "example"
      egress_traffic_map = {
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
        gsop                = "example"
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
        vlan_tag = {
          tag_protocol_id = "0x8100"
          vlan_action     = "add"
          vlan_id         = 0
        }
      }
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
      gsop_alias = "example"
      gsop_config = {
        alias      = "example"
        cluster_id = "example"
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
        gs_group     = "example"
        health_state = "green"
        health_state_reasons = [{
          message                               = "example"
          severity                              = "green"
          traffic_health_state_computation_type = "PORT_LOW_UTIL"
        }]
      }
      map_alias              = "example"
      pass_app_profile_alias = "example"
      priority               = 0
    }]
    sapf_profile = {
      alias = "example"
      bidi  = true
      buffering = {
        buffer_count_before_match = 3
        enabled                   = true
        protocol                  = "tcp"
      }
      cluster_id   = "example"
      packet_count = 0
      session_fields = [{
        pos  = 1
        type = "ipv4Addr"
      }]
      timeout = 10
    }
    sapf_profile_alias = "example"
  }
  associated_monitor_solution_alias = "example"
  cluster_id                        = "example"
  config_status                     = "example"
  delete_monitor_sol                = true
  egress_map_aliases_to_delete      = [ "example" ]
  exporter_aliases_to_delete        = [ "example" ]
  health_state                      = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  ingress_map_aliases_to_delete = [ "example" ]
  ingress_traffic_configs = [{
    ingress_traffic_map = {
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
    map_alias = "example"
  }]
  monitor_solution_config = {
    action         = "example"
    cluster_id     = "example"
    config_status  = "example"
    error_message  = "example"
    exporter_alias = "example"
    exporter_config = {
      alias                = "example"
      application_profiles = [ "example" ]
      cef = {
        active_timeout   = 1
        inactive_timeout = 1
      }
      description = "example"
      destination = {
        dscp         = 0
        ipv4_address = "example"
        l4_port_dst  = 1
        l4_port_src  = 1
        l4_protocol  = "udp"
        ttl          = 1
      }
      max_pkt_size = 0
      mobility_sam = {
        encoding        = "example"
        encoding_format = "hierarchy"
        event_enable = {
          modify = true
          update = true
        }
        trigger = "example"
      }
      monitor = {
        timeout = 60
      }
      netflow = {
        active_timeout   = 1
        inactive_timeout = 1
        template_refresh = 1
        template_type    = "cohesive"
        version          = "v5"
      }
      snmp = {
        enabled = true
      }
      source = {
        ip_interface = "example"
      }
      type = "cef"
    }
    gs_group             = "example"
    monitor_ip_interface = "example"
    ref_sols = [{
      alias          = "example"
      associated_map = "example"
      type           = "example"
    }]
    solution_alias = "example"
    vport_alias    = "example"
    vport_config = {
      alias              = "example"
      deferred_binding   = true
      fail_over_action   = "vport-bypass"
      gs_group           = "example"
      inline_status      = "up"
      inner_traffic_path = "to-inline-tool"
      metadata_monitoring = {
        action    = "enable"
        exporters = [ "example" ]
      }
      mode               = "none"
      outer_traffic_path = "to-inline-tool"
      sa_apf_profile     = "example"
    }
  }
  solution_alias  = "example"
  solution_desc   = "example"
  solution_status = "example"
  solution_type   = "monitor"
}
