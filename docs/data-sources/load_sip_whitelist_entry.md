---
page_title: "gigavuecore_load_sip_whitelist_entry Data Source - gigavuecore"
subcategory: ""
description: |-
  Check if active sessions for  Whitelist Entry exists
---

# gigavuecore_load_sip_whitelist_entry Data Source

Check if active sessions for  Whitelist Entry exists

## Example Usage

```terraform
data "gigavuecore_load_sip_whitelist_entry" "example" {
  alias = null
  caller_id = null
  ip_address = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target SIP Whitelist
* `caller_id` (String, optional) - sip caller id
* `ip_address` (Object({value}), optional) - ipaddress based whitelist entry being queried. Example:ipAddress=(value:1.1.1.1)
  * `value` (String, computed) - ipv4 or ipv6

### Attributes

In addition to all arguments above, the following attributes are exported:

* `active_sessions` (Number, computed) - Number of active sessions
* `caller_id` (String, computed) - sip caller id
* `id_range` (Object({value, value_max}), computed) - range of values from value to valueMax
  * `value` (String, computed)
  * `value_max` (String, computed)
* `ip_address` (Object({value}), computed) - ipaddress based whitelist entry being queried. Example:ipAddress=(value:1.1.1.1)
  * `value` (String, computed) - ipv4 or ipv6

