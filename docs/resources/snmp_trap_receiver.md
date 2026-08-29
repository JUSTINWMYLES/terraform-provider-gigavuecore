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
  auth_protocol  = "example"
  community      = "example"
  ip_address     = "example"
  priv_password  = "example"
  priv_protocol  = "example"
  security_level = "example"
  snmp_port      = 0
  snmp_retries   = 0
  snmp_timeout   = 0
  snmp_version   = "example"
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `auth_password` (String, computed) - Authentication Password
* `auth_protocol` (String, computed) - Authentication Protocol
* `community` (String, computed) - Community name for the transaction with the remote system
* `priv_password` (String, computed) - Private Password
* `priv_protocol` (String, computed) - Private Protocol
* `security_level` (String, computed) - Security Level
* `snmp_port` (Number, computed) - Destination port number
* `snmp_retries` (Number, computed) - Number of retries
* `snmp_timeout` (Number, computed) - Timeout (in milliseconds)
* `user_name` (String, computed) - Username


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_snmp_trap_receiver.example {alias}
```
