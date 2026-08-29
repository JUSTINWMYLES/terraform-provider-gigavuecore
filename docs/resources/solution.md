---
page_title: "gigavuecore_solution Resource - gigavuecore"
subcategory: ""
description: |-
  Load Apps Visibility Solutions by alias
---

# gigavuecore_solution Resource

Load Apps Visibility Solutions by alias

## Example Usage

```terraform
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
  ingress_map_aliases_to_delete     = [ "example" ]
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
```

## Schema

### Arguments

The following arguments are supported:

* `app_export_config` (Attributes, optional) (see [below for nested schema](#nestedatt--app_export_config))
* `app_filter_config` (Attributes, optional) (see [below for nested schema](#nestedatt--app_filter_config))
* `associated_monitor_solution_alias` (String, optional) - monitor object alias
* `cluster_id` (String, optional) - cluster id for which solution is getting created
* `config_status` (String, optional) - configuration status
* `delete_monitor_sol` (Boolean, optional)
* `egress_map_aliases_to_delete` (List of String, optional)
* `exporter_aliases_to_delete` (List of String, optional)
* `ingress_map_aliases_to_delete` (List of String, optional)
* `ingress_traffic_configs` (Attributes List, optional) - ingress traffic configuration (see [below for nested schema](#nestedatt--ingress_traffic_configs))
* `monitor_solution_config` (Attributes, optional) (see [below for nested schema](#nestedatt--monitor_solution_config))
* `solution_alias` (String, optional) - user defined application visibility solution alias
* `solution_desc` (String, optional) - user defined application visibility solution description
* `solution_status` (String, optional) - solution status
* `solution_type` (String, optional) - user intent solution type

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `app_export_config` (Attributes, computed) (see [below for nested schema](#nestedatt--app_export_config))
* `app_filter_config` (Attributes, computed) (see [below for nested schema](#nestedatt--app_filter_config))
* `associated_monitor_solution_alias` (String, computed) - monitor object alias
* `cluster_id` (String, computed) - cluster id for which solution is getting created
* `config_status` (String, computed) - configuration status
* `delete_monitor_sol` (Boolean, computed)
* `egress_map_aliases_to_delete` (List of String, computed)
* `exporter_aliases_to_delete` (List of String, computed)
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `ingress_map_aliases_to_delete` (List of String, computed)
* `ingress_traffic_configs` (Attributes List, computed) - ingress traffic configuration (see [below for nested schema](#nestedatt--ingress_traffic_configs))
* `monitor_solution_config` (Attributes, computed) (see [below for nested schema](#nestedatt--monitor_solution_config))
* `solution_alias` (String, computed) - user defined application visibility solution alias
* `solution_desc` (String, computed) - user defined application visibility solution description
* `solution_status` (String, computed) - solution status
* `solution_type` (String, computed) - user intent solution type

<a id="nestedatt--app_export_config"></a>
### Nested Schema for `app_export_config`

Optional:

* `cache_config` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config))
* `cache_config_alias` (String)
* `destination_configs` (Attributes List) (see [below for nested schema](#nestedatt--app_export_config--destination_configs))
* `gsop_alias` (String)
* `gsop_config` (Attributes) - GigaSMART Operation (see [below for nested schema](#nestedatt--app_export_config--gsop_config))
<a id="nestedatt--app_export_config--cache_config"></a>
### Nested Schema for `app_export_config.cache_config`

Required:

* `alias` (String)
Optional:

* `advance_hash` (Boolean) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
* `description` (String)
* `dpi_inject_limit` (Number)
* `event` (String)
* `exporters` (List of String) - alias of metadata exporters to attach this cache
* `flow_behavior` (String) - direction for flow identification
* `match` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match))
* `multi_collect` (Boolean) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
* `network_profiles` (List of String) - alias of metadata network profiles to attach this cache
* `observation_domain_id` (Number)
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--sampling))
* `size` (Attributes) - size of the flows (see [below for nested schema](#nestedatt--app_export_config--cache_config--size))
* `timeout` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--timeout))
<a id="nestedatt--app_export_config--cache_config--match"></a>
### Nested Schema for `app_export_config.cache_config.match`

Optional:

* `datalink` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--datalink))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--transport))
<a id="nestedatt--app_export_config--cache_config--match--datalink"></a>
### Nested Schema for `app_export_config.cache_config.match.datalink`

Optional:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--interface"></a>
### Nested Schema for `app_export_config.cache_config.match.interface`

Optional:

* `in_name_width` (Number)
* `in_physical_width` (Number)
<a id="nestedatt--app_export_config--cache_config--match--ip"></a>
### Nested Schema for `app_export_config.cache_config.match.ip`

Optional:

* `version` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv4"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv4--destination"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4.destination`

Optional:

* `prefix_min_mask` (String) - ipv4 destination prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_export_config--cache_config--match--ipv4--fragmentation"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv4--section"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_export_config--cache_config--match--ipv4--source"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4.source`

Optional:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_export_config--cache_config--match--ipv6"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--destination"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--fragmentation"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--length"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.length`

Optional:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--section"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--source"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.source`

Optional:

* `prefix_min_mask` (String) - ipv6 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_export_config--cache_config--match--transport"></a>
### Nested Schema for `app_export_config.cache_config.match.transport`

Optional:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--transport--udp))
<a id="nestedatt--app_export_config--cache_config--match--transport--icmp"></a>
### Nested Schema for `app_export_config.cache_config.match.transport.icmp`

Optional:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--transport--tcp"></a>
### Nested Schema for `app_export_config.cache_config.match.transport.tcp`

Optional:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--transport--udp"></a>
### Nested Schema for `app_export_config.cache_config.match.transport.udp`

Optional:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--app_export_config--cache_config--sampling"></a>
### Nested Schema for `app_export_config.cache_config.sampling`

Optional:

* `mode` (String)
* `single_sampling_rate` (Number) - Packet interval window size. Valid values: 10-16000 (in packets)
<a id="nestedatt--app_export_config--cache_config--size"></a>
### Nested Schema for `app_export_config.cache_config.size`

Optional:

* `flows` (Number) - size of flows in millions
<a id="nestedatt--app_export_config--cache_config--timeout"></a>
### Nested Schema for `app_export_config.cache_config.timeout`

Optional:

* `idle` (Number) - idle timeout in seconds. max value 7days. default 30 min
<a id="nestedatt--app_export_config--destination_configs"></a>
### Nested Schema for `app_export_config.destination_configs`

Optional:

* `application_names` (Attributes List) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--application_names))
* `destination_name` (String)
* `export_ip_interface` (String)
* `export_meta_app_profile` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile))
* `export_meta_app_profile_alias` (String)
* `exporter_alias` (String)
* `exporter_config` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config))
<a id="nestedatt--app_export_config--destination_configs--application_names"></a>
### Nested Schema for `app_export_config.destination_configs.application_names`

Required:

* `name` (String) - application name
Optional:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
<a id="nestedatt--app_export_config--destination_configs--application_names--attributes"></a>
### Nested Schema for `app_export_config.destination_configs.application_names.attributes`

Required:

* `name` (String) - attribute name
Optional:

* `value` (String) - application's attribute value
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile`

Required:

* `alias` (String) - application profile alias
Optional:

* `application_id` (Boolean) - only valid with 'export' type
* `applications` (Attributes List) - application and attributes. (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--counter))
* `datalink` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--datalink))
* `description` (String)
* `flow` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--gtpu))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport))
* `type` (String)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--applications"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.applications`

Required:

* `name` (String) - application name
Optional:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--applications--attributes"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.applications.attributes`

Required:

* `name` (String) - attribute name
Optional:

* `value` (String) - application's attribute value
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--counter"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.counter`

Optional:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `inner_byte` (Boolean)
* `inner_byte_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--datalink"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.datalink`

Optional:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--flow"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.flow`

Optional:

* `end_reason` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--gtpu"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.gtpu`

Optional:

* `qfi` (Boolean)
* `teid` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--interface"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.interface`

Optional:

* `in_name_width` (Number)
* `in_physical_width` (Number)
* `out_physical_width` (Number)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ip"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ip`

Optional:

* `version` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--destination"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--fragmentation"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--section"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--source"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4.source`

Optional:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--destination"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--fragmentation"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--length"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.length`

Optional:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--section"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--source"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.source`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--outer_ipv4"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.outer_ipv4`

Optional:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--outer_ipv6"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.outer_ipv6`

Optional:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--timestamp"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.timestamp`

Optional:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
* `sys_up_time_first` (Boolean)
* `sys_up_time_last` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.transport`

Optional:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--udp))
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--icmp"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.transport.icmp`

Optional:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--tcp"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.transport.tcp`

Optional:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--udp"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.transport.udp`

Optional:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--exporter_config"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config`

Required:

* `alias` (String)
Optional:

* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--cef))
* `description` (String)
* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--mobility_sam))
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--monitor))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--snmp))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--source))
* `type` (String)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--cef"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.cef`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
<a id="nestedatt--app_export_config--destination_configs--exporter_config--destination"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.destination`

Optional:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--mobility_sam"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.mobility_sam`

