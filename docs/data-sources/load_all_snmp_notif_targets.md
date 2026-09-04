---
page_title: "gigavuecore_load_all_snmp_notif_targets Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all SNMP Notification Targets
---

# gigavuecore_load_all_snmp_notif_targets Data Source

Load all SNMP Notification Targets

## Example Usage

```terraform
data "gigavuecore_load_all_snmp_notif_targets" "example" {
  cluster_id = "example"
  page       = "example"
  sort       = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `enabled` (Boolean) - temporarily enable/disable the notification destination
* `host` (String) - ipv4 or ipv6 or domain name
* `notify_config` (Attributes) - Notification Target configuration for specific Notification type (Trap/Inform) (see [below for nested schema](#nestedatt--items--notify_config))
* `notify_type` (String) - SNMP notification type to use

<a id="nestedatt--items--notify_config"></a>
### Nested Schema for `items.notify_config`

Read-Only:

* `auth_key` (String) - authentication password. required with 'v3user'
* `auth_protocol` (String) - authentication hash algorithm. required with 'v3user'
* `community` (String) - required when when 'version' is 'v2c'
* `engine_id` (String) - remote engineID. only valid with notifyType 'inform' and 'version' v3
* `port` (Number)
* `priv_key` (String) - privacy password
* `priv_protocol` (String) - privacy encryption
* `v3_user` (String) - required when when 'version' is 'v3'
* `version` (String) - SNMP version to use. v1 is only valid for traps. for v3, user name should be provided

