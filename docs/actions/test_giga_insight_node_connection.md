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
    provider_ = "example"
    proxy_url = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `bedrock_config` (Attributes, optional) - AWS Bedrock provider configuration (see [below for nested schema](#nestedatt--bedrock_config))
* `google_config` (Attributes, optional) - Google Vertex AI provider configuration (see [below for nested schema](#nestedatt--google_config))
* `node_id` (String, required)
* `open_ai_config` (Attributes, optional) - OpenAI provider configuration (see [below for nested schema](#nestedatt--open_ai_config))
* `private_config` (Attributes, optional) - Private (Azure AI Foundry) provider configuration (see [below for nested schema](#nestedatt--private_config))
* `provider_` (String, required) - The LLM provider type
* `proxy_url` (String, optional) - Optional proxy URL for the connection

<a id="nestedatt--bedrock_config"></a>
### Nested Schema for `bedrock_config`

Required:

* `access_key_id` (String) - AWS access key ID
* `region` (String) - AWS region
* `secret_access_key` (String) - AWS secret access key

<a id="nestedatt--google_config"></a>
### Nested Schema for `google_config`

Required:

* `project_id` (String) - Google Cloud project ID
* `region` (String) - Google Cloud region
* `service_account_key_json` (Dynamic) - Google service account key JSON

<a id="nestedatt--open_ai_config"></a>
### Nested Schema for `open_ai_config`

Required:

* `api_key` (String) - OpenAI API key

<a id="nestedatt--private_config"></a>
### Nested Schema for `private_config`

Required:

* `api_key` (String) - API key for the private endpoint
* `endpoint_url` (String) - Private LLM endpoint URL

