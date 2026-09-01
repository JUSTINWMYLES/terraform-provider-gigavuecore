---
page_title: "gigavuecore_get_ssh_ciphers Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all ciphers, Kex, Macs and HostKey
---

# gigavuecore_get_ssh_ciphers Data Source

Load all ciphers, Kex, Macs and HostKey

## Example Usage

```terraform
data "gigavuecore_get_ssh_ciphers" "example" {
}
```

## Schema

### Attributes

In addition to all arguments above, the following attributes are exported:

* `client_ciphers` (List of String, computed)
* `client_hostkey` (List of String, computed)
* `client_kex` (List of String, computed)
* `client_macs` (List of String, computed)
* `server_ciphers` (List of String, computed)
* `server_hostkey` (List of String, computed)
* `server_kex` (List of String, computed)
* `server_macs` (List of String, computed)


