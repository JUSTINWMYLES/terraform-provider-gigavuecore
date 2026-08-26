---
page_title: "gigavuecore_load_device_syslog Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Device Syslog
---

# gigavuecore_load_device_syslog Data Source

Load Device Syslog

## Example Usage

```terraform
data "gigavuecore_load_device_syslog" "example" {
  category        = null
  cluster_id      = null
  device_ip       = null
  end_time        = null
  hostname        = null
  page            = null
  process         = null
  resource_id     = null
  resource_type   = null
  severity        = null
  sort            = null
  start_time      = null
  syslog_priority = null
  syslog_version  = null
  type            = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `category` (String, optional) - Category of the log
* `cluster_id` (String, optional) - Cluster ID of the device
* `device_ip` (String, optional) - Device IP address
* `end_time` (String, optional) - End time to filter by. In UTC, ISO 8601 format
* `hostname` (String, optional) - Device host name
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident
* `process` (String, optional) - System process generating the log
* `resource_id` (String, optional) - Affected Entity of the log
* `resource_type` (String, optional) - Affected Entity Type of the log
* `severity` (String, optional) - Severity of the event
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `start_time` (String, optional) - Start time to filter by. In UTC, ISO 8601 format
* `syslog_priority` (String, optional) - Syslog Priority as per RFC 5424
* `syslog_version` (String, optional) - Syslog Version as per RFC 5424
* `type` (String, optional) - Log Type identifier. Example linkChangeNotify

### Attributes

In addition to all arguments above, the following attributes are exported:

* `category` (String, computed) - Category of the log
* `customer_id` (String, computed) - Customer Id of the log. Gigamon reserved customerId is 26866
* `description` (String, computed) - Log description
* `device_ip` (String, computed) - Device IP address
* `hostname` (String, computed) - Device host name
* `pid` (String, computed) - Process Id of the process generating the log
* `process` (String, computed) - System process generating the log
* `resource_id` (String, computed) - Affected Entity of the log
* `resource_type` (String, computed) - Affected Entity Type of the log
* `severity` (String, computed) - Severity of the event
* `source` (String, computed) - Log Source. Control Card or GigaSmart
* `structured_data` (String, computed) - Additional information about the log. Example: link speed=10000, portID=1/1/x1 for linkChangeNotify event
* `syslog_priority` (String, computed) - Syslog Priority as per RFC 5424
* `syslog_version` (String, computed) - Syslog Version as per RFC 5424
* `timestamp_utc_string` (String, computed) - Log timestamp in UTC represented in ISO 8601 format
* `type` (String, computed) - Log Type identifier. Example linkChangeNotify


