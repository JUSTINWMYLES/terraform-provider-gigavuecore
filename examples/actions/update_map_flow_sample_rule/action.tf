action "gigavuecore_update_map_flow_sample_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 1
    cluster_id   = "example"
    comment      = "example"
    gtp = {
      apn       = "example"
      eci       = "example"
      imei      = "example"
      imsi      = "example"
      interface = "Gn"
      msisdn    = "example"
      nas_5_qi  = "0"
      nci       = "example"
      plmn_id   = "example"
      qci       = 0
      snssai    = "0"
      tac       = "abc1"
      tac_5_g   = "example"
      version   = "any"
    }
    percentage      = 0
    periodic_recalc = true
    priority        = 1
    rule_id         = "example"
  }
}
