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
    client_ciphers = [ "example" ]
    client_hostkey = [ "example" ]
    client_kex = [ "example" ]
    client_macs = [ "example" ]
    server_ciphers = [ "example" ]
    server_hostkey = [ "example" ]
    server_kex = [ "example" ]
    server_macs = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `client_ciphers` (List(String), required)
* `client_hostkey` (List(String), required)
* `client_kex` (List(String), required)
* `client_macs` (List(String), required)
* `server_ciphers` (List(String), required)
* `server_hostkey` (List(String), required)
* `server_kex` (List(String), required)
* `server_macs` (List(String), required)
