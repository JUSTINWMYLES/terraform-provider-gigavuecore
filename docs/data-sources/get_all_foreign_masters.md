---
page_title: "gigavuecore_get_all_foreign_masters Data Source - gigavuecore"
subcategory: ""
description: |-
  deprecated: use GET /ptp/portState/foreignSource
---

# gigavuecore_get_all_foreign_masters Data Source

deprecated: use GET /ptp/portState/foreignSource

## Example Usage

```terraform
data "gigavuecore_get_all_foreign_masters" "example" {
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

* `best_master` (Number)
* `foreign_masters` (Attributes List) (see [below for nested schema](#nestedatt--items--foreign_masters))
* `port_id` (String)
<a id="nestedatt--items--foreign_masters"></a>
### Nested Schema for `items.foreign_masters`

Read-Only:

* `announce_msg_count` (Number) - Number of Announce messages from the foreign master that have been received within a time window
* `clock_address` (String)
* `grandmaster_clock_identity` (String) - grandmaster's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array
* `grandmaster_priority1` (Number)
* `grandmaster_priority2` (Number)
* `last_delay_response` (String)
* `last_follow_up` (String)
* `last_sync` (String)
* `mtsd_scaled_avar` (Number)
* `port_identity` (String)
* `port_module_number` (Number)
* `port_number` (Number)
* `ptp_protocol` (String)
* `quality` (Attributes) - Clock Quality determines the quality of a clock based on class and accuracy. Used to select master clock in Best Master Clock Algorithm. (deprecated: use primarySourceClockQuality instead of grandmasterClockQuality) (see [below for nested schema](#nestedatt--items--foreign_masters--quality))
* `steps_removed` (Number)
<a id="nestedatt--items--foreign_masters--quality"></a>
### Nested Schema for `items.foreign_masters.quality`

Read-Only:

* `accuracy` (String) - Indicates the expected accuracy of a clock when it is the grandmaster or in the event it becomes the grandmaster
* `class` (Number) - Denotes the traceability of the time or frequency distributed by the grandmaster clock
* `offset` (Number) - Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol

