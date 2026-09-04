---
page_title: "gigavuecore_create_api_token_tokens_create Ephemeral Resource - gigavuecore"
subcategory: ""
description: |-
  Create token for FM Users
---

# gigavuecore_create_api_token_tokens_create Ephemeral Resource

Create token for FM Users

~> **Note:** Ephemeral resources are only available within the context of a single Terraform operation and are never persisted to state or plan files.

~> **Note:** This ephemeral resource is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
ephemeral "gigavuecore_create_api_token_tokens_create" "example" {
  authentication_type = "example"
  created_by          = "example"
  created_ts          = 0
  expiry_time         = "example"
  expiry_ts           = "example"
  groups              = ["example"]
  token               = "example"
  token_id            = "example"
  token_name          = "example"
  usage_count         = 0
  username            = "example"
}
```
## Schema

### Arguments

The following arguments are supported:

* `authentication_type` (String, optional) - Authentication Type local/radius/tacacs+/external
* `created_by` (String, optional) - FM user who has created the token
* `created_ts` (Number, optional) - Token creation timestamp
* `expiry_time` (String, optional) - Expiry Time in number of days. Default would be 30 and Maximum of 105 days.
* `expiry_ts` (String, optional) - Expiry Timestamp
* `groups` (List of String, optional) - FM User Groups
* `token` (String, optional) - FM generated JWT token for FM REST API access
* `token_id` (String, optional) - Random alpha-numeric 64 digit string
* `token_name` (String, optional) - User defined token name
* `usage_count` (Number, optional) - Token Tracking count. Usagecount keep track of number times token was used to access API's
* `username` (String, optional) - FM Username


