resource "gigavuecore_inline_ssl_app" "example" {
  alias = "example"
  app_intent_configs = {
    black_list_config = {
      operation = "example"
      profile_list = {
        file_source = {
          hostname = "example"
          password = "example"
          path = "example"
          protocol = "example"
          username = "example"
        }
        list = "example"
      }
    }
    global_default_configs = {
      caching = {
        persistence = {
          enable = true
        }
      }
      cluster_id = "example"
      dhe_ciphersuit = "example"
      monitor = {
        enable = true
      }
      resumption = {
        client = {
          enable = true
        }
      }
      ssl_versions = {
        connection_reset_action_for_max_version = "example"
        connection_reset_action_for_min_version = "example"
        max_version = "example"
        min_version = "example"
      }
      start_tls = {
        enable = true
      }
    }
    gs_engines = [ "example" ]
    gs_group_alias = "example"
    gs_group_param_configs = {
      hsm_group = "example"
      session_logging = {
        interface = "example"
        log_level = "example"
        remote_syslog_ip = "example"
        remote_syslog_port = 1
      }
    }
    gsop_alias = "example"
    inline_ssl = {
      standalone = true
    }
    key_store_configs = {
      deployment_type = "example"
      inbound_keys = [{
        key_alias = "example"
        server_domain_alias = "example"
      }]
      outboundkeys = [{
        key_alias = "example"
        signing_for = "example"
      }]
    }
    m_tls_configs = {
      primary_signing = "example"
      secondary_signing = "example"
      trust_store = "example"
    }
    network_access_configs = [{
      network_access = {
        cluster_id = "example"
        dhcp = true
        dns = "example"
        eport = "example"
        gateway = "example"
        hw_address = "example"
        interface = "example"
        ip_address = "example"
        ip_mask = "example"
        mtu = 1
        proxy_server_profile = "example"
        status = "example"
        vlan = 1
      }
      operation = "example"
    }]
    ssl_path = [{
      alias = "example"
      flex_inline_map = {
        a_to_b = {
          ib_pathway = "example"
          tools = [ "example" ]
          type = "example"
        }
        b_to_a = {
          ib_pathway = "example"
          tools = [ "example" ]
          type = "example"
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
        svt_tag = 1
        tag = {
          tag_protocol_id = "example"
          type = "example"
          vlan_id = 1
        }
      }
    }]
    ssl_profile_alias = "example"
    ssl_profile_config = {
      alias = "example"
      certificate = {
        expired = "example"
        invalid = "example"
        revocation = {
          crl = {
            defer = 1
            enabled = true
            fail = "example"
          }
          ocsp = {
            defer = 1
            enabled = true
            fail = "example"
          }
        }
        self_signed = "example"
        unknown_ca = "example"
      }
      client_auth = {
        expired = "example"
        invalid = "example"
        revocation = {
          crl = {
            defer = 1
            enabled = true
            fail = "example"
          }
          ocsp = {
            defer = 1
            enabled = true
            fail = "example"
          }
        }
        self_signed = "example"
        unknown_ca = "example"
      }
      cluster_id = "example"
      decrypt = {
        tcp = {
          inactive_timeout = 1
          port_map = {
            default_out_port = 1
            ports = [{
              in_port = 1
              out_port = 1
              rule_id = 1
            }]
          }
        }
        tool_bypass = {
          enable = true
        }
      }
      default_action = "example"
      high_avail = {
        active_standby = {
          enable = true
        }
      }
      key_map = [{
        hostname = "example"
        key = "example"
        rule_id = 1
      }]
      monitor = "example"
      network_group = {
        multiple_entry = {
          enable = true
        }
      }
      no_decrypt = {
        tool_bypass = {
          enable = true
        }
      }
      non_ssl_tcp = {
        tool_bypass = {
          enable = true
        }
      }
      rules = null
      split_proxy = {
        mode = {
          enable = true
        }
        server_non_pfs_ciphers = {
          enable = true
        }
      }
      start_tls = {
        l4_port = [ 1 ]
      }
      tcp = {
        delayed_ack = true
        syn_retries = 1
        timewait_timeout = 1
      }
      tool = {
        early_engage = true
        fail_action = "example"
      }
      url_cache = {
        miss_action = "example"
        timeout = 1
      }
    }
    tag_protocol_id = "example"
    trust_store_configs = {
      trust_store_append_configs = {
        file = "example"
        file_source = {
          hostname = "example"
          password = "example"
          path = "example"
          protocol = "example"
          username = "example"
        }
      }
      trust_store_replace_configs = {
        file = "example"
        file_source = {
          hostname = "example"
          password = "example"
          path = "example"
          protocol = "example"
          username = "example"
        }
      }
    }
    vlan_id = 1
    vport_alias = "example"
    vport_config = {
      alias = "example"
      deferred_binding = true
      fail_over_action = "example"
      gs_group = "example"
      health_state = "example"
      health_state_reasons = [{
        message = "example"
        severity = "example"
        traffic_health_state_computation_type = "example"
      }]
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
    white_list_config = {
      operation = "example"
      profile_list = {
        file_source = {
          hostname = "example"
          password = "example"
          path = "example"
          protocol = "example"
          username = "example"
        }
        list = "example"
      }
    }
  }
  cluster_name = "example"
  config_status = "example"
  config_status_reasons = [ "example" ]
  health_state = "example"
  health_state_reasons = [{
    message = "example"
    severity = "example"
    traffic_health_state_computation_type = "example"
  }]
  m_tls = "example"
  ria_configs = [{
    cluster_name = "example"
    gs_engines = [ "example" ]
    gs_group_alias = "example"
    gsop_alias = "example"
    network_access_configs = [{
      network_access = {
        cluster_id = "example"
        dhcp = true
        dns = "example"
        eport = "example"
        gateway = "example"
        hw_address = "example"
        interface = "example"
        ip_address = "example"
        ip_mask = "example"
        mtu = 1
        proxy_server_profile = "example"
        status = "example"
        vlan = 1
      }
      operation = "example"
    }]
    ssl_profile_alias = "example"
    vport_alias = "example"
  }]
  ria_enabled = "example"
}
