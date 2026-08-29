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
  aging_interval = 0
  box_id         = "example"
  dst_port       = 0
  protocol_type  = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `aging_interval` (Number, optional) - Interval in sec. Valid range 300-1000000. Enter 0 to disable.
* `box_id` (String, required) - device box id. valid range 1 - 64. all is applicable only for post request.
* `dst_port` (Number, optional) - L4 destination port number.Valid value is between 0 to 65535.
* `protocol_type` (String, optional) - protocol type

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

