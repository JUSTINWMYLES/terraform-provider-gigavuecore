---
page_title: "gigavuecore_load_cluster_event_notification_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Cluster Event Notification Status
---

# gigavuecore_load_cluster_event_notification_status Data Source

Load Cluster Event Notification Status

## Example Usage

```terraform
data "gigavuecore_load_cluster_event_notification_status" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster ID whose event notification status is queried

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_name` (String, computed) - Name of the cluster
* `event_notif_status_list` (Attributes List, computed) - Status of event notification for every node in the cluster (see [below for nested schema](#nestedatt--event_notif_status_list))

<a id="nestedatt--event_notif_status_list"></a>
### Nested Schema for `event_notif_status_list`

Read-Only:

* `device_ip` (String) - IP address of cluster member
* `enabled` (Boolean) - Should events be streamed from this device
* `event_targets_status` (Attributes List) - Status of every notification target the device is configured to stream to (see [below for nested schema](#nestedatt--event_notif_status_list--event_targets_status))
<a id="nestedatt--event_notif_status_list--event_targets_status"></a>
### Nested Schema for `event_notif_status_list.event_targets_status`

Read-Only:

* `encoding` (String) - Encoding format the notification target is configured with
* `error` (String) - Errors encountered when connecting to notification target
* `protocol` (String) - Messaging Protocol the notification target is configured with
* `secured` (Boolean) - Is communication channel secured for the notification target
* `state` (String) - Indicates if the device can reach this target
* `target_address` (String) - IP address of the notification target
* `target_id` (String) - ID of the notification target
* `target_port` (Number) - The port where notification target is listening
* `username` (String) - Username of the notification target

