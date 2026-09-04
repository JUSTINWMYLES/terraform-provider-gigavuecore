---
page_title: "gigavuecore_snmp_trap_receiver Resource - gigavuecore"
subcategory: ""
description: |-
  Add External Trap Receiver
---

# gigavuecore_snmp_trap_receiver Resource

Add External Trap Receiver

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_snmp_trap_receiver.example {alias}
```
