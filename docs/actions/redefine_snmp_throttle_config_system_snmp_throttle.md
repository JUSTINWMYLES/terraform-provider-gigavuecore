---
page_title: "gigavuecore_redefine_snmp_throttle_config_system_snmp_throttle Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Snmp Throttle config
---

# gigavuecore_redefine_snmp_throttle_config_system_snmp_throttle Action

Redefine Snmp Throttle config

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_redefine_snmp_throttle_config_system_snmp_throttle" "example" {
  config {
    cluster_id = "example"
    throttle_config_details = [{
      interval         = 1
      notify_set       = "all"
      report_threshold = 0
      throttle_events  = ["secondflashboot"]
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

