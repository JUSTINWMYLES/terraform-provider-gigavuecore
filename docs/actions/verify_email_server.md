---
page_title: "gigavuecore_verify_email_server Action - gigavuecore"
subcategory: ""
description: |-
  Verify Email Server Configuration
---

# gigavuecore_verify_email_server Action

Verify Email Server Configuration

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
