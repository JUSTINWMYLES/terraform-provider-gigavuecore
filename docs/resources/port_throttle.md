---
page_title: "gigavuecore_port_throttle Resource - gigavuecore"
subcategory: ""
description: |-
  Create a Port Throttle
---

# gigavuecore_port_throttle Resource

Create a Port Throttle

## Example Usage

```terraform
resource "gigavuecore_port_throttle" "example" {
  alias = "example"
  ports_throttles = [{
    port  = "example"
    type  = "pps"
    value = 33
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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_port_throttle.example {alias}
```
