---
page_title: "gigavuecore_copilot_config Resource - gigavuecore"
subcategory: ""
description: |-
  Load the copilot configuration
---

# gigavuecore_copilot_config Resource

Load the copilot configuration

## Example Usage

```terraform
resource "gigavuecore_copilot_config" "example" {
  api_key = null
  bedrock_config = {}
  enabled = null
  id = null
  provider_ = null
  server_url = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `api_key` (String, optional) - API key for Copilot
* `bedrock_config` (Object({access_key_id, region, secret_access_key}), optional) - Bedrock specific configuration
  * `access_key_id` (String, required) - AWS Access Key ID for Bedrock service
  * `region` (String, required) - AWS region for Bedrock service
  * `secret_access_key` (String, required) - AWS Secret Access Key for Bedrock service
* `enabled` (Bool, required) - Enable/Disable the copilot feature
* `id` (String, required)
* `provider_` (String, optional) - Copilot provider name. Bedrock
* `server_url` (String, required) - Copilot VM URL

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `api_key` (String, computed) - API key for Copilot
* `bedrock_config` (Object({access_key_id, region, secret_access_key}), computed) - Bedrock specific configuration
  * `access_key_id` (String, required) - AWS Access Key ID for Bedrock service
  * `region` (String, required) - AWS region for Bedrock service
  * `secret_access_key` (String, required) - AWS Secret Access Key for Bedrock service
* `connection_status` (Object({message, status, status_code}), computed) - Connection Status Domain Class
  * `message` (String, computed) - Copilot Connection status message
  * `status` (String, computed) - Copilot Connection status code reason string
  * `status_code` (Number, computed)
* `provider_` (String, computed) - Copilot provider name. Bedrock

