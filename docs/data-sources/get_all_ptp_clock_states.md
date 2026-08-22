---
page_title: "gigavuecore_get_all_ptp_clock_states Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all ptp clock states data source.
---

# gigavuecore_get_all_ptp_clock_states Data Source

Reads the get all ptp clock states data source.

## Example Usage

```terraform
data "gigavuecore_get_all_ptp_clock_states" "example" {
  box_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({ptp_time, alias, box_id, clock_identity, clock_quality, domain, local_clock_time, local_priority, mean_path_delay, mode, offset_from_master, offset_from_source, port_ptp_count, priority2, step_type, steps_removed})), computed)

