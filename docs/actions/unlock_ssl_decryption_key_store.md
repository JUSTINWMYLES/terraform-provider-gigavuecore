---
page_title: "gigavuecore_unlock_ssl_decryption_key_store Action - gigavuecore"
subcategory: ""
description: |-
  Unlock keystore keychain
---

# gigavuecore_unlock_ssl_decryption_key_store Action

Unlock keystore keychain

## Example Usage

```terraform
action "gigavuecore_unlock_ssl_decryption_key_store" "example" {
  config {
    auto_login  = true
    cluster_id  = "example"
    ks_password = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `auto_login` (Boolean, optional) - If enabled, On every node reboot FM would auto login into device with the password configured.
* `cluster_id` (String, required) - Target Cluster ID
* `ks_password` (String, required) - Only valid for 'set' operations


