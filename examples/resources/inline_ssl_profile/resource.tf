resource "gigavuecore_inline_ssl_profile" "example" {
  alias = "example"
  certificate = {
    expired = "example"
    invalid = "example"
    revocation = {
      crl = {
        defer   = 0
        enabled = true
        fail    = "example"
      }
      ocsp = {
        defer   = 0
        enabled = true
        fail    = "example"
      }
    }
    self_signed = "example"
    unknown_ca  = "example"
  }
  cluster_id = "example"
  decrypt = {
    tcp = {
      inactive_timeout = 0
      port_map = {
        default_out_port = 0
        ports = [{
          in_port  = 0
          out_port = 0
          rule_id  = 0
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
  inbound_tool_early_inspect = {
    connection_timeout = 0
    mode = {
      enable = true
    }
  }
  key_map = [{
    hostname = "example"
    key      = "example"
    rule_id  = 0
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
  one_arm = "example"
  resilient_inline = {
    mode = {
      enable = true
    }
  }
  rules = [ "example" ]
  split_proxy = {
    mode = {
      enable = true
    }
    server_non_pfs_ciphers = {
      enable = true
    }
  }
  start_tls = {
    l4_port = [ 0 ]
  }
  tcp = {
    delayed_ack      = true
    syn_retries      = 0
    timewait_timeout = 0
  }
  tool = {
    early_engage = true
    fail_action  = "example"
  }
  tool_l3 = {
    cache_server_cert_timeout = 0
    http2_downgrade = {
      enable = true
    }
    nat_pat = {
      enable = true
    }
  }
  url_cache = {
    miss_action = "example"
    timeout     = 0
  }
}
