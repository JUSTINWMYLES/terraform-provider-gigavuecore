resource "gigavuecore_enhanced_slicing" "example" {
  alias        = "example"
  hash_field   = "5tuple"
  max_sessions = 4
  protocol_fields = [{
    gtp = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      offset   = 0
      protocol = "gtp"
    }
    gtp_u = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      l4_port  = 1
      offset   = 0
      protocol = "gtpu-tcp"
    }
    ip = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      offset   = 0
      protocol = "ip"
      value    = "inner"
    }
    none = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      offset = 0
    }
    transport = {
      flow_session = {
        action         = "slice"
        skip_pkt_count = 1
        timeout        = 10
        value          = "inner"
      }
      l4_port  = 1
      offset   = 0
      protocol = "tcp"
      value    = "inner"
    }
  }]
}
