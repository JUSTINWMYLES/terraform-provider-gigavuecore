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
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `community_config` (Object({community_strings, enable_community_auth, enable_community_auth_v1, enable_multi_community}), computed) - SNMP Server Community Strings config
  * `community_strings` (List(String), computed) - community string\[s\] used to connect to this node using SNMP. The default value is 'public'. If 'enableMultiCommunity' is enabled, multiple community strings for the node are allowed
  * `enable_community_auth` (Bool, computed) - turn on community-based authentication for the system
  * `enable_community_auth_v1` (Bool, computed) - turn on community-based authentication for the SNMP v1
  * `enable_multi_community` (Bool, computed) - allow configuration of multiple communities
* `notify_config` (Object({notify_events, notify_set, targets}), computed) - Node SNMP Server Notification settings
  * `notify_events` (Set(String), computed) - The set of notification event types
  * `notify_set` (String, computed) - When 'notifySet' is 'select', notifyEvents represents the subset of notification event types activated for dispatch. When set to 'none', effectively disables SNMP notifications from the node
  * `targets` (List(Object({enabled, host, notify_config, notify_type})), computed) - list of notification destinations
* `snmp_throttle_config` (Object({throttle_config_details}), computed) - SNMP Throttle Configuration
  * `throttle_config_details` (List(Object({interval, notify_set, report_threshold, throttle_events})), computed) - list of SNMP Throttle Config Details on the node
* `snmp_v3_config` (Object({snmp_v3_users}), computed) - SNMP Server v3 config
  * `snmp_v3_users` (List(Object({auth_key, auth_protocol, enabled, priv_key, priv_protocol, read_only, username})), computed) - list of SNMPv3 users on the node
* `system_config` (Object({enabled, engine_id, port, sys_contact, sys_descr, sys_location, sys_name}), computed) - SNMP Server System level config
  * `enabled` (Bool, computed) - Enables SNMP Server on the node
  * `engine_id` (String, computed) - local EngineID
  * `port` (Number, computed)
  * `sys_contact` (String, computed) - MIB-II 'sysContact'
  * `sys_descr` (String, computed) - MIB-II 'sysDescr'
  * `sys_location` (String, computed) - MIB-II 'sysLocation'
  * `sys_name` (String, computed) - MIB-II 'sysName'

