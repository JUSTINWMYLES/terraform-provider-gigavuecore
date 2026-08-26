---
page_title: "gigavuecore_loads_global_fm_template_with_values Data Source - gigavuecore"
subcategory: ""
description: |-
  new in FM 6.6
---

# gigavuecore_loads_global_fm_template_with_values Data Source

new in FM 6.6

## Example Usage

```terraform
data "gigavuecore_loads_global_fm_template_with_values" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `classic` (Attributes, computed) - System SSH Cipher (see [below for nested schema](#nestedatt--classic))
* `crypto` (Attributes, computed) - System SSH Cipher (see [below for nested schema](#nestedatt--crypto))
* `fips` (Attributes, computed) - System SSH Cipher (see [below for nested schema](#nestedatt--fips))

<a id="nestedatt--classic"></a>
### Nested Schema for `classic`

Read-Only:

* `client_ciphers` (List of String)
* `client_hostkey` (List of String)
* `client_kex` (List of String)
* `client_macs` (List of String)
* `server_ciphers` (List of String)
* `server_hostkey` (List of String)
* `server_kex` (List of String)
* `server_macs` (List of String)
<a id="nestedatt--crypto"></a>
### Nested Schema for `crypto`

Read-Only:

* `client_ciphers` (List of String)
* `client_hostkey` (List of String)
* `client_kex` (List of String)
* `client_macs` (List of String)
* `server_ciphers` (List of String)
* `server_hostkey` (List of String)
* `server_kex` (List of String)
* `server_macs` (List of String)
<a id="nestedatt--fips"></a>
### Nested Schema for `fips`

Read-Only:

* `client_ciphers` (List of String)
* `client_hostkey` (List of String)
* `client_kex` (List of String)
* `client_macs` (List of String)
* `server_ciphers` (List of String)
* `server_hostkey` (List of String)
* `server_kex` (List of String)
* `server_macs` (List of String)

