---
page_title: "gigavuecore_get_event_by_id Data Source - gigavuecore"
subcategory: ""
description: |-
  Find Event by ID
---

# gigavuecore_get_event_by_id Data Source

Find Event by ID

## Example Usage

```terraform
data "gigavuecore_get_event_by_id" "example" {
  event_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `event_id` (String, required) - ID of the target Event

### Attributes

In addition to all arguments above, the following attributes are exported:

* `description` (String, computed) - Event description
* `resource_id` (String, computed) - Affected Entity of the event
* `resource_type` (String, computed) - Affected Entity Type of the event
* `scope` (String, computed) - Scope of the event
* `severity` (String, computed) - Severity of the event
* `source` (String, computed) - Event Source. Device ID for node-originated events like Traps or Syslogs. Component ID for FM-originated events
* `ts` (String, computed) - Event timestamp in ISO 8601 format
* `ts_utc` (Number, computed) - Timestamp in UTC milliseconds
* `type` (String, computed) - Event Type identifier


