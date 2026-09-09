---
page_title: "gigavuecore_elb Resource - gigavuecore"
subcategory: ""
description: |-
  Create a Elb
---

# gigavuecore_elb Resource

Create a Elb

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

* `create` (Number) - A create timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `read` (Number) - A read timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `update` (Number) - An update timeout in seconds for this operation. Overrides the generator default (1200 seconds).
* `delete` (Number) - A delete timeout in seconds for this operation. Overrides the generator default (1200 seconds).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_elb.example {alias}
```
