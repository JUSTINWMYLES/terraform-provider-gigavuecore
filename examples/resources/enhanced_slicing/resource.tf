resource "gigavuecore_enhanced_slicing" "example" {
  alias = "example"
  hash_field = "example"
  max_sessions = 1
  protocol_fields = [{
    gtp = {
      flow_session = {
        action = "example"
        skip_pkt_count = 1
        timeout = 1
        value = "example"
      }
      offset = 1
      protocol = "example"
    }
    gtp_u = {
      flow_session = {
        action = "example"
        skip_pkt_count = 1
        timeout = 1
        value = "example"
      }
      l4_port = 1
      offset = 1
      protocol = "example"
    }
    ip = {
      flow_session = {
        action = "example"
        skip_pkt_count = 1
        timeout = 1
        value = "example"
      }
      offset = 1
      protocol = "example"
      value = "example"
    }
    none = {
      flow_session = {
        action = "example"
        skip_pkt_count = 1
        timeout = 1
        value = "example"
      }
      offset = 1
    }
    transport = {
      flow_session = {
        action = "example"
        skip_pkt_count = 1
        timeout = 1
        value = "example"
      }
      l4_port = 1
      offset = 1
      protocol = "example"
      value = "example"
    }
  }]
}
