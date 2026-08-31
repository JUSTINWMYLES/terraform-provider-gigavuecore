action "gigavuecore_update_map_flow_whitelist_overlap_rule" "example" {
  config {
    alias        = "example"
    body_rule_id = 1
    flow5_g = {
      dnn                 = "example"
      type                = "example"
      whitelist_databases = [ "example" ]
    }
    gtp = {
      apn                 = "example"
      interface           = "Gn"
      type                = "example"
      version             = "v1"
      whitelist_databases = [ "example" ]
    }
    rule_id = "example"
    sip = {
      type = "all"
    }
  }
}
