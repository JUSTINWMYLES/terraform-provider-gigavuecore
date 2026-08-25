resource "gigavuecore_tunnel_application" "example" {
  alias                 = "example"
  cluster_id            = "example"
  config_status         = "example"
  config_status_reasons = [ "example" ]
  decap                 = null
  description           = "example"
  dscp                  = 1
  encap = {
    additonalgsops = {
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
          length = 1
          offset = 1
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
          length = 1
          offset = 1
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
        vlan = 1
      }
      header_remove = {
        ah1                = "example"
        ah2                = "example"
        custom_len         = 1
        erspan_flow_id     = 1
        fp_dst_switch_id   = 1
        fp_src_switch_id   = 1
        header_count       = 1
        offset             = "example"
        offset_range_value = 1
        protocol           = "example"
        timestamp_format   = "example"
        vlan_header        = "example"
        vxlan_id           = 1
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
            avp_codevalue = 1
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
        length       = 1
        offset       = 1
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
        offset   = 1
        protocol = "example"
      }
      ssl_decrypt = {
        in_port  = 1
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
        gmip_port      = 1
        l2_gre_key     = 1
        tls_pcapng = {
          decap_key = "example"
          listener  = "example"
        }
        type = "example"
        vxlan = {
          port_dst = 1
          port_src = 1
          vni      = 1
        }
      }
      tunnel_encap = {
        gmip_config = {
          dscp       = 1
          dst_ip     = "example"
          dst_port   = 1
          flow_label = 1
          prec       = 1
          src_port   = 1
          ttl        = 1
        }
        l2_gre_config = {
          dscp          = 1
          dst_ip        = "example"
          flow_label    = 1
          key           = 1
          pg_dst        = "example"
          prec          = 1
          session_field = "example"
          session_pos   = "example"
          ttl           = 1
        }
        tls_pcapng = {
          exporter       = "example"
          exporter_group = "example"
        }
        type = "example"
        vxlan_config = {
          dscp     = 1
          dst_ip   = "example"
          dst_port = 1
          src_port = 1
          ttl      = 1
          vni      = 1
        }
      }
    }
    export_configs = [{
      exporter_alias          = "example"
      remote_application_port = 1
      remote_ip               = "example"
      source_application_port = 1
    }]
    exporter_group_alias = "example"
    gsop_alias           = "example"
    map_alias            = "example"
    rules = {
      drop_rules = null
      pass_rules = null
    }
    source_port = [ "example" ]
  }
  gsgroup          = "example"
  ip_interface_in  = "example"
  ip_interface_out = "example"
  local_key_alias  = "example"
  remote_key_alias = "example"
  ssl_profile = {
    cipher  = "example"
    mtls    = "example"
    version = "example"
  }
  ssl_profile_alias = "example"
  tcp_profile = {
    keep_alive_timer = 1
    selective_ack    = "example"
    syn_retries      = 1
  }
  tcp_profile_alias = "example"
  traffic_dir       = "example"
  ttl               = 1
  tunnel_type       = "example"
}