Optional:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--mobility_sam--event_enable))
* `trigger` (String)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--mobility_sam--event_enable"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.mobility_sam.event_enable`

Optional:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--monitor"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.monitor`

Optional:

* `timeout` (Number) - how often to export in seconds
<a id="nestedatt--app_export_config--destination_configs--exporter_config--netflow"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.netflow`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--snmp"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.snmp`

Optional:

* `enabled` (Boolean) - snmp reverse lookup enable/disable
<a id="nestedatt--app_export_config--destination_configs--exporter_config--source"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.source`

Optional:

* `ip_interface` (String)
<a id="nestedatt--app_export_config--gsop_config"></a>
### Nested Schema for `app_export_config.gsop_config`

Required:

* `alias` (String)
* `gs_apps` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps))
* `gs_group` (String) - Alias of referenced managing GsGroup
Optional:

* `cluster_id` (String) - id of the defining cluster
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--health_state_reasons))
<a id="nestedatt--app_export_config--gsop_config--gs_apps"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps`

Optional:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap))
<a id="nestedatt--app_export_config--gsop_config--gs_apps--apf"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.apf`

Optional:

* `enabled` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--dedup"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.dedup`

Optional:

* `enabled` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--diameter_whitelist"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.diameter_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--flow_filter"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.flow_filter`

Required:

* `type` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--flow_sampling"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.flow_sampling`

