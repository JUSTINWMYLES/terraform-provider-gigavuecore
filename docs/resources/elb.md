---
page_title: "gigavuecore_elb Resource - gigavuecore"
subcategory: ""
description: |-
  Load a Elb
---

# gigavuecore_elb Resource

Load a Elb

## Example Usage

```terraform
resource "gigavuecore_elb" "example" {
  alias = null
  hash_fields = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - app elb alias
* `hash_fields` (List(Object({hash_field, hash_location})), optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `hash_fields` (List(Object({hash_field, hash_location})), computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_elb.example {alias}
```
