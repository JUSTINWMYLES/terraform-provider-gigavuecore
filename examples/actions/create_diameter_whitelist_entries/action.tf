action "gigavuecore_create_diameter_whitelist_entries" "example" {
  config {
    alias = "example"
    entries = [{
      active_sessions = 0
      user_name       = "12345678901234"
    }]
  }
}
