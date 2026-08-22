---
page_title: "gigavuecore_redefine_snmp_throttle_config_system_snmp_throttle Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Snmp Throttle config
---

# gigavuecore_redefine_snmp_throttle_config_system_snmp_throttle Action

Redefine Snmp Throttle config

## Example Usage

```terraform
action "gigavuecore_redefine_snmp_throttle_config_system_snmp_throttle" "example" {
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
