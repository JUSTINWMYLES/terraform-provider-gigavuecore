---
page_title: "gigavuecore_load_events Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Events
---

# gigavuecore_load_events Data Source

Load Events

## Example Usage

```terraform
data "gigavuecore_load_events" "example" {
  end_time      = null
  page          = null
  resource_id   = null
  resource_type = null
  scope         = null
  severity      = null
  sort          = null
  source        = null
  start_time    = null
  type          = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `end_time` (String, optional) - End Time to filter by. In ISO 8601 format
* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned. Defaults to (1:100) to prevent reading entire DB by accident
* `resource_id` (String, optional) - Affected Entity to filter by
* `resource_type` (String, optional) - Affected Entity Type to filter by
* `scope` (String, optional) - Event Scope to filter by
* `severity` (String, optional) - Event Scope to filter by
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)
* `source` (String, optional) - Event Source identifier to filter by
* `start_time` (String, optional) - Start Time to filter by. In ISO 8601 format
* `type` (String, optional) - Event Type to filter by

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `description` (String) - Event description
* `resource_id` (String) - Affected Entity of the event
* `resource_type` (String) - Affected Entity Type of the event
* `scope` (String) - Scope of the event
* `severity` (String) - Severity of the event
* `source` (String) - Event Source. Device ID for node-originated events like Traps or Syslogs. Component ID for FM-originated events
* `ts` (String) - Event timestamp in ISO 8601 format
* `ts_utc` (Number) - Timestamp in UTC milliseconds
* `type` (String) - Event Type identifier

