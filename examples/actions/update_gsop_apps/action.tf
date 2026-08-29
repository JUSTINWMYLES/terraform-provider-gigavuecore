action "gigavuecore_update_gsop_apps" "example" {
  config {
    alias = "example"
    apf = {
      enabled = "example"
    }
    cluster_id = "example"
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
}
