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
  auth_type      = "ascii"
  cluster_id     = "example"
  enabled        = true
  port           = 0
  retries        = 0
  secret_key     = "example"
  server_address = "example"
  timeout        = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `auth_type` (String, required) - Specify whether this TACACS+ server uses ASCII or PAP authentication
* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Boolean, optional)
* `port` (Number, optional)
* `retries` (Number, optional) - value of 0 disables retries. Defaults to the value defined in the TacacsServerDefaults
* `secret_key` (String, required) - if not included, defaults to the value defined in the TacacsServerDefaults
* `server_address` (String, required) - ipv4 or ipv6 or hostname
* `timeout` (Number, optional) - in seconds. Defaults to the value defined in the TacacsServerDefaults

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tacacs_server.example {server_address}/{cluster_id}
```
