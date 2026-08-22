---
page_title: "gigavuecore_get_all_correlated_alarms Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Correlated Alarms
---

# gigavuecore_get_all_correlated_alarms Data Source

Load All Correlated Alarms

## Example Usage

```terraform
data "gigavuecore_get_all_correlated_alarms" "example" {
  acknowledged = null
  acknowledgedby = null
  alias = null
  cluster_id = null
  device_ip = null
  end_time = null
  hostname = null
  page = null
  resource_id = null
  resource_type = null
  severity = null
  sort = null
  start_time = null
  suppressed = null
  type = null
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
* `end_time` (String, optional) - End Time to filter by. In ISO 8601 format 'yyyy-MM-ddTHH:mm:ssZ'
* `hostname` (String, optional) - hostname to filter by
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:30) to prevent reading entire DB by accident
* `resource_id` (String, optional) - Affected entity to filter by
* `resource_type` (String, optional) - Affected entity type to filter by
* `severity` (String, optional) - Alarm severity to filter by
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is DESC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `start_time` (String, optional) - Start Time to filter by. In ISO 8601 format 'yyyy-MM-ddTHH:mm:ssZ'
* `suppressed` (String, optional) - filter by suppressed alarms
* `type` (String, optional) - Alarm Type to filter by
* `unacknowledgedby` (String, optional) - filter alarms that are unacknowledged by a specific user

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alarms` (List(Object({acknowledged, acknowledged_by, acknowledged_ts, alias, cluster_id, comment, description, device_ip, hostname, resource_id, resource_type, severity, suppressed, ts, type, unacknowledged_by, unacknowledged_ts})), computed)
* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type

