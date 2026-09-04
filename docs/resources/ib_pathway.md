---
page_title: "gigavuecore_ib_pathway Resource - gigavuecore"
subcategory: ""
description: |-
  Create a new Inter-broker Pathway
---

# gigavuecore_ib_pathway Resource

Create a new Inter-broker Pathway

## Example Usage

```terraform
resource "gigavuecore_ib_pathway" "example" {
  alias        = "example"
  comment      = "example"
  min_ports_up = 1
  ports        = ["example"]
  traffic_path = "bypass"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Inline resilient inter-broker pathway alias
* `comment` (String, optional)
* `min_ports_up` (Number, optional) - minimum number of ports in the 'up' state needed to declare the ib-pathway to be in the 'up' state.
* `ports` (List of String, optional) - list of local network ports with same speed
* `traffic_path` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `operational_state` (String, computed)

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
terraform import gigavuecore_ib_pathway.example {alias}
```
