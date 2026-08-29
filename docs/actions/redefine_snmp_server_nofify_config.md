---
page_title: "gigavuecore_redefine_snmp_server_nofify_config Action - gigavuecore"
subcategory: ""
description: |-
  Redefine Snmp Server Notification settings
---

# gigavuecore_redefine_snmp_server_nofify_config Action

Redefine Snmp Server Notification settings

## Example Usage

```terraform
action "gigavuecore_redefine_snmp_server_nofify_config" "example" {
  config {
    cluster_id    = "example"
    notify_events = [ "example" ]
    notify_set    = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `notify_events` (Set of String, optional) - The set of notification event types
* `notify_set` (String, optional) - when 'notifySet' is 'select', notifyEvents represents the subset of notification event types activated for dispatch. When set to 'none', effectively disables SNMP notifications from the node


