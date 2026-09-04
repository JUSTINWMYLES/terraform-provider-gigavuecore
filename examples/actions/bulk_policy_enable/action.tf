action "gigavuecore_bulk_policy_enable" "example" {
  config {
    enabled    = true
    policy_ids = ["example"]
    type       = "all"
  }
}
