---
page_title: "gigavuecore_update_snmp_throttle_config_system_snmp_throttle Action - gigavuecore"
subcategory: ""
description: |-
  Update Snmp Throttle config
---

# gigavuecore_update_snmp_throttle_config_system_snmp_throttle Action

Update Snmp Throttle config

## Example Usage

```terraform
action "gigavuecore_update_snmp_throttle_config_system_snmp_throttle" "example" {
  config {
    cluster_id = "example"
    throttle_config_details = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `throttle_config_details` (List(Dynamic), optional) - list of SNMP Throttle Config Details on the node
