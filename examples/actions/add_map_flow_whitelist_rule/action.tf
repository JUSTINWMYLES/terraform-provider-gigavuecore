action "gigavuecore_add_map_flow_whitelist_rule" "example" {
  config {
    alias = "example"
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
    rule_id = 0
    sip = {
      type = "example"
    }
  }
}
