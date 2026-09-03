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
  category        = "LINK"
  cluster_id      = "example"
  device_ip       = "example"
  end_time        = "example"
  hostname        = "example"
  page            = "example"
  process         = "example"
  resource_id     = "example"
  resource_type   = "Cpu"
  severity        = "EMERG"
  sort            = "example"
  start_time      = "example"
  syslog_priority = "example"
  syslog_version  = "example"
  type            = "example"
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

* `customer_id` (String, computed) - Customer Id of the log. Gigamon reserved customerId is 26866
* `description` (String, computed) - Log description
* `pid` (String, computed) - Process Id of the process generating the log
* `source` (String, computed) - Log Source. Control Card or GigaSmart
* `structured_data` (String, computed) - Additional information about the log. Example: link speed=10000, portID=1/1/x1 for linkChangeNotify event
* `timestamp_utc_string` (String, computed) - Log timestamp in UTC represented in ISO 8601 format


