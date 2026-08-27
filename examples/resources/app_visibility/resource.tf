resource "gigavuecore_app_visibility" "example" {
  app_export_config = {
    cache_config = {
      advance_hash     = true
      alias            = "example"
      description      = "example"
      dpi_inject_limit = 1
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
          in_name_width     = 1
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
      observation_domain_id = 1
      sampling = {
        mode                 = "example"
        single_sampling_rate = 1
      }
      size = {
        flows = 1
      }
      timeout = {
        idle = 1
      }
    }
    destination_configs = [{
      application_names = [{
        attributes = [{
          name  = "example"
          value = "example"
        }]
        is_user_defined = true
        name            = "example"
      }]
      destination_name = "example"
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
          in_physical_width  = 1
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
        type = "example"
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
          dscp         = 1
          ipv4_address = "example"
          l4_port_dst  = 1
          l4_port_src  = 1
          l4_protocol  = "example"
          ttl          = 1
        }
        max_pkt_size = 1
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
          timeout = 1
        }
        netflow = {
          active_timeout   = 1
          inactive_timeout = 1
          template_refresh = 1
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
      iface = "example"
    }]
  }
  app_filter_config = {
    egress_traffic_configs = [{
      drop_application_names = [{
        attributes = [{
          name  = "example"
          value = "example"
        }]
        is_user_defined = true
        name            = "example"
      }]
      pass_application_names = [{
        attributes = [{
          name  = "example"
          value = "example"
        }]
        is_user_defined = true
        name            = "example"
      }]
      tunnel_aliases = [ "example" ]
    }]
  }
  conn_id            = "example"
  distribute_traffic = true
  dynamic_scale_unit = true
  env_id             = "example"
  ingress_traffic_configs = [{
    source_selector_aliases      = [ "example" ]
    src_raw_end_point_interfaces = [ "example" ]
    tunnel_aliases               = [ "example" ]
    tunnel_interface_mappings = [{
      iface        = "example"
      tunnel_alias = "example"
    }]
  }]
  monitor_solution_config = {
    destination_config = {
      iface = "example"
    }
    mgmt_interface = "example"
  }
  scale_unit     = 1
  solution_alias = "example"
  solution_desc  = "example"
}
