action "gigavuecore_copy_rules" "example" {
  config {
    configs = [{
      flow_alias            = "example"
      is_all_rules          = true
      policy_id_or_alias    = "example"
      policy_updated_time   = 0
      rule_ids              = [0]
      source_and_rule_alias = "example"
      sub_flow_alias        = "example"
    }]
    rule_category = "SOURCE"
    rule_type     = "byRule"
  }
}
