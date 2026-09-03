action "gigavuecore_update_gs_group_params" "example" {
  config {
    alias = "example"
    app_tcp = {
      application  = "broadcast"
      load_balance = true
      tcp_control  = "broadcast"
    }
    cluster_id = "example"
    dedup = {
      action    = "count"
      ip_tclass = "include"
      ip_tos    = "include"
      tcp_seq   = "include"
      timer     = 10
      vlan      = "include"
    }
    diameter_packet = {
      timeout = 1
    }
    diameter_s6_a_session = {
      limit   = 1
      timeout = 30
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
      timestamp_format = "gigasmart"
    }
    flow_mask = {
      enabled = true
      length  = 1
      offset  = 0
    }
    flow_sampling = {
      ip_ranges = ["example"]
      rate      = 5
      timeout   = 1
      type      = "deviceIp"
    }
    generic_session_timeout = {
      time = 5
    }
    gpfcp_profiles = {
      g_pfcp_profiles = ["example"]
    }
    gs_group_system = {
      cpu_load_alarm_threshold = 20
    }
    gta_profiles = {
      gta_profiles = ["example"]
    }
    gtp_control_sampling = {
      enabled = true
    }
    gtp_flow = {
      timeout = 1
    }
    gtp_gpfcp_delay = {
      timeout = 0
    }
    gtp_persistence = {
      enabled          = true
      file_age_timeout = 10
      interval         = 10
      restart_age_time = 10
    }
    gtp_random_sampling = {
      enabled  = true
      interval = 12
    }
    gtp_whitelist = {
      multi_whitelists = ["example"]
      whitelist        = "example"
    }
    health_check = {
      action          = "pass"
      dst_port        = 1
      enabled         = true
      interval        = 5
      protocol        = "icmp"
      rcv_port        = 1
      retries         = 1
      round_trip_time = 1
      src_port        = 1
    }
    hsm_group = {
      hsm_group = "example"
    }
    ip_frag = {
      forward              = true
      head_session_timeout = 15
      timeout              = 5
    }
    load_balance = {
      failover = {
        enabled               = true
        threshold_lt_bw       = 50
        threshold_lt_pkt_rate = 500
      }
      link_weight_type = "speed"
      replicate_gtpc   = true
    }
    netflow = {
      monitor = "example"
    }
    node_role = {
      mob5_g_limit     = 1
      mob_lte_limit    = 1
      stand_alone_mode = true
      type             = "control"
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
        packet_buffer = 20
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
      log_level          = "err"
      remote_syslog_ip   = "example"
      remote_syslog_port = 0
    }
    sffp_profiles = {
      sffp_profiles = ["example"]
    }
    sip_media = {
      timeout = 30
    }
    sip_ports = {
      ports = [1]
    }
    sip_session = {
      timeout = 30
    }
    sip_tcp_idle_timeout = {
      time = 20
    }
    sip_whitelist = {
      whitelist = "example"
    }
    ssl_decrypt = {
      decrypt_fail_action = "drop"
      enabled             = true
      hsm_pkcs11 = {
        debug_level    = 0
        dynamic_object = true
        load_sharing   = true
      }
      hsm_timeout             = 2
      key_cache_timeout       = 1
      key_map                 = "example"
      non_ssl_traffic         = "drop"
      pending_session_timeout = 30
      session_timeout         = 30
      tcp_syn_timeout         = 20
      ticket_cache_timeout    = 1
    }
    xpkt_match = {
      enabled = true
    }
  }
}
