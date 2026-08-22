resource "gigavuecore_ntp_server" "example" {
  enabled = true
  key_enabled = true
  key_number = 1
  preferred = true
  server = "example"
  version = "example"
}
