---
page_title: "gigavuecore_get_all_time_stamping_ptp_configs Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all time stamping ptp configs data source.
---

# gigavuecore_get_all_time_stamping_ptp_configs Data Source

Reads the get all time stamping ptp configs data source.

## Example Usage

```terraform
data "gigavuecore_get_all_time_stamping_ptp_configs" "example" {
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

* `alias` (String) - alias of the time stamping ptp configuration
* `box_id` (Number) - device id
* `domain` (Number) - Values supported for domain are 0 and range of values from 24 to 43
* `local_priority` (Number)
* `mode` (String)
* `priority2` (Number)
* `step_type` (String)

