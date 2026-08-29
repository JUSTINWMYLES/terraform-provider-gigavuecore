---
page_title: "gigavuecore_create_sip_whitelist_entries Action - gigavuecore"
subcategory: ""
description: |-
  Create SIP Whitelist Entries
---

# gigavuecore_create_sip_whitelist_entries Action

Create SIP Whitelist Entries

## Example Usage

```terraform
action "gigavuecore_create_sip_whitelist_entries" "example" {
  config {
    alias = "example"
    entries = [{
      active_sessions = 0
      caller_id       = "example"
      id_range = {
        value     = "example"
        value_max = "example"
      }
      ip_address = {
        value = "example"
      }
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target SIP Whitelist
* `entries` (Attributes List, required) (see [below for nested schema](#nestedatt--entries))

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Optional:

* `active_sessions` (Number) - Number of active sessions
* `caller_id` (String) - sip caller id
* `id_range` (Attributes) - range of values from value to valueMax (see [below for nested schema](#nestedatt--entries--id_range))
* `ip_address` (Attributes) (see [below for nested schema](#nestedatt--entries--ip_address))
<a id="nestedatt--entries--id_range"></a>
### Nested Schema for `entries.id_range`

Required:

* `value` (String)
* `value_max` (String)
<a id="nestedatt--entries--ip_address"></a>
### Nested Schema for `entries.ip_address`

Required:

* `value` (String) - ipv4 or ipv6

