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
  alias          = null
  box_id         = null
  domain         = null
  local_priority = null
  mode           = null
  priority2      = null
  step_type      = null
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

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `domain` (Number, computed) - Values supported for domain are 0 and range of values from 24 to 43
* `local_priority` (Number, computed)
* `mode` (String, computed)
* `priority2` (Number, computed)
* `step_type` (String, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_ptp_config.example {alias}
```