Required:

* `type` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--gseries_header_add"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.gseries_header_add`

Required:

* `types` (Set of String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--gseries_header_remove"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.gseries_header_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--gseries_load_balance"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.gseries_load_balance`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--gseries_load_balance--variable_offset))
<a id="nestedatt--app_export_config--gsop_config--gs_apps--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.gseries_load_balance.fixed_offset`

Required:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--gseries_load_balance--variable_offset"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.gseries_load_balance.variable_offset`

Required:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--gseries_pattern_match"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.gseries_pattern_match`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--gseries_pattern_match--variable_offset))
<a id="nestedatt--app_export_config--gsop_config--gs_apps--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.gseries_pattern_match.fixed_offset`

Required:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.gseries_pattern_match.variable_offset`

Required:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--gtp_whitelist"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.gtp_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--header_add"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.header_add`

Required:

* `vlan` (Number)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--header_remove"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.header_remove`

Required:

* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series
Optional:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
<a id="nestedatt--app_export_config--gsop_config--gs_apps--icap"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.icap`

Required:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--app_export_config--gsop_config--gs_apps--inline_ssl"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.inline_ssl`

Required:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--app_export_config--gsop_config--gs_apps--load_balance"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.load_balance`

Optional:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--load_balance--stateless))
<a id="nestedatt--app_export_config--gsop_config--gs_apps--load_balance--enhanced"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.load_balance.enhanced`

Required:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--app_export_config--gsop_config--gs_apps--load_balance--stateful"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.load_balance.stateful`

Required:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
Optional:

* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--app_export_config--gsop_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.load_balance.stateful.diameter_key_multi_hash_type`

Required:

* `key` (String)
Optional:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
<a id="nestedatt--app_export_config--gsop_config--gs_apps--load_balance--stateless"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.load_balance.stateless`

Required:

* `hash_fields` (String)
Optional:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
<a id="nestedatt--app_export_config--gsop_config--gs_apps--masking"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.masking`

Required:

* `protocol` (String)
Optional:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
<a id="nestedatt--app_export_config--gsop_config--gs_apps--metadata_export"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.metadata_export`

Optional:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--app_export_config--gsop_config--gs_apps--netflow"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.netflow`

Optional:

* `enabled` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--sa_apf"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.sa_apf`

Optional:

* `enabled` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--sip_whitelist"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.sip_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--slicing"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.slicing`

Required:

* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
Optional:

* `enhanced` (String) - enhanced-slicing apps alias
<a id="nestedatt--app_export_config--gsop_config--gs_apps--ssl_decrypt"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.ssl_decrypt`

Optional:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--app_export_config--gsop_config--gs_apps--trailer_add"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.trailer_add`

Required:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--app_export_config--gsop_config--gs_apps--trailer_remove"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.trailer_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--tunnel_decap"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.tunnel_decap`

Required:

* `type` (String)
Optional:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--tunnel_decap--tls_pcapng))
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--tunnel_decap--vxlan))
<a id="nestedatt--app_export_config--gsop_config--gs_apps--tunnel_decap--custom"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.tunnel_decap.custom`

Required:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--app_export_config--gsop_config--gs_apps--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.tunnel_decap.tls_pcapng`

Optional:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--tunnel_decap--vxlan"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.tunnel_decap.vxlan`

Required:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.tunnel_encap`

Required:

* `type` (String)
Optional:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap--tls_pcapng))
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap--vxlan_config))
<a id="nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap--gmip_config"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.tunnel_encap.gmip_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
Optional:

* `dscp` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `ttl` (Number)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.tunnel_encap.l2_gre_config`

Required:

* `key` (Number)
Optional:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.tunnel_encap.tls_pcapng`

Optional:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--app_export_config--gsop_config--gs_apps--tunnel_encap--vxlan_config"></a>
### Nested Schema for `app_export_config.gsop_config.gs_apps.tunnel_encap.vxlan_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `vni` (Number)
Optional:

* `dscp` (Number)
* `ttl` (Number)
<a id="nestedatt--app_export_config--gsop_config--health_state_reasons"></a>
### Nested Schema for `app_export_config.gsop_config.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--app_filter_config"></a>
### Nested Schema for `app_filter_config`

Optional:

* `egress_traffic_config` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config))
* `sapf_profile` (Attributes) - Session-Aware APF (see [below for nested schema](#nestedatt--app_filter_config--sapf_profile))
* `sapf_profile_alias` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config"></a>
### Nested Schema for `app_filter_config.egress_traffic_config`

Optional:

* `drop_app_profile_alias` (String)
* `egress_traffic_map` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map))
* `gs_apps` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps))
* `gsop_alias` (String)
* `gsop_config` (Attributes) - GigaSMART Operation (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config))
* `map_alias` (String)
* `pass_app_profile_alias` (String)
* `priority` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map`

Required:

* `alias` (String) - unique map alias
* `dst_ports` (List of String) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
* `src_ports` (List of String) - list of the 'from' ports. Only port number is supported
* `type` (String) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
Optional:

* `ap_rules` (Attributes) - pass and drop application profile rules (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules))
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `egress_gigastream` (List of String) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Boolean) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String) - tunnel alias
* `flex_inline` (Attributes) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline))
* `flex_inline_failover` (String) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Attributes) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules))
* `fstype` (Attributes) - Type of Mobility Flowsampling & properties (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--fstype))
* `gs_rules` (Attributes) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules))
* `gsop` (String) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--ip_rewrite))
* `mod_time` (Number) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Boolean) - enabled when the dstPort is null
* `order` (Number) - relative order within per-source port map chain
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rewrite))
* `roles` (Attributes) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--roles))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules))
* `rx_cluster_ports` (List of String)
* `sub_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List of String)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--vlan_tag))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.ap_rules`

Optional:

* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules--drop_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.ap_rules.drop_rules`

Required:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.ap_rules.pass_rules`

Required:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline`

Optional:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--tag))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--a_to_b"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.a_to_b`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--b_to_a"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.b_to_a`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--oob_copy"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.oob_copy`

Required:

* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists
Optional:

* `direction` (String)
* `tag` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--oob_copy--tag))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--oob_copy--tag"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.oob_copy.tag`

Required:

* `type` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--tag"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.tag`

Required:

* `type` (String)
Optional:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `vlan_id` (Number) - only applicable when type is 'vlan'
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--drop_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules.drop_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--drop_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules.drop_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--pass_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules.pass_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_overlap_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_overlap_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
* `priority` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_diameter_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_diameter_rules.pass_rules`

Required:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_diameter_rules.pass_rules.diameter`

