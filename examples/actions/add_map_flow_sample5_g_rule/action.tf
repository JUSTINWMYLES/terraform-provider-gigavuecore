action "gigavuecore_add_map_flow_sample5_g_rule" "example" {
  config {
    alias   = "example"
    comment = "example"
    flow5_g = {
      dnn     = "example"
      gpsi    = "*"
      nci     = "a1b2c3d4e"
      nsiid   = "0"
      pei     = "*"
      plmn_id = "123.45"
      supi    = "*"
      tac     = "*"
    }
    percentage = 0
    priority   = 1
    rule_id    = 1
  }
}
