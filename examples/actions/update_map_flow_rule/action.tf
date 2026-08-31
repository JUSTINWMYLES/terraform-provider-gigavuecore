action "gigavuecore_update_map_flow_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 1
    cluster_id   = "example"
    gtp = {
      imei      = "example"
      imsi      = "example"
      interface = "Gn"
      msisdn    = "example"
      version   = "any"
    }
    rule_id   = "example"
    rule_type = "example"
  }
}
