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

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `classic` (Object({client_ciphers, client_hostkey, client_kex, client_macs, server_ciphers, server_hostkey, server_kex, server_macs}), computed) - System SSH Cipher
  * `client_ciphers` (List(String), computed)
  * `client_hostkey` (List(String), computed)
  * `client_kex` (List(String), computed)
  * `client_macs` (List(String), computed)
  * `server_ciphers` (List(String), computed)
  * `server_hostkey` (List(String), computed)
  * `server_kex` (List(String), computed)
  * `server_macs` (List(String), computed)
* `crypto` (Object({client_ciphers, client_hostkey, client_kex, client_macs, server_ciphers, server_hostkey, server_kex, server_macs}), computed) - System SSH Cipher
  * `client_ciphers` (List(String), computed)
  * `client_hostkey` (List(String), computed)
  * `client_kex` (List(String), computed)
  * `client_macs` (List(String), computed)
  * `server_ciphers` (List(String), computed)
  * `server_hostkey` (List(String), computed)
  * `server_kex` (List(String), computed)
  * `server_macs` (List(String), computed)
* `fips` (Object({client_ciphers, client_hostkey, client_kex, client_macs, server_ciphers, server_hostkey, server_kex, server_macs}), computed) - System SSH Cipher
  * `client_ciphers` (List(String), computed)
  * `client_hostkey` (List(String), computed)
  * `client_kex` (List(String), computed)
  * `client_macs` (List(String), computed)
  * `server_ciphers` (List(String), computed)
  * `server_hostkey` (List(String), computed)
  * `server_kex` (List(String), computed)
  * `server_macs` (List(String), computed)

