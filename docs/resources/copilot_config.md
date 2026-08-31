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
```

## Schema

### Arguments

The following arguments are supported:

* `api_key` (String, optional) - API key for Copilot
* `bedrock_config` (Attributes, optional) - Bedrock specific configuration (see [below for nested schema](#nestedatt--bedrock_config))
* `enabled` (Boolean, required) - Enable/Disable the copilot feature
* `provider_` (String, optional) - Copilot provider name. Bedrock
* `server_url` (String, required) - Copilot VM URL
* `test_connection` (Boolean, optional) - To decide whether to do connection validation on update

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `connection_status` (Attributes, computed) - Connection Status Domain Class (see [below for nested schema](#nestedatt--connection_status))
* `id` (String, computed)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

