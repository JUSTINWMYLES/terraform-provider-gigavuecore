---
page_title: "gigavuecore_get_ptp_parent_by_alias Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get ptp parent by alias data source.
---

# gigavuecore_get_ptp_parent_by_alias Data Source

Reads the get ptp parent by alias data source.

## Example Usage

```terraform
data "gigavuecore_get_ptp_parent_by_alias" "example" {
  alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the time stamping PTP configuration

### Attributes

In addition to all arguments above, the following attributes are exported:

* `box_id` (String, computed)
* `clock_identity` (String, computed) - Parent Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array
* `clock_port_id` (Number, computed) - Port Identity of the port on the source that issues the Sync message used in synchronizing this clock
* `grandmaster_clock_identity` (String, computed) - grandmaster's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array. (deprecated: use primarySourceClockIdentity)
* `grandmaster_clock_quality` (Attributes, computed) - Clock Quality determines the quality of a clock based on class and accuracy. Used to select master clock in Best Master Clock Algorithm. (deprecated: use primarySourceClockQuality instead of grandmasterClockQuality) (see [below for nested schema](#nestedatt--grandmaster_clock_quality))
* `grandmaster_priority1` (Number, computed) - deprecated: use primarySourcePriority1
* `grandmaster_priority2` (Number, computed) - deprecated: use primarySourcePriority1
* `observed_offset` (Number, computed) - Estimate of the parent's clock PTP variance as observed by the receiver clock
* `observed_phase_change_rate` (Number, computed) - Estimate of the parent's clock phase rate change as observed by the receiver clock
* `primary_source_clock_identity` (String, computed) - primary source's Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array
* `primary_source_clock_quality` (Attributes, computed) - Clock Quality determines the quality of a clock based on class and accuracy. Used to select source clock in Best Source Clock Algorithm (see [below for nested schema](#nestedatt--primary_source_clock_quality))
* `primary_source_priority1` (Number, computed)
* `primary_source_priority2` (Number, computed)
* `stats` (Boolean, computed) - indicates whether the values of observedOffset and observedPhaseChangeRate have been measured and are valid

<a id="nestedatt--grandmaster_clock_quality"></a>
### Nested Schema for `grandmaster_clock_quality`

Read-Only:

* `accuracy` (String) - Indicates the expected accuracy of a clock when it is the grandmaster or in the event it becomes the grandmaster
* `class` (Number) - Denotes the traceability of the time or frequency distributed by the grandmaster clock
* `offset` (Number) - Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol
<a id="nestedatt--primary_source_clock_quality"></a>
### Nested Schema for `primary_source_clock_quality`

Read-Only:

* `accuracy` (String) - Indicates the expected accuracy of a clock when it is the primary source or in the event it becomes the primary source
* `class` (Number) - Denotes the traceability of the time or frequency distributed by the primary source clock
* `offset` (Number) - Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol

