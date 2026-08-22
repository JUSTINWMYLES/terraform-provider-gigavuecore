resource "gigavuecore_radius_server" "example" {
  enabled = true
  port = 1
  retries = 1
  secret_key = "example"
  server_address = "example"
  timeout = 1
}
