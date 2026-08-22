---
page_title: "gigavuecore_enhanced_slicing Resource - gigavuecore"
subcategory: ""
description: |-
  new in version H 5.7
---

# gigavuecore_enhanced_slicing Resource

new in version H 5.7

## Example Usage

```terraform
resource "gigavuecore_enhanced_slicing" "example" {
  alias = null
  hash_field = null
  max_sessions = null
  protocol_fields = []
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - apps enhanced slicing alias
* `hash_field` (String, optional) - Hash field for session
* `max_sessions` (Number, optional) - Maximum number of session entries (in millions). Used only for flow-session option
* `protocol_fields` (List(Object({gtp, gtp_u, ip, none, transport})), required)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `hash_field` (String, computed) - Hash field for session
* `max_sessions` (Number, computed) - Maximum number of session entries (in millions). Used only for flow-session option

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_enhanced_slicing.example {alias}
```