Optional:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_overlap_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_overlap_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules`

Required:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules.sip`

Optional:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--id_range))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Required:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Required:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules.sip.id_range`

Required:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist5_g_overlap_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist5_g_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist5_g_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules.pass_rules`

Required:

* `rule_id` (Number)
Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--sip))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules.pass_rules`

Required:

* `rule_id` (Number)
Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--sip))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--fstype"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.fstype`

Optional:

* `offset` (Number) - Offset for Mobility Rotational Flowsampling
* `timer` (Number) - Timer for Mobility Rotational Flowsampling in minutes
* `type` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.gs_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules--drop_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.gs_rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.gs_rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--ip_rewrite"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rewrite"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--roles"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.roles`

Optional:

* `editors` (List of String)
* `listeners` (List of String)
* `owners` (List of String)
* `viewers` (List of String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--vlan_tag))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rules.drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--rewrite"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rules.drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rules.drop_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--vlan_tag))
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rules.pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--rewrite"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rules.pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.rules.pass_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--egress_traffic_map--vlan_tag"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.egress_traffic_map.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps`

Optional:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap))
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--apf"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.apf`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--dedup"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.dedup`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--diameter_whitelist"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.diameter_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--flow_filter"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.flow_filter`

Required:

* `type` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--flow_sampling"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.flow_sampling`

Required:

* `type` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_header_add"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.gseries_header_add`

Required:

* `types` (Set of String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_header_remove"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.gseries_header_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.gseries_load_balance`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance--variable_offset))
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.gseries_load_balance.fixed_offset`

Required:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance--variable_offset"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.gseries_load_balance.variable_offset`

Required:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.gseries_pattern_match`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match--variable_offset))
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.gseries_pattern_match.fixed_offset`

Required:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.gseries_pattern_match.variable_offset`

Required:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--gtp_whitelist"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.gtp_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--header_add"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.header_add`

Required:

* `vlan` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--header_remove"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.header_remove`

Required:

* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series
Optional:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--icap"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.icap`

Required:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--inline_ssl"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.inline_ssl`

Required:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.load_balance`

Optional:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateless))
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance--enhanced"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.load_balance.enhanced`

Required:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateful"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.load_balance.stateful`

Required:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
Optional:

* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.load_balance.stateful.diameter_key_multi_hash_type`

Required:

* `key` (String)
Optional:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateless"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.load_balance.stateless`

Required:

* `hash_fields` (String)
Optional:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--masking"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.masking`

Required:

* `protocol` (String)
Optional:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--metadata_export"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.metadata_export`

Optional:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--netflow"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.netflow`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--sa_apf"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.sa_apf`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--sip_whitelist"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.sip_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--slicing"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.slicing`

Required:

* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
Optional:

* `enhanced` (String) - enhanced-slicing apps alias
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--ssl_decrypt"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.ssl_decrypt`

Optional:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--trailer_add"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.trailer_add`

Required:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--trailer_remove"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.trailer_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.tunnel_decap`

Required:

* `type` (String)
Optional:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--tls_pcapng))
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--vxlan))
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--custom"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.tunnel_decap.custom`

Required:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.tunnel_decap.tls_pcapng`

Optional:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--vxlan"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.tunnel_decap.vxlan`

Required:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.tunnel_encap`

Required:

* `type` (String)
Optional:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--tls_pcapng))
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--vxlan_config))
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--gmip_config"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.tunnel_encap.gmip_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
Optional:

* `dscp` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `ttl` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.tunnel_encap.l2_gre_config`

Required:

* `key` (Number)
Optional:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.tunnel_encap.tls_pcapng`

Optional:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--vxlan_config"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gs_apps.tunnel_encap.vxlan_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `vni` (Number)
Optional:

* `dscp` (Number)
* `ttl` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config`

Required:

* `alias` (String)
* `gs_apps` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps))
* `gs_group` (String) - Alias of referenced managing GsGroup
Optional:

* `cluster_id` (String) - id of the defining cluster
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--health_state_reasons))
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps`

Optional:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap))
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--apf"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.apf`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--dedup"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.dedup`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--diameter_whitelist"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.diameter_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--flow_filter"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.flow_filter`

