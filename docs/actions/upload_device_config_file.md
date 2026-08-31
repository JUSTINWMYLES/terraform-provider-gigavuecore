---
page_title: "gigavuecore_upload_device_config_file Action - gigavuecore"
subcategory: ""
description: |-
  upload device config file
---

# gigavuecore_upload_device_config_file Action

upload device config file

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_upload_device_config_file" "example" {
  config {
    file = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `file` (String, required) - Attached Config file. In text format


