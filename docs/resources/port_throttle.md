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
  alias           = null
  ports_throttles = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of port throttle
* `ports_throttles` (Attributes List, required) (see [below for nested schema](#nestedatt--ports_throttles))

<a id="nestedatt--ports_throttles"></a>
### Nested Schema for `ports_throttles`

Required:

* `port` (String) - ports or gigastreams
* `type` (String)
* `value` (Number) - Port throttle value

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_throttle.example {alias}
```
