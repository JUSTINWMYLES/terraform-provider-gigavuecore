---
page_title: "gigavuecore_update_ssh_ciphers Action - gigavuecore"
subcategory: ""
description: |-
  Update all ciphers, Kex, Macs and HostKey
---

# gigavuecore_update_ssh_ciphers Action

Update all ciphers, Kex, Macs and HostKey

## Example Usage

```terraform
action "gigavuecore_update_ssh_ciphers" "example" {
  config {
    client_ciphers = [ "default" ]
    client_hostkey = [ "default" ]
    client_kex     = [ "default" ]
    client_macs    = [ "default" ]
    server_ciphers = [ "default" ]
    server_hostkey = [ "default" ]
    server_kex     = [ "default" ]
    server_macs    = [ "default" ]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `client_ciphers` (List of String, required)
* `client_hostkey` (List of String, required)
* `client_kex` (List of String, required)
* `client_macs` (List of String, required)
* `server_ciphers` (List of String, required)
* `server_hostkey` (List of String, required)
* `server_kex` (List of String, required)
* `server_macs` (List of String, required)


