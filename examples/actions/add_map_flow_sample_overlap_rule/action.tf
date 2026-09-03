action "gigavuecore_add_map_flow_sample_overlap_rule" "example" {
  config {
    alias   = "example"
    comment = "example"
    gtp = {
      apn       = "example"
      eci       = "a1b2c3d4"
      imei      = "*"
      imsi      = "*"
      interface = "Gn"
      msisdn    = "*"
      nas_5_qi  = "0"
      nci       = "*"
      plmn_id   = "123.45"
      qci       = 0
      snssai    = "0"
      tac       = "abc1"
      tac_5_g   = "*"
      version   = "any"
    }
    percentage      = 0
    periodic_recalc = true
    priority        = 1
    rule_id         = 1
  }
}
