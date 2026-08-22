---
page_title: "gigavuecore_reset_ssl_decryption_keys Action - gigavuecore"
subcategory: ""
description: |-
  Reset keystore keychain password and invalidate all keys
---

# gigavuecore_reset_ssl_decryption_keys Action

Reset keystore keychain password and invalidate all keys

## Example Usage

```terraform
action "gigavuecore_reset_ssl_decryption_keys" "example" {
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
