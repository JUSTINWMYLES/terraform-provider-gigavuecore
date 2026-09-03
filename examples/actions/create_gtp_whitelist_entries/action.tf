action "gigavuecore_create_gtp_whitelist_entries" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    entries = [{
      active_sessions = 0
      imsi            = "12345678901234"
      ran             = "123.45.0xabcd"
    }]
  }
}
