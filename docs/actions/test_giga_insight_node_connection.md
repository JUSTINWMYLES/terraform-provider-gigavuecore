---
page_title: "gigavuecore_test_giga_insight_node_connection Action - gigavuecore"
subcategory: ""
description: |-
  Test LLM provider connection via a GigaInsight Node
---

# gigavuecore_test_giga_insight_node_connection Action

Test LLM provider connection via a GigaInsight Node

## Example Usage

```terraform
action "gigavuecore_test_giga_insight_node_connection" "example" {
  config {
    bedrock_config = null
    google_config  = null
    node_id        = "example"
    open_ai_config = null
    private_config = null
    provider_      = "example"
    proxy_url      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `bedrock_config` (Dynamic, optional) - AWS Bedrock provider configuration
* `google_config` (Dynamic, optional) - Google Vertex AI provider configuration
* `node_id` (String, required)
* `open_ai_config` (Dynamic, optional) - OpenAI provider configuration
* `private_config` (Dynamic, optional) - Private (Azure AI Foundry) provider configuration
* `provider_` (String, required) - The LLM provider type
* `proxy_url` (String, optional) - Optional proxy URL for the connection


