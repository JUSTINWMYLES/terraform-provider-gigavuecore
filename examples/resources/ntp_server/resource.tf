resource "gigavuecore_ntp_server" "example" {
  cluster_id  = "example"
  enabled     = true
  key_enabled = true
  key_number  = 0
  preferred   = true
  server      = "example"
  version     = "example"
}
