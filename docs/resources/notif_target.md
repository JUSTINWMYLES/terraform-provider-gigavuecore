---
page_title: "gigavuecore_notif_target Resource - gigavuecore"
subcategory: ""
description: |-
  Find SNMP Notification Target by address
---

# gigavuecore_notif_target Resource

Find SNMP Notification Target by address

## Example Usage

```terraform
resource "gigavuecore_notif_target" "example" {
  enabled       = null
  host          = null
  notify_config = {}
  notify_type   = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `enabled` (Boolean, optional) - temporarily enable/disable the notification destination
* `host` (String, required) - ipv4 or ipv6 or domain name
* `notify_config` (Attributes, optional) - Notification Target configuration for specific Notification type (Trap/Inform) (see [below for nested schema](#nestedatt--notify_config))
* `notify_type` (String, optional) - SNMP notification type to use

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `enabled` (Boolean, computed) - temporarily enable/disable the notification destination
* `notif_target_address` (String, computed)
* `notify_config` (Attributes, computed) - Notification Target configuration for specific Notification type (Trap/Inform) (see [below for nested schema](#nestedatt--notify_config))
* `notify_type` (String, computed) - SNMP notification type to use

<a id="nestedatt--notify_config"></a>
### Nested Schema for `notify_config`

Optional:

* `auth_key` (String) - authentication password. required with 'v3user'
* `auth_protocol` (String) - authentication hash algorithm. required with 'v3user'
* `community` (String) - required when when 'version' is 'v2c'
* `engine_id` (String) - remote engineID. only valid with notifyType 'inform' and 'version' v3
* `port` (Number)
* `priv_key` (String) - privacy password
* `priv_protocol` (String) - privacy encryption
* `v3_user` (String) - required when when 'version' is 'v3'
* `version` (String) - SNMP version to use. v1 is only valid for traps. for v3, user name should be provided

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_notif_target.example {notif_target_address}
```
