---
page_title: "gigavuecore_load_snmp_throttle_status Data Source - gigavuecore"
subcategory: ""
description: |-
  Load Snmp Throttle Status
---

# gigavuecore_load_snmp_throttle_status Data Source

Load Snmp Throttle Status

## Example Usage

```terraform
data "gigavuecore_load_snmp_throttle_status" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `snmp_throttle_status` (Attributes List, computed) (see [below for nested schema](#nestedatt--snmp_throttle_status))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--snmp_throttle_status"></a>
### Nested Schema for `snmp_throttle_status`

Read-Only:

* `elapsed_time` (Number) - Elapsed Time
* `entity_id` (String) - Entity ID
* `event` (Set of String) - The set of notification event types
* `interval` (Number) - Interval (in seconds)
* `last_time_triggered` (String) - Date and time of last time the event has been triggered
* `trigger_count` (Number) - Trigger count

