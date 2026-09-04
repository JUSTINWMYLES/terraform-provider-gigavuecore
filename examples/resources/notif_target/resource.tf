resource "gigavuecore_notif_target" "example" {
  cluster_id = "example"
  enabled    = true
  host       = "example"
  notify_config = {
    auth_key      = "example-value"
    auth_protocol = "md5"
    community     = "example"
    engine_id     = "example"
    port          = 0
    priv_key      = "example-value"
    priv_protocol = "des"
    v3_user       = "example"
    version       = "v1"
  }
  notify_type = "trap"
}
