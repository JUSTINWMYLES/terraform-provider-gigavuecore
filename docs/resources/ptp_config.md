---
page_title: "gigavuecore_ptp_config Resource - gigavuecore"
subcategory: ""
description: |-
  Manages the ptp config resource.
---

# gigavuecore_ptp_config Resource

Manages the ptp config resource.

## Example Usage

```terraform
resource "gigavuecore_ptp_config" "example" {
  alias          = "example"
  box_id         = 0
  domain         = 0
  local_priority = 0
  mode           = "ordinary"
  priority2      = 0
  step_type      = "oneStep"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the time stamping ptp configuration
* `box_id` (Number, required) - device id
* `domain` (Number, optional) - Values supported for domain are 0 and range of values from 24 to 43
* `local_priority` (Number, optional)
* `mode` (String, optional)
* `priority2` (Number, optional)
* `step_type` (String, optional)

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_ptp_config.example {alias}
```
