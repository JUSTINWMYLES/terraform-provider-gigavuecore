resource "gigavuecore_intent_policy" "example" {
  comment = "example"
  deployed = true
  deployment_error = "example"
  deployment_percent = "example"
  dest_port_timestamp = "example"
  health_state = "example"
  health_state_reasons = [{
    message = "example"
    severity = "example"
    traffic_health_state_computation_type = "example"
  }]
  name = "example"
  policy_id = "example"
  policy_timestamp = "example"
  priority = true
  rules = null
  src_port_timestamp = "example"
  src_ports_info = {
    comment = "example"
    ports = [ "example" ]
    template_ids = [ "example" ]
  }
  tags = [{
    tag_key = "example"
    tag_values = [ "example" ]
  }]
}
