---
page_title: "gigavuecore_radius_server Resource - gigavuecore"
subcategory: ""
description: |-
  Find RADIUS Server by address
---

# gigavuecore_radius_server Resource

Find RADIUS Server by address

## Example Usage

```terraform
resource "gigavuecore_radius_server" "example" {
  enabled = null
  port = null
  retries = null
  secret_key = null
  server_address = null
  timeout = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `enabled` (Bool, optional)
* `port` (Number, optional)
* `retries` (Number, optional) - value of 0 disables retries. Defaults to the value defined in the RadiusServerDefaults
* `secret_key` (String, required) - if not included, defaults to the value defined in the RadiusServerDefaults
* `server_address` (String, required) - ipv4 or ipv6 or hostname
* `timeout` (Number, optional) - in seconds. Defaults to the value defined in the RadiusServerDefaults

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `enabled` (Bool, computed)
* `port` (Number, computed)
* `retries` (Number, computed) - value of 0 disables retries. Defaults to the value defined in the RadiusServerDefaults
* `timeout` (Number, computed) - in seconds. Defaults to the value defined in the RadiusServerDefaults

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_radius_server.example {server_address}
```
