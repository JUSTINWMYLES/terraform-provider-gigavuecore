---
page_title: "gigavuecore_get_supported_ciphers Data Source - gigavuecore"
subcategory: ""
description: |-
  Get supported ciphers (optionally filtered by protocol type)
---

# gigavuecore_get_supported_ciphers Data Source

Get supported ciphers (optionally filtered by protocol type)

## Example Usage

```terraform
data "gigavuecore_get_supported_ciphers" "example" {
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `type` (String, optional) - Cipher protocol type (optional filter). Allowed values -> TLS, SSH.

### Attributes

In addition to all arguments above, the following attributes are exported:

* `value` (Dynamic, computed)


