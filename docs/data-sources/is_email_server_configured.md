---
page_title: "gigavuecore_is_email_server_configured Data Source - gigavuecore"
subcategory: ""
description: |-
  Returns whether email server has been configured, or not
---

# gigavuecore_is_email_server_configured Data Source

Returns whether email server has been configured, or not

~> **Note:** This data source is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
data "gigavuecore_is_email_server_configured" "example" {
}
```
