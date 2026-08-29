---
page_title: "gigavuecore_ib_pathway Resource - gigavuecore"
subcategory: ""
description: |-
  Load Inter-broker Pathway by alias
---

# gigavuecore_ib_pathway Resource

Load Inter-broker Pathway by alias

## Example Usage

```terraform
resource "gigavuecore_ib_pathway" "example" {
  alias        = "example"
  comment      = "example"
  min_ports_up = 0
  ports        = [ "example" ]
  traffic_path = "example"
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

* `comment` (String, computed)
* `min_ports_up` (Number, computed) - minimum number of ports in the 'up' state needed to declare the ib-pathway to be in the 'up' state.
* `operational_state` (String, computed)
* `ports` (List of String, computed) - list of local network ports with same speed
* `traffic_path` (String, computed)


## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_ib_pathway.example {alias}
```
