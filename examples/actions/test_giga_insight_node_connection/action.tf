action "gigavuecore_test_giga_insight_node_connection" "example" {
  config {
    bedrock_config = {
      access_key_id     = "example"
      region            = "example"
      secret_access_key = "example"
    }
    google_config = {
      project_id               = "example"
      region                   = "example"
      service_account_key_json = "example"
    }
    node_id = "example"
    open_ai_config = {
      api_key = "example"
    }
    private_config = {
      api_key      = "example"
      endpoint_url = "example"
    }
    provider_ = "openAI"
    proxy_url = "example"
  }
}
