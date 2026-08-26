action "gigavuecore_redefine_ntp_server" "example" {
  config {
    cluster_id     = "example"
    enabled        = true
    key_enabled    = true
    key_number     = 1
    preferred      = true
    server         = "example"
    server_address = "example"
    version        = "example"
  }
}
