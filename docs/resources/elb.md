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
  alias       = null
  hash_fields = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - app elb alias
* `hash_fields` (Attributes List, optional) (see [below for nested schema](#nestedatt--hash_fields))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `hash_fields` (Attributes List, computed) (see [below for nested schema](#nestedatt--hash_fields))

<a id="nestedatt--hash_fields"></a>
### Nested Schema for `hash_fields`

Required:

* `hash_field` (String)
Optional:

* `hash_location` (String) - Ignored when 'hashField' == 'gtpuTeid'. required otherwise

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_elb.example {alias}
```
