---
page_title: "gigavuecore_get_all_foreign_sources Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get all foreign sources data source.
---

# gigavuecore_get_all_foreign_sources Data Source

Reads the get all foreign sources data source.

## Example Usage

```terraform
data "gigavuecore_get_all_foreign_sources" "example" {
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

* `best_source` (Number)
* `foreign_sources` (Attributes List) (see [below for nested schema](#nestedatt--items--foreign_sources))
* `port_id` (String)
<a id="nestedatt--items--foreign_sources"></a>
### Nested Schema for `items.foreign_sources`

Read-Only:

* `announce_msg_count` (Number) - Number of Announce messages from the foreign source that have been received within a time window
* `clock_address` (String)
* `last_delay_response` (String)
* `last_follow_up` (String)
* `last_sync` (String)
* `mtsd_scaled_avar` (Number)
* `port_identity` (String)
* `port_module_number` (Number)
* `port_number` (Number)
* `primary_source_clock_identity` (String) - primary source's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array
* `primary_source_priority1` (Number)
* `primary_source_priority2` (Number)
* `ptp_protocol` (String)
* `quality` (Attributes) - Clock Quality determines the quality of a clock based on class and accuracy. Used to select source clock in Best Source Clock Algorithm (see [below for nested schema](#nestedatt--items--foreign_sources--quality))
* `steps_removed` (Number)
<a id="nestedatt--items--foreign_sources--quality"></a>
### Nested Schema for `items.foreign_sources.quality`

Read-Only:

* `accuracy` (String) - Indicates the expected accuracy of a clock when it is the primary source or in the event it becomes the primary source
* `class` (Number) - Denotes the traceability of the time or frequency distributed by the primary source clock
* `offset` (Number) - Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol

