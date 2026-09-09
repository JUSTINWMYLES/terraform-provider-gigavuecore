---
page_title: "gigavuecore_tacacs_server Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new TACACS+ Server
---

# gigavuecore_tacacs_server Resource

Create a new TACACS+ Server

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_tacacs_server.example {server_address}/{cluster_id}
```
