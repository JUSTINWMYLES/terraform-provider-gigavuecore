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
  alias = "example"
  ports_throttles = [{
    port  = "example"
    type  = "example"
    value = 0
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of port throttle
* `ports_throttles` (Attributes List, required) (see [below for nested schema](#nestedatt--ports_throttles))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--ports_throttles"></a>
### Nested Schema for `ports_throttles`

Required:

* `port` (String) - ports or gigastreams
* `type` (String)
* `value` (Number) - Port throttle value
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - A update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_throttle.example {alias}
```
