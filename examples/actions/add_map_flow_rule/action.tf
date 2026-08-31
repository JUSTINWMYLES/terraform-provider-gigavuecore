action "gigavuecore_add_map_flow_rule" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    gtp = {
      imei      = "example"
      imsi      = "example"
      interface = "Gn"
      msisdn    = "example"
      version   = "any"
    }
    rule_id   = 1
    rule_type = "example"
  }
}
