---
page_title: "gigavuecore_load_snmp_server_config Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Snmp Server config
---

# gigavuecore_load_snmp_server_config Data Source

Load Snmp Server config

## Example Usage

```terraform
data "gigavuecore_load_snmp_server_config" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `community_config` (Attributes, computed) - SNMP Server Community Strings config (see [below for nested schema](#nestedatt--community_config))
* `notify_config` (Attributes, computed) - Node SNMP Server Notification settings (see [below for nested schema](#nestedatt--notify_config))
* `snmp_throttle_config` (Attributes, computed) - SNMP Throttle Configuration (see [below for nested schema](#nestedatt--snmp_throttle_config))
* `snmp_v3_config` (Attributes, computed) - SNMP Server v3 config (see [below for nested schema](#nestedatt--snmp_v3_config))
* `system_config` (Attributes, computed) - SNMP Server System level config (see [below for nested schema](#nestedatt--system_config))

<a id="nestedatt--community_config"></a>
### Nested Schema for `community_config`

Read-Only:

* `community_strings` (List of String) - community string\[s\] used to connect to this node using SNMP. The default value is 'public'. If 'enableMultiCommunity' is enabled, multiple community strings for the node are allowed
* `enable_community_auth` (Boolean) - turn on community-based authentication for the system
* `enable_community_auth_v1` (Boolean) - turn on community-based authentication for the SNMP v1
* `enable_multi_community` (Boolean) - allow configuration of multiple communities
<a id="nestedatt--notify_config"></a>
### Nested Schema for `notify_config`

Read-Only:

* `notify_events` (Set of String) - The set of notification event types
* `notify_set` (String) - When 'notifySet' is 'select', notifyEvents represents the subset of notification event types activated for dispatch. When set to 'none', effectively disables SNMP notifications from the node
* `targets` (Attributes List) - list of notification destinations (see [below for nested schema](#nestedatt--notify_config--targets))
<a id="nestedatt--notify_config--targets"></a>
### Nested Schema for `notify_config.targets`

Read-Only:

* `enabled` (Boolean) - temporarily enable/disable the notification destination
* `host` (String) - ipv4 or ipv6 or domain name
* `notify_config` (Attributes) - Notification Target configuration for specific Notification type (Trap/Inform) (see [below for nested schema](#nestedatt--notify_config--targets--notify_config))
* `notify_type` (String) - SNMP notification type to use
<a id="nestedatt--notify_config--targets--notify_config"></a>
### Nested Schema for `notify_config.targets.notify_config`

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
<a id="nestedatt--snmp_throttle_config"></a>
### Nested Schema for `snmp_throttle_config`

Read-Only:

* `throttle_config_details` (Attributes List) - list of SNMP Throttle Config Details on the node (see [below for nested schema](#nestedatt--snmp_throttle_config--throttle_config_details))
<a id="nestedatt--snmp_throttle_config--throttle_config_details"></a>
### Nested Schema for `snmp_throttle_config.throttle_config_details`

Read-Only:

* `interval` (Number) - Time interval at which the throttle should occur (in seconds)
* `notify_set` (String) - When 'notifySet' is 'select', throttleEvents represents the subset of event types for SNMP throttling. When set to 'none', effectively disables SNMP throttle from the node.
* `report_threshold` (Number) - Minimum count threshold to send the throttle report
* `throttle_events` (Set of String) - The set of notification event types
<a id="nestedatt--snmp_v3_config"></a>
### Nested Schema for `snmp_v3_config`

Read-Only:

* `snmp_v3_users` (Attributes List) - list of SNMPv3 users on the node (see [below for nested schema](#nestedatt--snmp_v3_config--snmp_v3_users))
<a id="nestedatt--snmp_v3_config--snmp_v3_users"></a>
### Nested Schema for `snmp_v3_config.snmp_v3_users`

Read-Only:

* `auth_key` (String)
* `auth_protocol` (String)
* `enabled` (Boolean)
* `priv_key` (String)
* `priv_protocol` (String)
* `read_only` (Boolean) - This property is introduced for GUI's purpose to restrict the edit/delete on the SNMPv3 user used by FM. If this property doesn't exist, consider it false. And this property is ignored if it exists in the create/update spec.
* `username` (String)
<a id="nestedatt--system_config"></a>
### Nested Schema for `system_config`

Read-Only:

* `enabled` (Boolean) - Enables SNMP Server on the node
* `engine_id` (String) - local EngineID
* `port` (Number)
* `sys_contact` (String) - MIB-II 'sysContact'
* `sys_descr` (String) - MIB-II 'sysDescr'
* `sys_location` (String) - MIB-II 'sysLocation'
* `sys_name` (String) - MIB-II 'sysName'

