---
page_title: "gigavuecore_get_email_notif_config_spec Data Source - gigavuecore"
subcategory: ""
description: |-
  Get system email notification settings
---

# gigavuecore_get_email_notif_config_spec Data Source

Get system email notification settings

## Example Usage

```terraform
data "gigavuecore_get_email_notif_config_spec" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `notify_events` (Set of String, computed) - when 'notifySet' is 'select', this represents the subset of notification event types activated for dispatch
* `notify_set` (String, computed) - when set to 'none', effectively disables email notifications from the node