Required:

* `type` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--flow_sampling"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.flow_sampling`

Required:

* `type` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_header_add"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_header_add`

Required:

* `types` (Set of String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_header_remove"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_header_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_load_balance`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance--variable_offset))
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_load_balance.fixed_offset`

Required:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance--variable_offset"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_load_balance.variable_offset`

Required:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_pattern_match`

Optional:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match--variable_offset))
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_pattern_match.fixed_offset`

Required:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_pattern_match.variable_offset`

Required:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gtp_whitelist"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.gtp_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--header_add"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.header_add`

Required:

* `vlan` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--header_remove"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.header_remove`

Required:

* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series
Optional:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--icap"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.icap`

Required:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--inline_ssl"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.inline_ssl`

Required:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance`

Optional:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateless))
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--enhanced"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance.enhanced`

Required:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateful"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance.stateful`

Required:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
Optional:

* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance.stateful.diameter_key_multi_hash_type`

Required:

* `key` (String)
Optional:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateless"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance.stateless`

Required:

* `hash_fields` (String)
Optional:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--masking"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.masking`

Required:

* `protocol` (String)
Optional:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--metadata_export"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.metadata_export`

Optional:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--netflow"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.netflow`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--sa_apf"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.sa_apf`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--sip_whitelist"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.sip_whitelist`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--slicing"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.slicing`

Required:

* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
Optional:

* `enhanced` (String) - enhanced-slicing apps alias
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--ssl_decrypt"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.ssl_decrypt`

Optional:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--trailer_add"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.trailer_add`

Required:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--trailer_remove"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.trailer_remove`

Optional:

* `enabled` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_decap`

Required:

* `type` (String)
Optional:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--tls_pcapng))
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--vxlan))
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--custom"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_decap.custom`

Required:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_decap.tls_pcapng`

Optional:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--vxlan"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_decap.vxlan`

Required:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap`

Required:

* `type` (String)
Optional:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--tls_pcapng))
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--vxlan_config))
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--gmip_config"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap.gmip_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
Optional:

* `dscp` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `ttl` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap.l2_gre_config`

Required:

* `key` (Number)
Optional:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap.tls_pcapng`

Optional:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--vxlan_config"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap.vxlan_config`

Required:

* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `vni` (Number)
Optional:

* `dscp` (Number)
* `ttl` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_config--gsop_config--health_state_reasons"></a>
### Nested Schema for `app_filter_config.egress_traffic_config.gsop_config.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--app_filter_config--sapf_profile"></a>
### Nested Schema for `app_filter_config.sapf_profile`

Required:

* `alias` (String)
* `session_fields` (Attributes Set) - A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20 (see [below for nested schema](#nestedatt--app_filter_config--sapf_profile--session_fields))
Optional:

* `bidi` (Boolean) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
* `buffering` (Attributes) - Session-Aware APF Buffereing settings (see [below for nested schema](#nestedatt--app_filter_config--sapf_profile--buffering))
* `cluster_id` (String) - id of the defining cluster
* `packet_count` (Number) - For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.
* `timeout` (Number) - in seconds
<a id="nestedatt--app_filter_config--sapf_profile--session_fields"></a>
### Nested Schema for `app_filter_config.sapf_profile.session_fields`

Required:

* `type` (String)
Optional:

* `pos` (Number) - Value of 1 also an alias for 'outer'. Value of 2 also an alias for 'inner'. Not applicable for 'gtpuTeid'. For 'fiveTuple' 'outer' is not supported
<a id="nestedatt--app_filter_config--sapf_profile--buffering"></a>
### Nested Schema for `app_filter_config.sapf_profile.buffering`

Required:

* `enabled` (Boolean)
Optional:

* `buffer_count_before_match` (Number) - Maximum number of packets BSAPF will buffer per session before APF match
* `protocol` (String) - When buffering is enabled, changing protocol requires reboot before new value takes effect.
<a id="nestedatt--ingress_traffic_configs"></a>
### Nested Schema for `ingress_traffic_configs`

Optional:

* `ingress_traffic_map` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map))
* `map_alias` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map`

Required:

* `alias` (String) - unique map alias
* `dst_ports` (List of String) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
* `src_ports` (List of String) - list of the 'from' ports. Only port number is supported
* `type` (String) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
Optional:

* `ap_rules` (Attributes) - pass and drop application profile rules (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--ap_rules))
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `egress_gigastream` (List of String) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Boolean) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String) - tunnel alias
* `flex_inline` (Attributes) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline))
* `flex_inline_failover` (String) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Attributes) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules))
* `fstype` (Attributes) - Type of Mobility Flowsampling & properties (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--fstype))
* `gs_rules` (Attributes) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--gs_rules))
* `gsop` (String) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--health_state_reasons))
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--ip_rewrite))
* `mod_time` (Number) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Boolean) - enabled when the dstPort is null
* `order` (Number) - relative order within per-source port map chain
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rewrite))
* `roles` (Attributes) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--roles))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rules))
* `rx_cluster_ports` (List of String)
* `sub_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List of String)
* `updated_time` (Number) - Last Updated time of the fabric map
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--vlan_tag))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--ap_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.ap_rules`

Optional:

* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--ap_rules--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--ap_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--ap_rules--drop_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.ap_rules.drop_rules`

