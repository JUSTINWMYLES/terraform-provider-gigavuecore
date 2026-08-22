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
* `foreign_masters` (List(Object({announce_msg_count, clock_address, grandmaster_clock_identity, grandmaster_priority1, grandmaster_priority2, last_delay_response, last_follow_up, last_sync, mtsd_scaled_avar, port_identity, port_module_number, port_number, ptp_protocol, quality, steps_removed})), computed)

