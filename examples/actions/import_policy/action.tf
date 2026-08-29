action "gigavuecore_import_policy" "example" {
  config {
    criteria_bindings              = { "key" = "example" }
    name                           = "example"
    packet_transformation_bindings = { "key" = "example" }
    priority                       = true
    rules = [{
      criteria = [{
        criteria_name = "example"
        filters = [{
          properties = [{
            key    = "example"
            values = [ "example" ]
          }]
          type   = "example"
          values = [ "example" ]
        }]
      }]
      high_priority_drop = true
      no_expansion_tags  = [ "example" ]
      packet_transformation = {
        apps = [{
          parameters = "example"
          type       = "example"
        }]
        engine_ports = [ "example" ]
      }
      rule_name = "example"
      tags = [{
        key    = "example"
        values = [ "example" ]
      }]
      tools = [{
        ports = [ "example" ]
        type  = "example"
      }]
      type = "example"
    }]
    source_bindings = { "key" = "example" }
    sources = [{
      ports = [ "example" ]
      type  = "example"
    }]
    tags = [{
      key    = "example"
      values = [ "example" ]
    }]
  }
}
