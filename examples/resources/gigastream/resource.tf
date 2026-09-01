resource "gigavuecore_gigastream" "example" {
  alias           = "example"
  cluster_id      = "example"
  comment         = "example"
  drop_weight     = 0
  failover_status = "enable"
  hash_size       = 1
  hash_tool_port = [{
    hash_bucket_ids = [0]
    tool_ports      = ["example"]
  }]
  hash_type    = "advanced"
  hash_weights = [0]
  health_state = "green"
  health_state_reasons = [{
    message                               = "example"
    severity                              = "green"
    traffic_health_state_computation_type = "PORT_LOW_UTIL"
  }]
  ports              = ["example"]
  threshold_level    = "Global-Level"
  variance_threshold = "example"
}
