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

