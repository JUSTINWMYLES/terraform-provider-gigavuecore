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

* `ptp_time` (String)
* `alias` (String)
* `box_id` (String) - device id
* `clock_identity` (String) - Clock Identifier as per the IEEE 1588 standard which is represented by an integer 8-octet array
* `clock_quality` (Attributes) - Clock Quality determines the quality of a clock based on class and accuracy. Used to select source clock in Best Source Clock Algorithm (see [below for nested schema](#nestedatt--items--clock_quality))
* `domain` (Number)
* `local_clock_time` (String) - \[RFC 3339\](https://tools.ietf.org/html/rfc3339) format
* `local_priority` (Number)
* `mean_path_delay` (String) - Represents current value of the mean propagation time between source and receiver clock as computed by the slave. format hh:mm:ss.secfrac
* `mode` (String)
* `offset_from_master` (String) - Represents current value of time difference between master and a slave as computed by the slave. format hh:mm:ss. (deprecated: use offsetFromSource)
* `offset_from_source` (String) - Represents current value of time difference between source and a receiver as computed by the receiver. format hh:mm:ss
* `port_ptp_count` (Number) - Total number of PTP ports per node
* `priority2` (Number)
* `step_type` (String)
* `steps_removed` (Number) - Represents the number of communication paths traversed between the local clock and the primary source clock

<a id="nestedatt--items--clock_quality"></a>
### Nested Schema for `items.clock_quality`

Read-Only:

* `accuracy` (String) - Indicates the expected accuracy of a clock when it is the primary source or in the event it becomes the primary source
* `class` (Number) - Denotes the traceability of the time or frequency distributed by the primary source clock
* `offset` (Number) - Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol

