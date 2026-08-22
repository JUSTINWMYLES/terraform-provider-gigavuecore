---
page_title: "gigavuecore_get_all_external_trap_receiver List Resource - gigavuecore"
subcategory: ""
description: |-
  Get All External Trap Receiver
---

# gigavuecore_get_all_external_trap_receiver List Resource

Get All External Trap Receiver

## Example Usage

```terraform
list "gigavuecore_get_all_external_trap_receiver" "example" {
  provider = gigavuecore
  limit = 100
  config {
    alias = "example"
    auth_protocol = "example"
    community = "example"
    ip_address = "example"
    page = "example"
    priv_protocol = "example"
    security_level = "example"
    snmp_port = "example"
    snmp_retries = "example"
    snmp_timeout = "example"
    snmp_version = "example"
    sort = "example"
    user_name = "example"
  }
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

### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)
