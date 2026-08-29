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
  alias     = "example"
  caller_id = "example"
  ip_address = {
  }
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target SIP Whitelist
* `caller_id` (String, optional) - sip caller id
* `ip_address` (Attributes, optional) - ipaddress based whitelist entry being queried. Example:ipAddress=(value:1.1.1.1) (see [below for nested schema](#nestedatt--ip_address))

### Attributes

In addition to all arguments above, the following attributes are exported:

* `active_sessions` (Number, computed) - Number of active sessions
* `caller_id` (String, computed) - sip caller id
* `id_range` (Attributes, computed) - range of values from value to valueMax (see [below for nested schema](#nestedatt--id_range))
* `ip_address` (Attributes, computed) - ipaddress based whitelist entry being queried. Example:ipAddress=(value:1.1.1.1) (see [below for nested schema](#nestedatt--ip_address))

<a id="nestedatt--ip_address"></a>
### Nested Schema for `ip_address`

Read-Only:

* `value` (String) - ipv4 or ipv6
<a id="nestedatt--id_range"></a>
### Nested Schema for `id_range`

Read-Only:

* `value` (String)
* `value_max` (String)

