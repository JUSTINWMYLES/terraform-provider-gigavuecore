---
page_title: "gigavuecore_get_foreign_sources Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get foreign sources data source.
---

# gigavuecore_get_foreign_sources Data Source

Reads the get foreign sources data source.

## Example Usage

```terraform
data "gigavuecore_get_foreign_sources" "example" {
  port_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `port_id` (String, required) - id of the target device Port (format: boxId\_slotId\_port, example: 1\_1\_c1)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `best_source` (Number, computed)
* `foreign_sources` (List(Object({announce_msg_count, clock_address, last_delay_response, last_follow_up, last_sync, mtsd_scaled_avar, port_identity, port_module_number, port_number, primary_source_clock_identity, primary_source_priority1, primary_source_priority2, ptp_protocol, quality, steps_removed})), computed)

