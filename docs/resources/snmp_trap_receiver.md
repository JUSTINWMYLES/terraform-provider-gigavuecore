---
page_title: "gigavuecore_snmp_trap_receiver Resource - gigavuecore"
subcategory: ""
description: |-
  Get External Trap Receiver
---

# gigavuecore_snmp_trap_receiver Resource

Get External Trap Receiver

## Example Usage

```terraform
resource "gigavuecore_snmp_trap_receiver" "example" {
  alias          = "example"
  auth_password  = "example"
  auth_protocol  = "SHA"
  community      = "example"
  ip_address     = "example"
  priv_password  = "example"
  priv_protocol  = "DES"
  security_level = "noAuthNoPriv"
  snmp_port      = 0
  snmp_retries   = 0
  snmp_timeout   = 0
  snmp_version   = "v2c"
  user_name      = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the External Trap Receiver
* `auth_password` (String, optional) - Authentication Password
* `auth_protocol` (String, optional) - Authentication Protocol
* `community` (String, optional) - Community name for the transaction with the remote system
* `ip_address` (String, required) - IP Address of the External Trap Receiver
* `priv_password` (String, optional) - Private Password
* `priv_protocol` (String, optional) - Private Protocol
* `security_level` (String, optional) - Security Level
* `snmp_port` (Number, optional) - Destination port number
* `snmp_retries` (Number, optional) - Number of retries
* `snmp_timeout` (Number, optional) - Timeout (in milliseconds)
* `snmp_version` (String, required) - SNMP Version
* `user_name` (String, optional) - Username

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_snmp_trap_receiver.example {alias}
```
