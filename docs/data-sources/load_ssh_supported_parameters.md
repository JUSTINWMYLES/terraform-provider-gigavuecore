---
page_title: "gigavuecore_load_ssh_supported_parameters Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all SSH supported Parameters
---

# gigavuecore_load_ssh_supported_parameters Data Source

Load all SSH supported Parameters

## Example Usage

```terraform
data "gigavuecore_load_ssh_supported_parameters" "example" {
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