Required:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--ap_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.ap_rules.pass_rules`

Required:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flex_inline`

Optional:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--tag))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--a_to_b"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flex_inline.a_to_b`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--b_to_a"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flex_inline.b_to_a`

Required:

* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
Optional:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--oob_copy"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flex_inline.oob_copy`

Required:

* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists
Optional:

* `direction` (String)
* `tag` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--oob_copy--tag))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--oob_copy--tag"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flex_inline.oob_copy.tag`

Required:

* `type` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flex_inline--tag"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flex_inline.tag`

Required:

* `type` (String)
Optional:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `vlan_id` (Number) - only applicable when type is 'vlan'
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules--drop_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_rules.drop_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules--drop_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_rules.drop_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules--pass_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_rules.pass_rules.gtp`

Optional:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_overlap_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_overlap_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_rules.pass_rules`

Required:

* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
* `priority` (Number)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_diameter_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_diameter_rules.pass_rules`

Required:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_diameter_rules.pass_rules.diameter`

Optional:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_overlap_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_overlap_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_rules.pass_rules`

Required:

* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `rule_id` (Number)
Optional:

* `comment` (String)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules`

Required:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules.sip`

Optional:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--id_range))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Required:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Required:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules.sip.id_range`

Required:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist5_g_overlap_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist5_g_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist5_g_rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules.pass_rules`

Required:

* `rule_id` (Number)
Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--sip))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules`

Optional:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules.pass_rules`

Required:

* `rule_id` (Number)
Optional:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--gtp))
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--sip))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules.pass_rules.flow5_g`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules.pass_rules.gtp`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules.pass_rules.sip`

Optional:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--fstype"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.fstype`

Optional:

* `offset` (Number) - Offset for Mobility Rotational Flowsampling
* `timer` (Number) - Timer for Mobility Rotational Flowsampling in minutes
* `type` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--gs_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.gs_rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--gs_rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--gs_rules--drop_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.gs_rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--gs_rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.gs_rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
Optional:

* `comment` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--health_state_reasons"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.health_state_reasons`

Optional:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--ip_rewrite"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rewrite"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--roles"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.roles`

Optional:

