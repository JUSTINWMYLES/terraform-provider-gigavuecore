resource "gigavuecore_tacacs_server" "example" {
  auth_type      = "example"
  enabled        = true
  port           = 1
  retries        = 1
  secret_key     = "example"
  server_address = "example"
  timeout        = 1
}
