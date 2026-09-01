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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `elapsed_time` (Number) - Elapsed Time
* `entity_id` (String) - Entity ID
* `event` (Set of String) - The set of notification event types
* `interval` (Number) - Interval (in seconds)
* `last_time_triggered` (String) - Date and time of last time the event has been triggered
* `trigger_count` (Number) - Trigger count

