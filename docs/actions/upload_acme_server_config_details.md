---
page_title: "gigavuecore_upload_acme_server_config_details Action - gigavuecore"
subcategory: ""
description: |-
  upload ACME server details
---

# gigavuecore_upload_acme_server_config_details Action

upload ACME server details

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_upload_acme_server_config_details" "example" {
  config {
    acme_server_url = "example"
    alias           = "example"
    certificate     = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `acme_server_url` (String, required) - ACME server url
* `alias` (String, required) - ACME server alias
* `certificate` (String, required) - User uploaded file, only .crt/.cert or .pem format is supported


