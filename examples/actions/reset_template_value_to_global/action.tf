action "gigavuecore_reset_template_value_to_global" "example" {
  config {
    body_config_type = "SNMPTRAPS"
    config = {
      aaa_auth_config = {
        auth_sequence = [ "local" ]
        external_login_mapping = {
          default_local_user = "example"
          user_map_order     = "localOnly"
        }
      }
      acme_certificate = [{
        acme_server_url = "example"
        algorithm       = "rsa-2048"
        operation_type  = "issue"
        renew_days      = 0
      }]
      device_ssl_certificate_configs = [{
        issuer              = "example"
        not_after           = "example"
        not_before          = "example"
        operation_type      = "add"
        signature_algorithm = "example"
        subject             = "example"
        trusted_ca = {
          name = "example"
        }
        upload_spec = {
          info = {
            comment    = "example"
            name       = "example"
            passphrase = "example"
            type       = "privateKey"
          }
          pem = "example"
        }
      }]
      export_metadata_app_profile = {
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
      giga_port_neighbors_discovery_config = {
        resource_configs = [{
          cdp       = true
          gdp       = true
          lldp      = true
          port_type = "example"
        }]
      }
      giga_stream_threshold_config = {
        giga_stream_type_thresholds = [{
          type               = "all"
          variance_threshold = 1.0
        }]
      }
      giga_user_defined_application_config = {
        alias    = "example"
        app_id   = 0
        priority = 0
        rules = {
          rule = [{
            address              = "example"
            code                 = "example"
            common_name          = "example"
            content              = "example"
            cts_cookie           = "example"
            cts_page_url         = "example"
            cts_referer          = "example"
            cts_server           = "example"
            cts_uri              = "example"
            cts_user_agent       = "example"
            dscp                 = "example"
            mime_type            = "example"
            mindata              = 0
            port                 = "example"
            resolv_name          = "example"
            stc_location         = "example"
            stc_server_agent     = "example"
            stc_subject_alt_name = "example"
            stream               = "example"
            typeval              = "example"
            user_agent           = "example"
          }]
        }
      }
      ldap_server_system_config = "example"
      ldap_servers = [{
        order          = "example"
        server_address = "example"
      }]
      metadata_exporter = {
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
      port_packet_threshold_config = {
        drop_threshold = {
          rx = [{
            count_    = 0
            percent   = 1.0
            port_type = "all"
          }]
          tx = [{
            count_    = 0
            percent   = 1.0
            port_type = "all"
          }]
        }
        error_threshold = {
          rx = [{
            count_    = 0
            percent   = 1.0
            port_type = "all"
          }]
          tx = [{
            count_    = 0
            percent   = 1.0
            port_type = "all"
          }]
        }
      }
      proxy_server_profile = {
        alias                       = "example"
        auth_type                   = "none"
        comment                     = "example"
        password                    = "example"
        periodic_ping               = "enable"
        periodic_ping_failure_retry = 1
        periodic_ping_interval      = 1
        periodic_ping_type          = "http-connect"
        port                        = 1
        protocol                    = "http"
        proxy_address               = "example"
        ssl_apps = {
          cluster_name = [ "example" ]
        }
        username = "example"
      }
      snmp_trap_event_configs = [{
        enabled      = true
        notify_event = "example"
      }]
      snmp_v3_users_config = {
        snmp_v3_user = [{
          auth_key          = "example"
          auth_protocol     = "md5"
          min_sw_version    = "example"
          previous_username = "example"
          priv_key          = "example"
          priv_protocol     = "des"
          username          = "example"
        }]
      }
      ssh_ciphers_config = {
        classic = {
          client_ciphers = [ "default" ]
          client_hostkey = [ "default" ]
          client_kex     = [ "default" ]
          client_macs    = [ "default" ]
          server_ciphers = [ "default" ]
          server_hostkey = [ "default" ]
          server_kex     = [ "default" ]
          server_macs    = [ "default" ]
        }
        crypto = {
          client_ciphers = [ "default" ]
          client_hostkey = [ "default" ]
          client_kex     = [ "default" ]
          client_macs    = [ "default" ]
          server_ciphers = [ "default" ]
          server_hostkey = [ "default" ]
          server_kex     = [ "default" ]
          server_macs    = [ "default" ]
        }
        fips = {
          client_ciphers = [ "default" ]
          client_hostkey = [ "default" ]
          client_kex     = [ "default" ]
          client_macs    = [ "default" ]
          server_ciphers = [ "default" ]
          server_hostkey = [ "default" ]
          server_kex     = [ "default" ]
          server_macs    = [ "default" ]
        }
      }
    }
    config_level       = "GLOBAL"
    config_level_value = [ "example" ]
    config_resource    = "example"
    config_type        = "example"
    modifiable         = true
    ref_count          = 0
    ref_object = {
      ref_object_type = "PHYSICAL"
    }
    template_name = "example"
    update_time   = "example"
  }
}
