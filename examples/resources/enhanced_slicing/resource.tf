resource "gigavuecore_enhanced_slicing" "example" {
  alias        = "example"
  hash_field   = "example"
  max_sessions = 0
  protocol_fields = [{
    gtp = {
      flow_session = {
        action         = "example"
        skip_pkt_count = 0
        timeout        = 0
        value          = "example"
      }
      offset   = 0
      protocol = "example"
    }
    gtp_u = {
      flow_session = {
        action         = "example"
        skip_pkt_count = 0
        timeout        = 0
        value          = "example"
      }
      l4_port  = 0
      offset   = 0
      protocol = "example"
    }
    ip = {
      flow_session = {
        action         = "example"
        skip_pkt_count = 0
        timeout        = 0
        value          = "example"
      }
      offset   = 0
      protocol = "example"
      value    = "example"
    }
    none = {
      flow_session = {
        action         = "example"
        skip_pkt_count = 0
        timeout        = 0
        value          = "example"
      }
      offset = 0
    }
    transport = {
      flow_session = {
        action         = "example"
        skip_pkt_count = 0
        timeout        = 0
        value          = "example"
      }
      l4_port  = 0
      offset   = 0
      protocol = "example"
      value    = "example"
    }
  }]
}
