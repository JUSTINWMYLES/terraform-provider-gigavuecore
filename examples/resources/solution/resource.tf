resource "gigavuecore_solution" "example" {
  app_export_config = {
    cache_config = {
      advance_hash = true
      alias = "example"
      description = "example"
      dpi_inject_limit = 1
      event = "example"
      exporters = [ "example" ]
      flow_behavior = "example"
      match = {
        datalink = {
          mac_dst = true
          mac_src = true
          vlan = true
        }
        interface = {
          in_name_width = 1
          in_physical_width = 1
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
            flags = true
            offset = true
          }
          header_len = true
          option_map = true
          precedence = true
          protocol = true
          section = {
            header_size = 1
            payload_size = 1
          }
          source = {
            prefix_min_mask = "example"
          }
          tos = true
          total_length = true
          ttl = true
        }
        ipv6 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp = true
          extension_map = true
          flow_label = true
          fragmentation = {
            flags = true
            offset = true
          }
          hop_limit = true
          length = {
            header = true
            payload = true
            total = true
          }
          next_header = true
          precedence = true
          section = {
            header_size = 1
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
            ack_number = true
            dst_port = true
            flags = true
            header_len = true
            seq_number = true
            src_port = true
            urgent_ptr = true
            window_size = true
          }
          udp = {
            dst_port = true
            msg_len = true
            src_port = true
          }
        }
      }
      multi_collect = true
      network_profiles = [ "example" ]
      observation_domain_id = 1
      sampling = {
        mode = "example"
        single_sampling_rate = 1
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
          name = "example"
          value = "example"
        }]
        is_user_defined = true
        name = "example"
      }]
      destination_name = "example"
      export_ip_interface = "example"
      export_meta_app_profile = {
        alias = "example"
        application_id = true
        applications = [{
          attributes = [{
            name = "example"
            value = "example"
          }]
          is_user_defined = true
          name = "example"
        }]
        counter = {
          bytes = true
          bytes_long = true
          inner_byte = true
          inner_byte_long = true
          packets = true
          packets_long = true
        }
        datalink = {
          mac_dst = true
          mac_src = true
          vlan = true
        }
        description = "example"
        flow = {
          end_reason = true
        }
        gtpu = {
          qfi = true
          teid = true
        }
        interface = {
          in_name_width = 1
          in_physical_width = 1
          out_physical_width = 1
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
            flags = true
            offset = true
          }
          header_len = true
          option_map = true
          precedence = true
          protocol = true
          section = {
            header_size = 1
            payload_size = 1
          }
          source = {
            prefix_min_mask = "example"
          }
          tos = true
          total_length = true
          ttl = true
        }
        ipv6 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp = true
          extension_map = true
          flow_label = true
          fragmentation = {
            flags = true
            offset = true
          }
          hop_limit = true
          length = {
            header = true
            payload = true
            total = true
          }
          next_header = true
          precedence = true
          section = {
            header_size = 1
            payload_size = 1
          }
          source = {
            prefix_min_mask = "example"
          }
          traffic_class = true
        }
        outer_ipv4 = {
          destination = true
          source = true
        }
        outer_ipv6 = {
          destination = true
          source = true
        }
        timestamp = {
          flow_end_msec = true
          flow_endsec = true
          flow_start_msec = true
          flow_startsec = true
          sys_up_time_first = true
          sys_up_time_last = true
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
            ack_number = true
            dst_port = true
            flags = true
            header_len = true
            seq_number = true
            src_port = true
            urgent_ptr = true
            window_size = true
          }
          udp = {
            dst_port = true
            msg_len = true
            src_port = true
          }
        }
        type = "example"
      }
      export_meta_app_profile_alias = "example"
      exporter_alias = "example"
      exporter_config = {
        alias = "example"
        application_profiles = [ "example" ]
        cef = {
          active_timeout = 1
          inactive_timeout = 1
        }
        description = "example"
        destination = {
          dscp = 1
          ipv4_address = "example"
          l4_port_dst = 1
          l4_port_src = 1
          l4_protocol = "example"
          ttl = 1
        }
        max_pkt_size = 1
        mobility_sam = {
          encoding = "example"
          encoding_format = "example"
          event_enable = {
            modify = true
            update = true
          }
          trigger = "example"
        }
        monitor = {
          timeout = 1
        }
        netflow = {
          active_timeout = 1
          inactive_timeout = 1
          template_refresh = 1
          template_type = "example"
          version = "example"
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
      alias = "example"
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
            hash = "example"
            length = 1
            offset = 1
          }
          variable_offset = {
            end_delim = "example"
            hash = "example"
            start_delim = "example"
            start_field = "example"
          }
        }
        gseries_pattern_match = {
          fixed_offset = {
            length = 1
            offset = 1
          }
          variable_offset = {
            end_delim = "example"
            start_delim = "example"
          }
        }
        gtp_whitelist = {
          enabled = "example"
        }
        header_add = {
          vlan = 1
        }
        header_remove = {
          ah1 = "example"
          ah2 = "example"
          custom_len = 1
          erspan_flow_id = 1
          fp_dst_switch_id = 1
          fp_src_switch_id = 1
          header_count = 1
          offset = "example"
          offset_range_value = 1
          protocol = "example"
          timestamp_format = "example"
          vlan_header = "example"
          vxlan_id = 1
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
            app_type = "example"
            diameter_key_hash_type = "example"
            diameter_key_multi_hash_type = [{
              avp_codevalue = 1
              key = "example"
            }]
            gtp_key_hash_type = "example"
            lb_type = "example"
            sip_key_hash_type = "example"
          }
          stateless = {
            field_location = "example"
            hash_fields = "example"
          }
        }
        masking = {
          content_type = "example"
          length = 1
          offset = 1
          pattern = "example"
          protocol = "example"
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
          offset = 1
          protocol = "example"
        }
        ssl_decrypt = {
          in_port = 1
          out_port = 1
        }
        trailer_add = {
          types = [ "example" ]
        }
        trailer_remove = {
          enabled = "example"
        }
        tunnel_decap = {
          custom = {
            port_dst = 1
            port_src = 1
          }
          erspan_flow_id = 1
          gmip_port = 1
          l2_gre_key = 1
          tls_pcapng = {
            decap_key = "example"
            listener = "example"
          }
          type = "example"
          vxlan = {
            port_dst = 1
            port_src = 1
            vni = 1
          }
        }
        tunnel_encap = {
          gmip_config = {
            dscp = 1
            dst_ip = "example"
            dst_port = 1
            flow_label = 1
            prec = 1
            src_port = 1
            ttl = 1
          }
          l2_gre_config = {
            dscp = 1
            dst_ip = "example"
            flow_label = 1
            key = 1
            pg_dst = "example"
            prec = 1
            session_field = "example"
            session_pos = "example"
            ttl = 1
          }
          tls_pcapng = {
            exporter = "example"
            exporter_group = "example"
          }
          type = "example"
          vxlan_config = {
            dscp = 1
            dst_ip = "example"
            dst_port = 1
            src_port = 1
            ttl = 1
            vni = 1
          }
        }
      }
      gs_group = "example"
      health_state = "example"
      health_state_reasons = [{
        message = "example"
        severity = "example"
        traffic_health_state_computation_type = "example"
      }]
    }
  }
  app_filter_config = {
    egress_traffic_config = null
    sapf_profile = {
      alias = "example"
      bidi = true
      buffering = {
        buffer_count_before_match = 1
        enabled = true
        protocol = "example"
      }
      cluster_id = "example"
      packet_count = 1
      session_fields = [{
        pos = 1
        type = "example"
      }]
      timeout = 1
    }
    sapf_profile_alias = "example"
  }
  associated_monitor_solution_alias = "example"
  cluster_id = "example"
  config_status = "example"
  delete_monitor_sol = true
  egress_map_aliases_to_delete = [ "example" ]
  exporter_aliases_to_delete = [ "example" ]
  ingress_map_aliases_to_delete = [ "example" ]
  ingress_traffic_configs = null
  monitor_solution_config = {
    action = "example"
    cluster_id = "example"
    config_status = "example"
    error_message = "example"
    exporter_alias = "example"
    exporter_config = {
      alias = "example"
      application_profiles = [ "example" ]
      cef = {
        active_timeout = 1
        inactive_timeout = 1
      }
      description = "example"
      destination = {
        dscp = 1
        ipv4_address = "example"
        l4_port_dst = 1
        l4_port_src = 1
        l4_protocol = "example"
        ttl = 1
      }
      max_pkt_size = 1
      mobility_sam = {
        encoding = "example"
        encoding_format = "example"
        event_enable = {
          modify = true
          update = true
        }
        trigger = "example"
      }
      monitor = {
        timeout = 1
      }
      netflow = {
        active_timeout = 1
        inactive_timeout = 1
        template_refresh = 1
        template_type = "example"
        version = "example"
      }
      snmp = {
        enabled = true
      }
      source = {
        ip_interface = "example"
      }
      type = "example"
    }
    gs_group = "example"
    monitor_ip_interface = "example"
    ref_sols = [{
      alias = "example"
      associated_map = "example"
      type = "example"
    }]
    solution_alias = "example"
    vport_alias = "example"
    vport_config = {
      alias = "example"
      deferred_binding = true
      fail_over_action = "example"
      gs_group = "example"
      inline_status = "example"
      inner_traffic_path = "example"
      metadata_monitoring = {
        action = "example"
        exporters = [ "example" ]
      }
      mode = "example"
      outer_traffic_path = "example"
      sa_apf_profile = "example"
    }
  }
  solution_alias = "example"
  solution_desc = "example"
  solution_status = "example"
  solution_type = "example"
}
