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
  api_key        = null
  bedrock_config = {}
  enabled        = null
  id             = null
  provider_      = null
  server_url     = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `api_key` (String, optional) - API key for Copilot
* `bedrock_config` (Attributes, optional) - Bedrock specific configuration (see [below for nested schema](#nestedatt--bedrock_config))
* `enabled` (Boolean, required) - Enable/Disable the copilot feature
* `id` (String, required)
* `provider_` (String, optional) - Copilot provider name. Bedrock
* `server_url` (String, required) - Copilot VM URL

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `api_key` (String, computed) - API key for Copilot
* `bedrock_config` (Attributes, computed) - Bedrock specific configuration (see [below for nested schema](#nestedatt--bedrock_config))
* `connection_status` (Attributes, computed) - Connection Status Domain Class (see [below for nested schema](#nestedatt--connection_status))
* `provider_` (String, computed) - Copilot provider name. Bedrock

<a id="nestedatt--bedrock_config"></a>
### Nested Schema for `bedrock_config`

Required:

* `access_key_id` (String) - AWS Access Key ID for Bedrock service
* `region` (String) - AWS region for Bedrock service
* `secret_access_key` (String) - AWS Secret Access Key for Bedrock service
<a id="nestedatt--connection_status"></a>
### Nested Schema for `connection_status`

Read-Only:

* `message` (String) - Copilot Connection status message
* `status` (String) - Copilot Connection status code reason string
* `status_code` (Number)

