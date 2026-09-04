---
page_title: "gigavuecore_update_email_notif_config_spec Action - gigavuecore"
subcategory: ""
description: |-
  Update system email notification settings
---

# gigavuecore_update_email_notif_config_spec Action

Update system email notification settings

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_email_notif_config_spec" "example" {
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
* `notify_events` (Set of String, optional) - when 'notifySet' is 'select', this represents the subset of notification event types activated for dispatch
* `notify_set` (String, optional) - when set to 'none', effectively disables email notifications from the node


