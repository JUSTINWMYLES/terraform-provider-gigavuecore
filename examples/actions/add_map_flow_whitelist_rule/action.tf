action "gigavuecore_add_map_flow_whitelist_rule" "example" {
  config {
    alias = "example"
    flow5_g = {
      dnn                 = "example"
      type                = "example"
      whitelist_databases = ["example"]
    }
    gtp = {
      apn                 = "example"
      interface           = "Gn"
      type                = "example"
      version             = "v1"
      whitelist_databases = ["example"]
    }
    rule_id = 1
    sip = {
      type = "all"
    }
  }
}
