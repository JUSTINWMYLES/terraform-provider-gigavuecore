---
page_title: "gigavuecore_modify_email_server Action - gigavuecore"
subcategory: ""
description: |-
  Modify system email server
---

# gigavuecore_modify_email_server Action

Modify system email server

## Example Usage

```terraform
action "gigavuecore_modify_email_server" "example" {
  config {
    domain_name               = "example"
    enable_auto_support_notif = true
    enable_smtp_auth          = true
    include_hostname          = true
    mail_hub_port             = 0
    password                  = "example"
    return_address            = "example"
    smtp_server               = "example"
    username                  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `domain_name` (String, optional) - Domain Name
* `enable_auto_support_notif` (Boolean, optional) - Enable Auto Support Notifications
* `enable_smtp_auth` (Boolean, optional) - Enable SMTP auth
* `include_hostname` (Boolean, optional) - Include Hostname
* `mail_hub_port` (Number, optional)
* `password` (String, optional) - SMTP Password
* `return_address` (String, optional) - Return Address
* `smtp_server` (String, optional) - SMTP Server
* `username` (String, optional) - SMTP Username


