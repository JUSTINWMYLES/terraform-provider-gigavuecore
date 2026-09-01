action "gigavuecore_add_map_flow_sample_diameter_rule" "example" {
  config {
    alias = "example"
    diameter = {
      user_name = "*"
    }
    interface  = "s6a"
    percentage = 0
    rule_id    = 1
  }
}
