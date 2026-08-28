---
page_title: "gigavuecore_load_all_alarms List Resource - gigavuecore"
subcategory: ""
description: |-
  Load All Alarms
---

# gigavuecore_load_all_alarms List Resource

Load All Alarms

## Example Usage

```terraform
list "gigavuecore_load_all_alarms" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    acknowledged     = "example"
    acknowledgedby   = "example"
    alias            = "example"
    cluster_id       = "example"
    device_ip        = "example"
    end_time         = "example"
    hostname         = "example"
    page             = "example"
    resource_id      = "example"
    resource_type    = "example"
    severity         = "example"
    sort             = "example"
    start_time       = "example"
    suppressed       = "example"
    type             = "example"
    unacknowledgedby = "example"
  }
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


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alarm_id` (String, computed)


