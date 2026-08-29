resource "gigavuecore_solution" "example" {
  app_export_config = {
    cache_config = {
      advance_hash     = true
      alias            = "example"
      description      = "example"
      dpi_inject_limit = 0
      event            = "example"
      exporters        = [ "example" ]
      flow_behavior    = "example"
      match = {
        datalink = {
          mac_dst = true
          mac_src = true
          vlan    = true
        }
        interface = {
          in_name_width     = 0
          in_physical_width = 0
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
            header_size  = 0
            payload_size = 0
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
            header_size  = 0
            payload_size = 0
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
        mode                 = "example"
        single_sampling_rate = 0
      }
      size = {
        flows = 0
      }
      timeout = {
        idle = 0
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
          in_name_width      = 0
          in_physical_width  = 0
          out_physical_width = 0
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
            header_size  = 0
            payload_size = 0
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
            header_size  = 0
            payload_size = 0
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
        type = "example"
      }
      export_meta_app_profile_alias = "example"
      exporter_alias                = "example"
      exporter_config = {
        alias                = "example"
        application_profiles = [ "example" ]
        cef = {
          active_timeout   = 0
          inactive_timeout = 0
        }
        description = "example"
        destination = {
          dscp         = 0
          ipv4_address = "example"
          l4_port_dst  = 0
          l4_port_src  = 0
          l4_protocol  = "example"
          ttl          = 0
        }
        max_pkt_size = 0
        mobility_sam = {
          encoding        = "example"
          encoding_format = "example"
          event_enable = {
            modify = true
            update = true
          }
          trigger = "example"
        }
        monitor = {
          timeout = 0
        }
        netflow = {
          active_timeout   = 0
          inactive_timeout = 0
          template_refresh = 0
          template_type    = "example"
          version          = "example"
        }
        snmp = {
          enabled = true
        }
        source = {
          ip_interface = "example"
        }
        type = "example"
      }
    }]
    gsop_alias = "example"
    gsop_config = {
      alias      = "example"
      cluster_id = "example"
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
      gs_group     = "example"
      health_state = "example"
      health_state_reasons = [{
        message                               = "example"
        severity                              = "example"
        traffic_health_state_computation_type = "example"
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
            rule_id             = 0
          }]
          pass_rules = [{
            application_profile = "example"
            rule_id             = 0
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
        fstype = {
          offset = 0
          timer  = 0
          type   = "example"
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
        gsop                = "example"
        inline_traffic_path = "example"
        inline_traffic_type = "example"
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
        vlan_tag = {
          tag_protocol_id = "example"
          vlan_action     = "example"
          vlan_id         = 0
        }
      }
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
      gsop_alias = "example"
      gsop_config = {
        alias      = "example"
        cluster_id = "example"
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
        gs_group     = "example"
        health_state = "example"
        health_state_reasons = [{
          message                               = "example"
          severity                              = "example"
          traffic_health_state_computation_type = "example"
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
        buffer_count_before_match = 0
        enabled                   = true
        protocol                  = "example"
      }
      cluster_id   = "example"
      packet_count = 0
      session_fields = [{
        pos  = 0
        type = "example"
      }]
      timeout = 0
    }
    sapf_profile_alias = "example"
  }
  associated_monitor_solution_alias = "example"
  cluster_id                        = "example"
  config_status                     = "example"
  delete_monitor_sol                = true
  egress_map_aliases_to_delete      = [ "example" ]
  exporter_aliases_to_delete        = [ "example" ]
  health_state                      = "example"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "example"
    traffic_health_state_computation_type = "example"
  }]
  ingress_map_aliases_to_delete = [ "example" ]
  ingress_traffic_configs = [{
    ingress_traffic_map = {
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
      fstype = {
        offset = 0
        timer  = 0
        type   = "example"
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
      updated_time     = 1.0
      vlan_tag = {
        tag_protocol_id = "example"
        vlan_action     = "example"
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
        active_timeout   = 0
        inactive_timeout = 0
      }
      description = "example"
      destination = {
        dscp         = 0
        ipv4_address = "example"
        l4_port_dst  = 0
        l4_port_src  = 0
        l4_protocol  = "example"
        ttl          = 0
      }
      max_pkt_size = 0
      mobility_sam = {
        encoding        = "example"
        encoding_format = "example"
        event_enable = {
          modify = true
          update = true
        }
        trigger = "example"
      }
      monitor = {
        timeout = 0
      }
      netflow = {
        active_timeout   = 0
        inactive_timeout = 0
        template_refresh = 0
        template_type    = "example"
        version          = "example"
      }
      snmp = {
        enabled = true
      }
      source = {
        ip_interface = "example"
      }
      type = "example"
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
      fail_over_action   = "example"
      gs_group           = "example"
      inline_status      = "example"
      inner_traffic_path = "example"
      metadata_monitoring = {
        action    = "example"
        exporters = [ "example" ]
      }
      mode               = "example"
      outer_traffic_path = "example"
      sa_apf_profile     = "example"
    }
  }
  solution_alias  = "example"
  solution_desc   = "example"
  solution_status = "example"
  solution_type   = "example"
}
