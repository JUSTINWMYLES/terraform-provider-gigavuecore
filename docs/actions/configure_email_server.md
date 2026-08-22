---
page_title: "gigavuecore_configure_email_server Action - gigavuecore"
subcategory: ""
description: |-
  Configure Email Server
---

# gigavuecore_configure_email_server Action

Configure Email Server

## Example Usage

```terraform
action "gigavuecore_configure_email_server" "example" {
  config {
    email_host = "example"
    enable_smtp_auth = true
    from = "example"
    password = "example"
    port = 1
    update_secret = true
    user_name = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `email_host` (String, optional) - Address of the SMTP Server
* `enable_smtp_auth` (Bool, optional) - Enable/Disable SMTP Authentication
* `from` (String, optional) - From address for the email
* `password` (String, optional) - Password for the SMTP Server
* `port` (Number, optional) - SMTP Server Port
* `update_secret` (Bool, optional) - updateSecret
* `user_name` (String, optional) - Username for the SMTP Server
