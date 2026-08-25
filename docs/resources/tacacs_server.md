---
page_title: "gigavuecore_tacacs_server Resource - gigavuecore"
subcategory: ""
description: |-
  Find TACACS+ Server by address
---

# gigavuecore_tacacs_server Resource

Find TACACS+ Server by address

## Example Usage

```terraform
resource "gigavuecore_tacacs_server" "example" {
  auth_type      = null
  enabled        = null
  port           = null
  retries        = null
  secret_key     = null
  server_address = null
  timeout        = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `auth_type` (String, required) - Specify whether this TACACS+ server uses ASCII or PAP authentication
* `enabled` (Boolean, optional)
* `port` (Number, optional)
* `retries` (Number, optional) - value of 0 disables retries. Defaults to the value defined in the TacacsServerDefaults
* `secret_key` (String, required) - if not included, defaults to the value defined in the TacacsServerDefaults
* `server_address` (String, required) - ipv4 or ipv6 or hostname
* `timeout` (Number, optional) - in seconds. Defaults to the value defined in the TacacsServerDefaults

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `enabled` (Boolean, computed)
* `port` (Number, computed)
* `retries` (Number, computed) - value of 0 disables retries. Defaults to the value defined in the TacacsServerDefaults
* `timeout` (Number, computed) - in seconds. Defaults to the value defined in the TacacsServerDefaults


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tacacs_server.example {server_address}
```
