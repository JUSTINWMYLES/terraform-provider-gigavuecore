---
page_title: "gigavuecore_get_all_time_properties Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all time properties data source.
---

# gigavuecore_get_all_time_properties Data Source

Reads the get all time properties data source.

## Example Usage

```terraform
data "gigavuecore_get_all_time_properties" "example" {
  box_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({ptp_time_scale, alias, box_id, current_utc_offset, frequency_traceable, leap59, leap61, time_source, time_traceable})), computed)

