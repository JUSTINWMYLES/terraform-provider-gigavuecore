---
page_title: "gigavuecore_get_all_ptp_parent Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all ptp parent data source.
---

# gigavuecore_get_all_ptp_parent Data Source

Reads the get all ptp parent data source.

## Example Usage

```terraform
data "gigavuecore_get_all_ptp_parent" "example" {
  box_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (Number, optional) - specify the cluster node by boxId. By default all nodes are selected.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, box_id, clock_identity, clock_port_id, grandmaster_clock_identity, grandmaster_clock_quality, grandmaster_priority1, grandmaster_priority2, observed_offset, observed_phase_change_rate, primary_source_clock_identity, primary_source_clock_quality, primary_source_priority1, primary_source_priority2, stats})), computed)

