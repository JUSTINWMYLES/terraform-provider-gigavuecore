resource "gigavuecore_tacacs_server" "example" {
  auth_type      = "ascii"
  cluster_id     = "example"
  enabled        = true
  port           = 0
  retries        = 0
  secret_key     = "example"
  server_address = "example"
  timeout        = 0
}
