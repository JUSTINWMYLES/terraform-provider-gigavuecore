resource "gigavuecore_fm_template" "example" {
  config = {
    aaa_auth_config = {
      auth_sequence = [ "example" ]
      external_login_mapping = {
        default_local_user = "example"
        user_map_order = "example"
      }
    }
    acme_certificate = [{
      acme_server_url = "example"
      algorithm = "example"
      operation_type = "example"
      renew_days = 1
    }]
    device_ssl_certificate_configs = [{
      issuer = "example"
      not_after = "example"
      not_before = "example"
      operation_type = "example"
      signature_algorithm = "example"
      subject = "example"
      trusted_ca = {
        name = "example"
      }
      upload_spec = {
        info = {
          comment = "example"
          name = "example"
          passphrase = "example"
          type = "example"
        }
        pem = "example"
      }
    }]
    export_metadata_app_profile = {
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
    giga_port_neighbors_discovery_config = {
      resource_configs = null
    }
    giga_stream_threshold_config = {
      giga_stream_type_thresholds = [{
        type = "example"
        variance_threshold = 1.0
      }]
    }
    giga_user_defined_application_config = {
      alias = "example"
      app_id = 1
      priority = 1
      rules = {
        rule = [{
          address = "example"
          code = "example"
          common_name = "example"
          content = "example"
          cts_cookie = "example"
          cts_page_url = "example"
          cts_referer = "example"
          cts_server = "example"
          cts_uri = "example"
          cts_user_agent = "example"
          dscp = "example"
          mime_type = "example"
          mindata = 1
          port = "example"
          resolv_name = "example"
          stc_location = "example"
          stc_server_agent = "example"
          stc_subject_alt_name = "example"
          stream = "example"
          typeval = "example"
          user_agent = "example"
        }]
      }
    }
    ldap_server_system_config = null
    ldap_servers = [{
      order = "example"
      server_address = "example"
    }]
    metadata_exporter = {
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
    port_packet_threshold_config = {
      drop_threshold = {
        rx = [{
          count_ = 1
          percent = 1.0
          port_type = "example"
        }]
        tx = [{
          count_ = 1
          percent = 1.0
          port_type = "example"
        }]
      }
      error_threshold = {
        rx = [{
          count_ = 1
          percent = 1.0
          port_type = "example"
        }]
        tx = [{
          count_ = 1
          percent = 1.0
          port_type = "example"
        }]
      }
    }
    proxy_server_profile = {
      alias = "example"
      auth_type = "example"
      comment = "example"
      password = "example"
      periodic_ping = "example"
      periodic_ping_failure_retry = 1
      periodic_ping_interval = 1
      periodic_ping_type = "example"
      port = 1
      protocol = "example"
      proxy_address = "example"
      ssl_apps = {
        cluster_name = [ "example" ]
      }
      username = "example"
    }
    snmp_trap_event_configs = [{
      enabled = true
      notify_event = "example"
    }]
    snmp_v3_users_config = {
      snmp_v3_user = [{
        auth_key = "example"
        auth_protocol = "example"
        min_sw_version = "example"
        previous_username = "example"
        priv_key = "example"
        priv_protocol = "example"
        username = "example"
      }]
    }
    ssh_ciphers_config = {
      classic = {
        client_ciphers = [ "example" ]
        client_hostkey = [ "example" ]
        client_kex = [ "example" ]
        client_macs = [ "example" ]
        server_ciphers = [ "example" ]
        server_hostkey = [ "example" ]
        server_kex = [ "example" ]
        server_macs = [ "example" ]
      }
      crypto = {
        client_ciphers = [ "example" ]
        client_hostkey = [ "example" ]
        client_kex = [ "example" ]
        client_macs = [ "example" ]
        server_ciphers = [ "example" ]
        server_hostkey = [ "example" ]
        server_kex = [ "example" ]
        server_macs = [ "example" ]
      }
      fips = {
        client_ciphers = [ "example" ]
        client_hostkey = [ "example" ]
        client_kex = [ "example" ]
        client_macs = [ "example" ]
        server_ciphers = [ "example" ]
        server_hostkey = [ "example" ]
        server_kex = [ "example" ]
        server_macs = [ "example" ]
      }
    }
  }
  config_level = "example"
  config_level_value = [ "example" ]
  config_resource = null
  config_type = "example"
  modifiable = true
  ref_count = 1
  ref_object = {
    ref_object_type = "example"
  }
  template_name = "example"
  update_time = "example"
}
