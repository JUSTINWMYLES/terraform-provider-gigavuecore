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
  box_id = 0
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `ptp_time_scale` (Number)
* `alias` (String)
* `box_id` (Number)
* `current_utc_offset` (Number)
* `frequency_traceable` (Number)
* `leap59` (Number)
* `leap61` (Number)
* `time_source` (String)
* `time_traceable` (Number)

