---
page_title: "gigavuecore_set_ssl_decryption_key_store_password Action - gigavuecore"
subcategory: ""
description: |-
  Set keystore keychain password
---

# gigavuecore_set_ssl_decryption_key_store_password Action

Set keystore keychain password

## Example Usage

```terraform
action "gigavuecore_set_ssl_decryption_key_store_password" "example" {
  config {
    cluster_id = "example"
    ks_password = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `ks_password` (String, required) - Only valid for 'set' operations
