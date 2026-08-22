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
  alias = null
  auth_protocol = null
  community = null
  ip_address = null
  page = null
  priv_protocol = null
  security_level = null
  snmp_port = null
  snmp_retries = null
  snmp_timeout = null
  snmp_version = null
  sort = null
  user_name = null
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

* `items` (List(Object({alias, auth_password, auth_protocol, community, ip_address, priv_password, priv_protocol, security_level, snmp_port, snmp_retries, snmp_timeout, snmp_version, user_name})), computed)

