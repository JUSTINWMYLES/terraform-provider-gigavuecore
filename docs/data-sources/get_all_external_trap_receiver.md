---
page_title: "gigavuecore_get_all_external_trap_receiver Data Source - gigavuecore"
subcategory: ""
description: |-
  Get All External Trap Receiver
---

# gigavuecore_get_all_external_trap_receiver Data Source

Get All External Trap Receiver

## Example Usage

```terraform
data "gigavuecore_get_all_external_trap_receiver" "example" {
  alias          = null
  auth_protocol  = null
  community      = null
  ip_address     = null
  page           = null
  priv_protocol  = null
  security_level = null
  snmp_port      = null
  snmp_retries   = null
  snmp_timeout   = null
  snmp_version   = null
  sort           = null
  user_name      = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - alias
* `auth_protocol` (String, optional) - authProtocol
* `community` (String, optional) - community
* `ip_address` (String, optional) - ipAddress
* `page` (String, optional) - page
* `priv_protocol` (String, optional) - privProtocol
* `security_level` (String, optional) - securityLevel
* `snmp_port` (String, optional) - snmpPort
* `snmp_retries` (String, optional) - snmpRetries
* `snmp_timeout` (String, optional) - snmpTimeout
* `snmp_version` (String, optional) - snmpVersion
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `user_name` (String, optional) - userName

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - Alias of the External Trap Receiver
* `auth_password` (String) - Authentication Password
* `auth_protocol` (String) - Authentication Protocol
* `community` (String) - Community name for the transaction with the remote system
* `ip_address` (String) - IP Address of the External Trap Receiver
* `priv_password` (String) - Private Password
* `priv_protocol` (String) - Private Protocol
* `security_level` (String) - Security Level
* `snmp_port` (Number) - Destination port number
* `snmp_retries` (Number) - Number of retries
* `snmp_timeout` (Number) - Timeout (in milliseconds)
* `snmp_version` (String) - SNMP Version
* `user_name` (String) - Username

