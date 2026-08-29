action "gigavuecore_update_map_flow_whitelist_overlap_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 0
    flow5_g = {
      dnn                 = "example"
      type                = "example"
      whitelist_databases = [ "example" ]
    }
    gtp = {
      apn                 = "example"
      interface           = "example"
      type                = "example"
      version             = "example"
      whitelist_databases = [ "example" ]
    }
    rule_id = "example"
    sip = {
      type = "example"
    }
  }
}
