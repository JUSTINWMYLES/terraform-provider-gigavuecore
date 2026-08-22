resource "gigavuecore_gigastream" "example" {
  alias = "example"
  cluster_id = "example"
  comment = "example"
  drop_weight = 1
  failover_status = "example"
  hash_size = 1
  hash_tool_port = [{
    hash_bucket_ids = [ 1 ]
    tool_ports = [ "example" ]
  }]
  hash_type = "example"
  hash_weights = [ 1 ]
  health_state = "example"
  health_state_reasons = [{
    message = "example"
    severity = "example"
    traffic_health_state_computation_type = "example"
  }]
  ports = [ "example" ]
  threshold_level = "example"
  variance_threshold = "example"
}
