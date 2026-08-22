---
page_title: "gigavuecore_port_throttle Resource - gigavuecore"
subcategory: ""
description: |-
  Load Port Throttle by alias
---

# gigavuecore_port_throttle Resource

Load Port Throttle by alias

## Example Usage

```terraform
resource "gigavuecore_port_throttle" "example" {
  alias = null
  ports_throttles = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of port throttle
* `ports_throttles` (List(Object({port, type, value})), required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_throttle.example {alias}
```
