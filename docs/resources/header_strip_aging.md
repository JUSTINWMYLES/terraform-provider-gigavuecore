---
page_title: "gigavuecore_header_strip_aging Resource - gigavuecore"
subcategory: ""
description: |-
  Load header strip aging for target box
---

# gigavuecore_header_strip_aging Resource

Load header strip aging for target box

## Example Usage

```terraform
resource "gigavuecore_header_strip_aging" "example" {
  aging_interval = null
  box_id         = null
  dst_port       = null
  protocol_type  = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `aging_interval` (Number, optional) - Interval in sec. Valid range 300-1000000. Enter 0 to disable.
* `box_id` (String, required) - device box id. valid range 1 - 64. all is applicable only for post request.
* `dst_port` (Number, optional) - L4 destination port number.Valid value is between 0 to 65535.
* `protocol_type` (String, optional) - protocol type

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `aging_interval` (Number, computed) - Interval in sec. Valid range 300-1000000. Enter 0 to disable.
* `dst_port` (Number, computed) - L4 destination port number.Valid value is between 0 to 65535.
* `id` (String, computed)
* `protocol_type` (String, computed) - protocol type


