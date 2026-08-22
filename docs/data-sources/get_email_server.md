---
page_title: "gigavuecore_get_email_server Data Source - gigavuecore"
subcategory: ""
description: |-
  Get system email server
---

# gigavuecore_get_email_server Data Source

Get system email server

## Example Usage

```terraform
data "gigavuecore_get_email_server" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `domain_name` (String, computed) - Domain Name
* `enable_auto_support_notif` (Bool, computed) - Enable Auto Support Notifications
* `enable_smtp_auth` (Bool, computed) - Enable SMTP auth
* `include_hostname` (Bool, computed) - Include Hostname
* `mail_hub_port` (Number, computed)
* `password` (String, computed) - SMTP Password
* `return_address` (String, computed) - Return Address
* `smtp_server` (String, computed) - SMTP Server
* `username` (String, computed) - SMTP Username

