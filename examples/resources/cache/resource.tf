resource "gigavuecore_cache" "example" {
  advance_hash     = true
  alias            = "example"
  description      = "example"
  dpi_inject_limit = 0
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
      in_name_width     = 0
      in_physical_width = 0
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
        header_size  = 0
        payload_size = 0
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
        header_size  = 0
        payload_size = 0
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
  observation_domain_id = 0
  sampling = {
    mode                 = "example"
    single_sampling_rate = 0
  }
  size = {
    flows = 0
  }
  timeout = {
    idle = 0
  }
}
