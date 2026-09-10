resource "gigavuecore_inline_ssl_app" "example" {
  alias = "example"
  app_intent_configs = {
    black_list_config = {
      operation = "ADD"
      profile_list = {
        file_source = {
          hostname = "example"
          password = "example"
          path     = "example"
          protocol = "scp"
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
      cluster_id     = "example"
      dhe_ciphersuit = "disable"
      monitor = {
        enable = true
      }
      resumption = {
        client = {
          enable = true
        }
      }
      ssl_versions = {
        connection_reset_action_for_max_version = "no-decrypt"
        connection_reset_action_for_min_version = "no-decrypt"
        max_version                             = "sslv3"
        min_version                             = "sslv3"
      }
      start_tls = {
        enable = true
      }
    }
    gs_engines     = ["example"]
    gs_group_alias = "example"
    gs_group_param_configs = {
      hsm_group = "example"
      session_logging = {
        interface          = "example"
        log_level          = "err"
        remote_syslog_ip   = "example"
        remote_syslog_port = 0
      }
    }
    gsop_alias = "example"
    inline_ssl = {
      standalone = true
    }
    key_store_configs = {
      deployment_type = "Inbound"
      inbound_keys = [{
        key_alias           = "example"
        server_domain_alias = "example"
      }]
      outboundkeys = [{
        key_alias   = "example"
        signing_for = "Primary"
      }]
    }
    m_tls_configs = {
      primary_signing   = "example"
      secondary_signing = "example"
      trust_store       = "example"
    }
    network_access_configs = [{
      network_access = {
        cluster_id           = "example"
        dhcp                 = true
        dns                  = "example"
        eport                = "example"
        gateway              = "example"
        hw_address           = "example"
        interface            = "eth2"
        ip_address           = "example"
        ip_mask              = "example"
        mtu                  = 68
        proxy_server_profile = "example"
        status               = "up"
        vlan                 = 20
      }
      operation = "Add"
    }]
    ssl_path = [{
      alias = "example"
      flex_inline_map = {
        a_to_b = {
          ib_pathway = "example"
          tools      = ["example"]
          type       = "bypass"
        }
        b_to_a = {
          ib_pathway = "example"
          tools      = ["example"]
          type       = "bypass"
        }
        oob_copy = [{
          direction = "aToB"
          dst_ports = ["example"]
          src_ports = ["example"]
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
    }]
    ssl_profile_alias = "example"
    ssl_profile_config = {
      alias = "example"
      certificate = {
        expired = "decrypt"
        invalid = "decrypt"
        revocation = {
          crl = {
            defer   = 20
            enabled = true
            fail    = "soft"
          }
          ocsp = {
            defer   = 20
            enabled = true
            fail    = "soft"
          }
        }
        self_signed = "decrypt"
        unknown_ca  = "decrypt"
      }
      client_auth = {
        expired = "decrypt"
        invalid = "decrypt"
        revocation = {
          crl = {
            defer   = 20
            enabled = true
            fail    = "soft"
          }
          ocsp = {
            defer   = 20
            enabled = true
            fail    = "soft"
          }
        }
        self_signed = "decrypt"
        unknown_ca  = "decrypt"
      }
      cluster_id = "example"
      decrypt = {
        tcp = {
          inactive_timeout = 2
          port_map = {
            default_out_port = 0
            ports = [{
              in_port  = 1
              out_port = 1
              rule_id  = 0
            }]
          }
        }
        tool_bypass = {
          enable = true
        }
      }
      default_action = "decrypt"
      high_avail = {
        active_standby = {
          enable = true
        }
      }
      key_map = [{
        hostname = "example"
        key      = "example"
        rule_id  = 0
      }]
      monitor = "enable"
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
      rules = ["example"]
      split_proxy = {
        mode = {
          enable = true
        }
        server_non_pfs_ciphers = {
          enable = true
        }
      }
      start_tls = {
        l4_port = [0]
      }
      tcp = {
        delayed_ack      = true
        syn_retries      = 0
        timewait_timeout = 0
      }
      tool = {
        early_engage = true
        fail_action  = "fail-open"
      }
      url_cache = {
        miss_action = "decrypt"
        timeout     = 1
      }
    }
    tag_protocol_id = "0x8100"
    trust_store_configs = {
      trust_store_append_configs = {
        file = "example"
        file_source = {
          hostname = "example"
          password = "example"
          path     = "example"
          protocol = "scp"
          username = "example"
        }
      }
      trust_store_replace_configs = {
        file = "example"
        file_source = {
          hostname = "example"
          password = "example"
          path     = "example"
          protocol = "scp"
          username = "example"
        }
      }
    }
    vlan_id     = 0
    vport_alias = "example"
    vport_config = {
      alias            = "example"
      deferred_binding = true
      fail_over_action = "vport-bypass"
      gs_group         = "example"
      health_state     = "green"
      health_state_reasons = [{
        message                               = "example"
        severity                              = "green"
        traffic_health_state_computation_type = "PORT_LOW_UTIL"
      }]
      inline_status      = "up"
      inner_traffic_path = "to-inline-tool"
      metadata_monitoring = {
        action    = "enable"
        exporters = ["example"]
      }
      mode               = "none"
      outer_traffic_path = "to-inline-tool"
      sa_apf_profile     = "example"
    }
    white_list_config = {
      operation = "ADD"
      profile_list = {
        file_source = {
          hostname = "example"
          password = "example"
          path     = "example"
          protocol = "scp"
          username = "example"
        }
        list = "example"
      }
    }
  }
  cluster_id            = "example"
  cluster_name          = "example"
  config_status         = "SUCCESS"
  config_status_reasons = ["example"]
  health_state          = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  m_tls = "false"
  ria_configs = [{
    cluster_name   = "example"
    gs_engines     = ["example"]
    gs_group_alias = "example"
    gsop_alias     = "example"
    network_access_configs = [{
      network_access = {
        cluster_id           = "example"
        dhcp                 = true
        dns                  = "example"
        eport                = "example"
        gateway              = "example"
        hw_address           = "example"
        interface            = "eth2"
        ip_address           = "example"
        ip_mask              = "example"
        mtu                  = 68
        proxy_server_profile = "example"
        status               = "up"
        vlan                 = 20
      }
      operation = "Add"
    }]
    ssl_profile_alias = "example"
    vport_alias       = "example"
  }]
  ria_enabled = "false"
}
