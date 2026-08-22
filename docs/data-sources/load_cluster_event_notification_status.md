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
* `event_notif_status_list` (List(Object({device_ip, enabled, event_targets_status})), computed) - Status of event notification for every node in the cluster

