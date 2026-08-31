action "gigavuecore_create_sip_whitelist_entries" "example" {
  config {
    alias = "example"
    entries = [{
      active_sessions = 0
      caller_id       = "example"
      id_range = {
        value     = "example"
        value_max = "example"
      }
      ip_address = {
        value = "example"
      }
    }]
  }
}
