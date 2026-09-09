action "gigavuecore_add_map_flow_rule" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    gtp = {
      imei      = "*"
      imsi      = "*"
      interface = "Gn"
      msisdn    = "*"
      version   = "any"
    }
    rule_id   = 1
    rule_type = "example"
  }
}
