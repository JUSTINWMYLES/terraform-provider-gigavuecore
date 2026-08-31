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
  alias = "example"
  hash_fields = [{
    hash_field    = "ip"
    hash_location = "inner"
  }]
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - app elb alias
* `hash_fields` (Attributes List, optional) (see [below for nested schema](#nestedatt--hash_fields))

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--hash_fields"></a>
### Nested Schema for `hash_fields`

Required:

* `hash_field` (String)

Optional:

* `hash_location` (String) - Ignored when 'hashField' == 'gtpuTeid'. required otherwise
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
terraform import gigavuecore_elb.example {alias}
```
