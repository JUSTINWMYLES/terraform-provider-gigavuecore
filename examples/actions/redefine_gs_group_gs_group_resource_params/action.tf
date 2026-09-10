action "gigavuecore_redefine_gs_group_gs_group_resource_params" "example" {
  config {
    alias           = "example"
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
}
