resource "gigavuecore_port_filter" "example" {
  cluster_id = "example"
  port       = "example"
  rules = {
    drop_rules = [{
      comment = "example"
      matches = [ "example" ]
      rule_id = 1
    }]
    pass_rules = [{
      comment = "example"
      matches = [ "example" ]
      rule_id = 1
    }]
  }
}
