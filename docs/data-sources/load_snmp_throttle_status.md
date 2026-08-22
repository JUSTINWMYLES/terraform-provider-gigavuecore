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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Object({page_no, page_size, sort, total_items}), computed) - Gigamon query result context
  * `page_no` (Number, computed) - page number of the returned result set
  * `page_size` (Number, computed) - page size of the returned result set
  * `sort` (List(String), computed) - sorting info of the returned result set. list of fields in the array indicate sorting order
  * `total_items` (Number, computed) - total number of items in the queried entity type
* `snmp_throttle_status` (List(Object({elapsed_time, entity_id, event, interval, last_time_triggered, trigger_count})), computed)

