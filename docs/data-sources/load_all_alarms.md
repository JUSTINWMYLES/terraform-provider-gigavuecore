---
page_title: "gigavuecore_load_all_alarms Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Alarms
---

# gigavuecore_load_all_alarms Data Source

Load All Alarms

## Example Usage

```terraform
data "gigavuecore_load_all_alarms" "example" {
  acknowledged     = null
  acknowledgedby   = null
  alias            = null
  cluster_id       = null
  device_ip        = null
  end_time         = null
  hostname         = null
  page             = null
  resource_id      = null
  resource_type    = null
  severity         = null
  sort             = null
  start_time       = null
  suppressed       = null
  type             = null
  unacknowledgedby = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `acknowledged` (String, optional) - filter by acknowledged alarms
* `acknowledgedby` (String, optional) - filter alarms that are acknowledged by a specific user
* `alias` (String, optional) - resource alias name to filter by
* `cluster_id` (String, optional) - Cluster ID to filter by
* `device_ip` (String, optional) - IP address of Device to filter by
* `end_time` (String, optional) - Filter End timestamps to include reports ending by. In ISO-8601 date format 'yyyy-MM-ddTHH:mm:ssZ'
* `hostname` (String, optional) - hostname to filter by
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident
* `resource_id` (String, optional) - Affected entity to filter by
* `resource_type` (String, optional) - Affected entity type to filter by
* `severity` (String, optional) - Alarm severity to filter by
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `start_time` (String, optional) - Filter Start timestamps to include reports starting from. In ISO-8601 date format 'yyyy-MM-ddTHH:mm:ssZ'
* `suppressed` (String, optional) - filter by suppressed alarms
* `type` (String, optional) - Alarm Type to filter by
* `unacknowledgedby` (String, optional) - filter alarms that are unacknowledged by a specific user

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `acknowledged` (Boolean) - True if the alarm is acknowledged by the user
* `acknowledged_by` (String) - User who acknowledged the alarm
* `acknowledged_ts` (String) - Alarm acknowledged timestamp in ISO 8601 format
* `alias` (String) - Alias of the Resource ID
* `cluster_id` (String) - Cluster ID of the device
* `comment` (String) - User comments
* `description` (String) - Alarm description
* `device_ip` (String) - IP address of the device
* `hostname` (String) - Hostname of the Resource
* `resource_id` (String) - Resource ID of the alarm
* `resource_type` (String) - Resource Type of the alarm
* `severity` (String) - Severity of the alarm
* `suppressed` (Boolean) - True if the alarm is suppressed
* `ts` (String) - Alarm timestamp in ISO 8601 format
* `type` (String) - Alarm Type identifier
* `unacknowledged_by` (String) - User who unacknowledged the alarm
* `unacknowledged_ts` (String) - Alarm unacknowledged timestamp in ISO 8601 format

