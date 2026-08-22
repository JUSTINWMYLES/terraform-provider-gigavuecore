---
page_title: "gigavuecore_get_alarm_by_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Find Alarm by ID
---

# gigavuecore_get_alarm_by_id Data Source

Find Alarm by ID

## Example Usage

```terraform
data "gigavuecore_get_alarm_by_id" "example" {
  alarm_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alarm_id` (String, required) - ID of the target Alarm

### Attributes

In addition to all arguments above, the following attributes are exported:

* `acknowledged` (Bool, computed) - True if the alarm is acknowledged by the user
* `acknowledged_by` (String, computed) - User who acknowledged the alarm
* `acknowledged_ts` (String, computed) - Alarm acknowledged timestamp in ISO 8601 format
* `alias` (String, computed) - Alias of the Resource ID
* `cluster_id` (String, computed) - Cluster ID of the device
* `comment` (String, computed) - User comments
* `description` (String, computed) - Alarm description
* `device_ip` (String, computed) - IP address of the device
* `hostname` (String, computed) - Hostname of the Resource
* `resource_id` (String, computed) - Resource ID of the alarm
* `resource_type` (String, computed) - Resource Type of the alarm
* `severity` (String, computed) - Severity of the alarm
* `suppressed` (Bool, computed) - True if the alarm is suppressed
* `ts` (String, computed) - Alarm timestamp in ISO 8601 format
* `type` (String, computed) - Alarm Type identifier
* `unacknowledged_by` (String, computed) - User who unacknowledged the alarm
* `unacknowledged_ts` (String, computed) - Alarm unacknowledged timestamp in ISO 8601 format

