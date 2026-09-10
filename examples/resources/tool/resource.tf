resource "gigavuecore_tool" "example" {
  alias                      = "example"
  cluster_id                 = "example"
  combined_heart_beat_status = "up"
  comment                    = "example"
  enabled                    = true
  failover_action            = "toolBypass"
  flex_status                = "forwarding"
  flex_traffic_path          = "drop"
  health_state               = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  heartbeat = {
    enabled      = true
    ip_address_a = "example"
    ip_address_b = "example"
    profile      = "example"
    status = {
      heartbeat_passing = "up"
      stats_ato_b = {
        rx_count = 0
        tx_count = 0
      }
      stats_bto_a = {
        rx_count = 0
        tx_count = 0
      }
    }
  }
  inline_tool_type = "external"
  negative_heartbeat = {
    enabled = true
    profile = "example"
    status = {
      heartbeat_passing = "up"
      stats_ato_b = {
        rx_count = 0
        tx_count = 0
      }
      stats_bto_a = {
        rx_count = 0
        tx_count = 0
      }
    }
  }
  operational_state = "disable"
  port_a            = "example"
  port_a_status     = "up"
  port_b            = "example"
  port_b_status     = "up"
  recovery_mode     = "manual"
  shared            = true
  timestamp         = "example"
}
