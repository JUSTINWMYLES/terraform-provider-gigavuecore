---
page_title: "gigavuecore_update_email_notif_config_spec Action - gigavuecore"
subcategory: ""
description: |-
  Update system email notification settings
---

# gigavuecore_update_email_notif_config_spec Action

Update system email notification settings

## Example Usage

```terraform
action "gigavuecore_update_email_notif_config_spec" "example" {
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
* `notify_events` (Set of String, optional) - when 'notifySet' is 'select', this represents the subset of notification event types activated for dispatch
* `notify_set` (String, optional) - when set to 'none', effectively disables email notifications from the node


