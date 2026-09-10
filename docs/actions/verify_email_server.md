---
page_title: "gigavuecore_verify_email_server Action - gigavuecore"
subcategory: ""
description: |-
  Verify Email Server Configuration
---

# gigavuecore_verify_email_server Action

Verify Email Server Configuration

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_verify_email_server" "example" {
  config {
    recipients = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `recipients` (String, required) - recipients


