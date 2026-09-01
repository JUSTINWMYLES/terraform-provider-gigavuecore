---
page_title: "gigavuecore_update_snmp_server_nofify_config Action - gigavuecore"
subcategory: ""
description: |-
  Update Snmp Server Notification settings
---

# gigavuecore_update_snmp_server_nofify_config Action

Update Snmp Server Notification settings

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_snmp_server_nofify_config" "example" {
  config {
    cluster_id    = "example"
    notify_events = ["secondflashboot"]
    notify_set    = "all"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `notify_events` (Set of String, optional) - The set of notification event types
* `notify_set` (String, optional) - when 'notifySet' is 'select', notifyEvents represents the subset of notification event types activated for dispatch. When set to 'none', effectively disables SNMP notifications from the node


