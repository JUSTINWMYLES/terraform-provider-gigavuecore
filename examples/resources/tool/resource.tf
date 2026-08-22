resource "gigavuecore_tool" "example" {
  alias = "example"
  combined_heart_beat_status = "example"
  comment = "example"
  enabled = true
  failover_action = "example"
  flex_status = "example"
  flex_traffic_path = "example"
  health_state = "example"
  health_state_reasons = [{
    message = "example"
    severity = "example"
    traffic_health_state_computation_type = "example"
  }]
  heartbeat = {
    enabled = true
    ip_address_a = "example"
    ip_address_b = "example"
    profile = "example"
    status = {
      heartbeat_passing = "example"
      stats_ato_b = {
        rx_count = 1
        tx_count = 1
      }
      stats_bto_a = {
        rx_count = 1
        tx_count = 1
      }
    }
  }
  inline_tool_type = "example"
  negative_heartbeat = {
    enabled = true
    profile = "example"
    status = {
      heartbeat_passing = "example"
      stats_ato_b = {
        rx_count = 1
        tx_count = 1
      }
      stats_bto_a = {
        rx_count = 1
        tx_count = 1
      }
    }
  }
  operational_state = "example"
  port_a = "example"
  port_a_status = "example"
  port_b = "example"
  port_b_status = "example"
  recovery_mode = "example"
  shared = true
  timestamp = "example"
}
