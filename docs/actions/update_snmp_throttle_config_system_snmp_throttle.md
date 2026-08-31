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
    throttle_config_details = [{
      interval         = 1
      notify_set       = "all"
      report_threshold = 0
      throttle_events  = [ "secondflashboot" ]
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `throttle_config_details` (Attributes List, optional) - list of SNMP Throttle Config Details on the node (see [below for nested schema](#nestedatt--throttle_config_details))

<a id="nestedatt--throttle_config_details"></a>
### Nested Schema for `throttle_config_details`

Optional:

* `interval` (Number) - Time interval at which the throttle should occur (in seconds)
* `notify_set` (String) - When 'notifySet' is 'select', throttleEvents represents the subset of event types for SNMP throttling. When set to 'none', effectively disables SNMP throttle from the node.
* `report_threshold` (Number) - Minimum count threshold to send the throttle report
* `throttle_events` (Set of String) - The set of notification event types

