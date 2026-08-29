action "gigavuecore_redefine_gs_group_params" "example" {
  config {
    alias = "example"
    app_tcp = {
      application  = "example"
      load_balance = true
      tcp_control  = "example"
    }
    cluster_id = "example"
    dedup = {
      action    = "example"
      ip_tclass = "example"
      ip_tos    = "example"
      tcp_seq   = "example"
      timer     = 0
      vlan      = "example"
    }
    diameter_packet = {
      timeout = 0
    }
    diameter_s6_a_session = {
      limit   = 0
      timeout = 0
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
      timestamp_format = "example"
    }
    flow_mask = {
      enabled = true
      length  = 0
      offset  = 0
    }
    flow_sampling = {
      ip_ranges = [ "example" ]
      rate      = 0
      timeout   = 0
      type      = "example"
    }
    generic_session_timeout = {
      time = 0
    }
    gpfcp_profiles = {
      g_pfcp_profiles = [ "example" ]
    }
    gs_group_system = {
      cpu_load_alarm_threshold = 0
    }
    gta_profiles = {
      gta_profiles = [ "example" ]
    }
    gtp_control_sampling = {
      enabled = true
    }
    gtp_flow = {
      timeout = 0
    }
    gtp_gpfcp_delay = {
      timeout = 0
    }
    gtp_persistence = {
      enabled          = true
      file_age_timeout = 0
      interval         = 0
      restart_age_time = 0
    }
    gtp_random_sampling = {
      enabled  = true
      interval = 0
    }
    gtp_whitelist = {
      multi_whitelists = [ "example" ]
      whitelist        = "example"
    }
    health_check = {
      action          = "example"
      dst_port        = 0
      enabled         = true
      interval        = 0
      protocol        = "example"
      rcv_port        = 0
      retries         = 0
      round_trip_time = 0
      src_port        = 0
    }
    hsm_group = {
      hsm_group = "example"
    }
    ip_frag = {
      forward              = true
      head_session_timeout = 0
      timeout              = 0
    }
    load_balance = {
      failover = {
        enabled               = true
        threshold_lt_bw       = 0
        threshold_lt_pkt_rate = 0
      }
      link_weight_type = "example"
      replicate_gtpc   = true
    }
    netflow = {
      monitor = "example"
    }
    node_role = {
      mob5_g_limit     = 0
      mob_lte_limit    = 0
      stand_alone_mode = true
      type             = "example"
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
        packet_buffer = 0
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
      log_level          = "example"
      remote_syslog_ip   = "example"
      remote_syslog_port = 0
    }
    sffp_profiles = {
      sffp_profiles = [ "example" ]
    }
    sip_media = {
      timeout = 0
    }
    sip_ports = {
      ports = [ 0 ]
    }
    sip_session = {
      timeout = 0
    }
    sip_tcp_idle_timeout = {
      time = 0
    }
    sip_whitelist = {
      whitelist = "example"
    }
    ssl_decrypt = {
      decrypt_fail_action = "example"
      enabled             = true
      hsm_pkcs11 = {
        debug_level    = 0
        dynamic_object = true
        load_sharing   = true
      }
      hsm_timeout             = 0
      key_cache_timeout       = 0
      key_map                 = "example"
      non_ssl_traffic         = "example"
      pending_session_timeout = 0
      session_timeout         = 0
      tcp_syn_timeout         = 0
      ticket_cache_timeout    = 0
    }
    xpkt_match = {
      enabled = true
    }
  }
}
