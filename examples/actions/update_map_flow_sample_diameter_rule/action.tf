action "gigavuecore_update_map_flow_sample_diameter_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 1
    diameter = {
      user_name = "example"
    }
    interface  = "s6a"
    percentage = 0
    rule_id    = "example"
  }
}
