---
page_title: "gigavuecore_configure_email_server Action - gigavuecore"
subcategory: ""
description: |-
  Configure Email Server
---

# gigavuecore_configure_email_server Action

Configure Email Server

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (password), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_configure_email_server" "example" {
  config {
    email_host       = "example"
    enable_smtp_auth = true
    from             = "example"
    password         = "example"
    port             = 0
    update_secret    = true
    user_name        = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `email_host` (String, optional) - Address of the SMTP Server
* `enable_smtp_auth` (Boolean, optional) - Enable/Disable SMTP Authentication
* `from` (String, optional) - From address for the email
* `password` (String, optional) - Password for the SMTP Server
* `port` (Number, optional) - SMTP Server Port
* `update_secret` (Boolean, optional) - updateSecret
* `user_name` (String, optional) - Username for the SMTP Server


