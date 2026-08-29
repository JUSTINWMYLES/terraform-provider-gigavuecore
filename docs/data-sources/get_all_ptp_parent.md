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

* `alias` (String)
* `box_id` (String)
* `clock_identity` (String) - Parent Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array
* `clock_port_id` (Number) - Port Identity of the port on the source that issues the Sync message used in synchronizing this clock
* `grandmaster_clock_identity` (String) - grandmaster's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array. (deprecated: use primarySourceClockIdentity)
* `grandmaster_clock_quality` (Attributes) - Clock Quality determines the quality of a clock based on class and accuracy. Used to select master clock in Best Master Clock Algorithm. (deprecated: use primarySourceClockQuality instead of grandmasterClockQuality) (see [below for nested schema](#nestedatt--items--grandmaster_clock_quality))
* `grandmaster_priority1` (Number) - deprecated: use primarySourcePriority1
* `grandmaster_priority2` (Number) - deprecated: use primarySourcePriority1
* `observed_offset` (Number) - Estimate of the parent's clock PTP variance as observed by the receiver clock
* `observed_phase_change_rate` (Number) - Estimate of the parent's clock phase rate change as observed by the receiver clock
* `primary_source_clock_identity` (String) - primary source's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array
* `primary_source_clock_quality` (Attributes) - Clock Quality determines the quality of a clock based on class and accuracy. Used to select source clock in Best Source Clock Algorithm (see [below for nested schema](#nestedatt--items--primary_source_clock_quality))
* `primary_source_priority1` (Number)
* `primary_source_priority2` (Number)
* `stats` (Boolean) - indicates whether the values of observedOffset and observedPhaseChangeRate have been measured and are valid
<a id="nestedatt--items--grandmaster_clock_quality"></a>
### Nested Schema for `items.grandmaster_clock_quality`

Read-Only:

* `accuracy` (String) - Indicates the expected accuracy of a clock when it is the grandmaster or in the event it becomes the grandmaster
* `class` (Number) - Denotes the traceability of the time or frequency distributed by the grandmaster clock
* `offset` (Number) - Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol
<a id="nestedatt--items--primary_source_clock_quality"></a>
### Nested Schema for `items.primary_source_clock_quality`

Read-Only:

* `accuracy` (String) - Indicates the expected accuracy of a clock when it is the primary source or in the event it becomes the primary source
* `class` (Number) - Denotes the traceability of the time or frequency distributed by the primary source clock
* `offset` (Number) - Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol

