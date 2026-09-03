action "gigavuecore_update_map_flow_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 1
    cluster_id   = "example"
    gtp = {
      imei      = "*"
      imsi      = "*"
      interface = "Gn"
      msisdn    = "*"
      version   = "any"
    }
    rule_id   = "example"
    rule_type = "example"
  }
}
