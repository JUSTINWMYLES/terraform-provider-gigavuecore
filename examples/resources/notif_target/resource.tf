resource "gigavuecore_notif_target" "example" {
  enabled = true
  host    = "example"
  notify_config = {
    auth_key      = "example"
    auth_protocol = "example"
    community     = "example"
    engine_id     = "example"
    port          = 1
    priv_key      = "example"
    priv_protocol = "example"
    v3_user       = "example"
    version       = "example"
  }
  notify_type = "example"
}
