resource "gigavuecore_copilot_config" "example" {
  api_key = "example"
  bedrock_config = {
    access_key_id     = "example"
    region            = "example"
    secret_access_key = "example"
  }
  enabled         = true
  provider_       = "example"
  server_url      = "example"
  test_connection = true
}
