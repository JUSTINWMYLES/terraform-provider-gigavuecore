---
page_title: "gigavuecore_get_ptp_clock_state_query_response Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get ptp clock state query response data source.
---

# gigavuecore_get_ptp_clock_state_query_response Data Source

Reads the get ptp clock state query response data source.

## Example Usage

```terraform
data "gigavuecore_get_ptp_clock_state_query_response" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the time stamping PTP configuration

### Attributes

In addition to all arguments above, the following attributes are exported:

* `box_id` (String, computed) - device id
* `clock_identity` (String, computed) - Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array
* `clock_quality` (Attributes, computed) - Clock Quality determines the quality of a clock based on class and accuracy. Used to select source clock in Best Source Clock Algorithm (see [below for nested schema](#nestedatt--clock_quality))
* `domain` (Number, computed)
* `local_clock_time` (String, computed) - \[RFC 3339\](https://tools.ietf.org/html/rfc3339) format
* `local_priority` (Number, computed)
* `mean_path_delay` (String, computed) - Represents current value of the mean propagation time between source and receiver clock as computed by the slave. format hh:mm:ss.secfrac
* `mode` (String, computed)
* `offset_from_master` (String, computed) - Represents current value of time difference between master and a slave as computed by the slave. format hh:mm:ss. (deprecated: use offsetFromSource)
* `offset_from_source` (String, computed) - Represents current value of time difference between source and a receiver as computed by the receiver. format hh:mm:ss
* `port_ptp_count` (Number, computed) - Total number of PTP ports per node
* `priority2` (Number, computed)
* `ptp_time` (String, computed)
* `step_type` (String, computed)
* `steps_removed` (Number, computed) - Represents the number of communication paths traversed between the local clock and the primary source clock

<a id="nestedatt--clock_quality"></a>
### Nested Schema for `clock_quality`

Read-Only:

* `accuracy` (String) - Indicates the expected accuracy of a clock when it is the primary source or in the event it becomes the primary source
* `class` (Number) - Denotes the traceability of the time or frequency distributed by the primary source clock
* `offset` (Number) - Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol

