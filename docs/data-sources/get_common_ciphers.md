---
page_title: "gigavuecore_get_common_ciphers Data Source - gigavuecore"
subcategory: ""
description: |-
  Get common (intersection) ciphers by protocol type
---

# gigavuecore_get_common_ciphers Data Source

Get common (intersection) ciphers by protocol type

## Example Usage

```terraform
data "gigavuecore_get_common_ciphers" "example" {
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `type` (String, required) - Cipher protocol type

### Attributes

In addition to all arguments above, the following attributes are exported:

* `classic` (Object({tls1_2, tls1_3}), computed) - Intersection of supported TLS ciphers by protocol variant within a single device mode.
  * `tls1_2` (List(String), computed) - Common TLS 1.2 cipher suite names for this mode.
  * `tls1_3` (List(String), computed) - Common TLS 1.3 cipher suite names for this mode.
* `crypto` (Object({tls1_2, tls1_3}), computed) - Intersection of supported TLS ciphers by protocol variant within a single device mode.
  * `tls1_2` (List(String), computed) - Common TLS 1.2 cipher suite names for this mode.
  * `tls1_3` (List(String), computed) - Common TLS 1.3 cipher suite names for this mode.
* `fips` (Object({tls1_2, tls1_3}), computed) - Intersection of supported TLS ciphers by protocol variant within a single device mode.
  * `tls1_2` (List(String), computed) - Common TLS 1.2 cipher suite names for this mode.
  * `tls1_3` (List(String), computed) - Common TLS 1.3 cipher suite names for this mode.

