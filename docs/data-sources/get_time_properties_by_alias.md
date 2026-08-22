---
page_title: "gigavuecore_get_time_properties_by_alias Data Source - gigavuecore"
subcategory: ""
description: |-
  Reads the get time properties by alias data source.
---

# gigavuecore_get_time_properties_by_alias Data Source

Reads the get time properties by alias data source.

## Example Usage

```terraform
data "gigavuecore_get_time_properties_by_alias" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the time stamping PTP configuration

### Attributes

In addition to all arguments above, the following attributes are exported:

* `box_id` (Number, computed)
* `current_utc_offset` (Number, computed)
* `frequency_traceable` (Number, computed)
* `leap59` (Number, computed)
* `leap61` (Number, computed)
* `ptp_time_scale` (Number, computed)
* `time_source` (String, computed)
* `time_traceable` (Number, computed)