* `editors` (List of String)
* `listeners` (List of String)
* `owners` (List of String)
* `viewers` (List of String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rules`

Optional:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rules.drop_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--vlan_tag))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rules.drop_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--rewrite"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rules.drop_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rules.drop_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rules.pass_rules`

Required:

* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rule_id` (Number)
Optional:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--ip_rewrite))
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--rewrite))
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--vlan_tag))
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rules.pass_rules.ip_rewrite`

Optional:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--rewrite"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rules.pass_rules.rewrite`

Optional:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.rules.pass_rules.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--ingress_traffic_configs--ingress_traffic_map--vlan_tag"></a>
### Nested Schema for `ingress_traffic_configs.ingress_traffic_map.vlan_tag`

Required:

* `vlan_action` (String)
Optional:

* `tag_protocol_id` (String)
* `vlan_id` (Number)
<a id="nestedatt--monitor_solution_config"></a>
### Nested Schema for `monitor_solution_config`

Optional:

* `action` (String)
* `cluster_id` (String) - cluster id
* `config_status` (String)
* `error_message` (String)
* `exporter_alias` (String) - exporter alias associated with monitor
* `exporter_config` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--exporter_config))
* `gs_group` (String) - gsgroup alias associated with monitor
* `monitor_ip_interface` (String) - ip interface alias associated with monitor
* `ref_sols` (Attributes List) (see [below for nested schema](#nestedatt--monitor_solution_config--ref_sols))
* `solution_alias` (String) - solution alias
* `vport_alias` (String) - vport alias associated with monitor
* `vport_config` (Attributes) - GigaSMART vPort (see [below for nested schema](#nestedatt--monitor_solution_config--vport_config))
Read-Only:

* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--monitor_solution_config--health_state_reasons))
* `user_defined_application_profile` (Dynamic)
* `user_defined_applications` (List of List of String)
<a id="nestedatt--monitor_solution_config--exporter_config"></a>
### Nested Schema for `monitor_solution_config.exporter_config`

Required:

* `alias` (String)
Optional:

* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--exporter_config--cef))
* `description` (String)
* `destination` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--exporter_config--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--exporter_config--mobility_sam))
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--exporter_config--monitor))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--exporter_config--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--exporter_config--snmp))
* `source` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--exporter_config--source))
* `type` (String)
<a id="nestedatt--monitor_solution_config--exporter_config--cef"></a>
### Nested Schema for `monitor_solution_config.exporter_config.cef`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
<a id="nestedatt--monitor_solution_config--exporter_config--destination"></a>
### Nested Schema for `monitor_solution_config.exporter_config.destination`

Optional:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)
<a id="nestedatt--monitor_solution_config--exporter_config--mobility_sam"></a>
### Nested Schema for `monitor_solution_config.exporter_config.mobility_sam`

Optional:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--exporter_config--mobility_sam--event_enable))
* `trigger` (String)
<a id="nestedatt--monitor_solution_config--exporter_config--mobility_sam--event_enable"></a>
### Nested Schema for `monitor_solution_config.exporter_config.mobility_sam.event_enable`

Optional:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--monitor_solution_config--exporter_config--monitor"></a>
### Nested Schema for `monitor_solution_config.exporter_config.monitor`

Optional:

* `timeout` (Number) - how often to export in seconds
<a id="nestedatt--monitor_solution_config--exporter_config--netflow"></a>
### Nested Schema for `monitor_solution_config.exporter_config.netflow`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)
<a id="nestedatt--monitor_solution_config--exporter_config--snmp"></a>
### Nested Schema for `monitor_solution_config.exporter_config.snmp`

Optional:

* `enabled` (Boolean) - snmp reverse lookup enable/disable
<a id="nestedatt--monitor_solution_config--exporter_config--source"></a>
### Nested Schema for `monitor_solution_config.exporter_config.source`

Optional:

* `ip_interface` (String)
<a id="nestedatt--monitor_solution_config--ref_sols"></a>
### Nested Schema for `monitor_solution_config.ref_sols`

Optional:

* `alias` (String)
* `associated_map` (String)
* `type` (String)
<a id="nestedatt--monitor_solution_config--vport_config"></a>
### Nested Schema for `monitor_solution_config.vport_config`

Required:

* `alias` (String)
* `gs_group` (String) - Alias of referenced managing GsGroup
Optional:

* `deferred_binding` (Boolean) - enable/disable deferred-binding
* `fail_over_action` (String)
* `inline_status` (String)
* `inner_traffic_path` (String) - Similar to inline-network traffic-path, applicable for inner map
* `metadata_monitoring` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--vport_config--metadata_monitoring))
* `mode` (String)
* `outer_traffic_path` (String) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
* `sa_apf_profile` (String) - ASF session profile
Read-Only:

* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--monitor_solution_config--vport_config--health_state_reasons))
<a id="nestedatt--monitor_solution_config--vport_config--metadata_monitoring"></a>
### Nested Schema for `monitor_solution_config.vport_config.metadata_monitoring`

Optional:

* `action` (String) - metadata monitoring action
* `exporters` (List of String)
<a id="nestedatt--monitor_solution_config--vport_config--health_state_reasons"></a>
### Nested Schema for `monitor_solution_config.vport_config.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--monitor_solution_config--health_state_reasons"></a>
### Nested Schema for `monitor_solution_config.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_solution.example {solution_alias}
```
