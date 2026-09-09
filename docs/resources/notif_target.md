---
page_title: "gigavuecore_notif_target Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new SNMP Notification Target
---

# gigavuecore_notif_target Resource

Create a new SNMP Notification Target

## Example Usage

```terraform
resource "gigavuecore_notif_target" "example" {
  cluster_id = "example"
  enabled    = true
  host       = "example"
  notify_config = {
    auth_key      = "example-value"
    auth_protocol = "md5"
    community     = "example"
    engine_id     = "example"
    port          = 0
    priv_key      = "example-value"
    priv_protocol = "des"
    v3_user       = "example"
    version       = "v1"
  }
  notify_type = "trap"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Boolean, optional) - temporarily enable/disable the notification destination
* `host` (String, required) - ipv4 or ipv6 or domain name
* `notify_config` (Attributes, optional) - Notification Target configuration for specific Notification type (Trap/Inform) (see [below for nested schema](#nestedatt--notify_config))
* `notify_type` (String, optional) - SNMP notification type to use

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `notif_target_address` (String, computed) - address of the target SNMP Notification Target

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
terraform import gigavuecore_notif_target.example {notif_target_address}/{cluster_id}
```
