action "gigavuecore_redefine_netflow_exporter_filter" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    rules = [{
      pass_rules = [{
        matches = ["example"]
        rule_id = 1
      }]
    }]
  }
}
