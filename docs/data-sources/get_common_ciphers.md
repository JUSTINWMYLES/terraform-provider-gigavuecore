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

* `classic` (Attributes, computed) - Intersection of supported TLS ciphers by protocol variant within a single device mode. (see [below for nested schema](#nestedatt--classic))
* `crypto` (Attributes, computed) - Intersection of supported TLS ciphers by protocol variant within a single device mode. (see [below for nested schema](#nestedatt--crypto))
* `fips` (Attributes, computed) - Intersection of supported TLS ciphers by protocol variant within a single device mode. (see [below for nested schema](#nestedatt--fips))

<a id="nestedatt--classic"></a>
### Nested Schema for `classic`

Read-Only:

* `tls1_2` (List of String) - Common TLS 1.2 cipher suite names for this mode.
* `tls1_3` (List of String) - Common TLS 1.3 cipher suite names for this mode.
<a id="nestedatt--crypto"></a>
### Nested Schema for `crypto`

Read-Only:

* `tls1_2` (List of String) - Common TLS 1.2 cipher suite names for this mode.
* `tls1_3` (List of String) - Common TLS 1.3 cipher suite names for this mode.
<a id="nestedatt--fips"></a>
### Nested Schema for `fips`

Read-Only:

* `tls1_2` (List of String) - Common TLS 1.2 cipher suite names for this mode.
* `tls1_3` (List of String) - Common TLS 1.3 cipher suite names for this mode.

