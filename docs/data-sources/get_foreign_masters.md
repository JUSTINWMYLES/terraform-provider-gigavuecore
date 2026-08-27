---
page_title: "gigavuecore_get_foreign_masters Data Source - gigavuecore"
subcategory: ""
description: |-
  deprecated: use GET /ptp/portState/{portId}/foreignSource
---

# gigavuecore_get_foreign_masters Data Source

deprecated: use GET /ptp/portState/{portId}/foreignSource

## Example Usage

```terraform
data "gigavuecore_get_foreign_masters" "example" {
  port_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `port_id` (String, required) - id of the target device Port (format: boxId\_slotId\_port, example: 1\_1\_c1)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `best_master` (Number, computed)
* `foreign_masters` (Attributes List, computed) (see [below for nested schema](#nestedatt--foreign_masters))

<a id="nestedatt--foreign_masters"></a>
### Nested Schema for `foreign_masters`

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
* `quality` (Attributes) - Clock Quality determines the quality of a clock based on class and accuracy. Used to select master clock in Best Master Clock Algorithm. (deprecated: use primarySourceClockQuality instead of grandmasterClockQuality) (see [below for nested schema](#nestedatt--foreign_masters--quality))
* `steps_removed` (Number)
<a id="nestedatt--foreign_masters--quality"></a>
### Nested Schema for `foreign_masters.quality`

Read-Only:

* `accuracy` (String) - Indicates the expected accuracy of a clock when it is the grandmaster or in the event it becomes the grandmaster
* `class` (Number) - Denotes the traceability of the time or frequency distributed by the grandmaster clock
* `offset` (Number) - Indicates the estimate of the variations of the local clock from a linear timescale when it is not synchronized to another clock using the protocol

