resource "gigavuecore_network_profile" "example" {
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
  ipv4_subnets = ["example"]
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
  ipv6_masks = ["example"]
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
