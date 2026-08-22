---
page_title: "gigavuecore_redefine_snmp_server_system_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Snmp Server System config
---

# gigavuecore_redefine_snmp_server_system_config Action

Redefine Snmp Server System config

## Example Usage

```terraform
action "gigavuecore_redefine_snmp_server_system_config" "example" {
  config {
    cluster_id = "example"
    enabled = true
    engine_id = "example"
    port = 1
    sys_contact = "example"
    sys_descr = "example"
    sys_location = "example"
    sys_name = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `enabled` (Bool, optional) - Enables SNMP Server on the node
* `engine_id` (String, optional) - local EngineID
* `port` (Number, optional)
* `sys_contact` (String, optional) - MIB-II 'sysContact'
* `sys_descr` (String, optional) - MIB-II 'sysDescr'
* `sys_location` (String, optional) - MIB-II 'sysLocation'
* `sys_name` (String, optional) - MIB-II 'sysName'
