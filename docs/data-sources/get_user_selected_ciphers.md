---
page_title: "gigavuecore_get_user_selected_ciphers Data Source - gigavuecore"
subcategory: ""
description: |-
  Get user-selected ciphers by protocol type
---

# gigavuecore_get_user_selected_ciphers Data Source

Get user-selected ciphers by protocol type

## Example Usage

```terraform
data "gigavuecore_get_user_selected_ciphers" "example" {
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `type` (String, required) - Cipher protocol type (for example, TLS, SSH)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `value` (Dynamic, computed)

