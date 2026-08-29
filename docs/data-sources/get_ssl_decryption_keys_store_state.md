---
page_title: "gigavuecore_get_ssl_decryption_keys_store_state Data Source - gigavuecore"
subcategory: ""
description: |-
  Get keystore keychain State
---

# gigavuecore_get_ssl_decryption_keys_store_state Data Source

Get keystore keychain State

## Example Usage

```terraform
data "gigavuecore_get_ssl_decryption_keys_store_state" "example" {
  cluster_id = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID

### Attributes

In addition to all arguments above, the following attributes are exported:

* `locked` (Boolean, computed) - indicates whether SSL Decryption Key Store is unlocked
* `password_set` (Boolean, computed) - indicates whether SSL Decryption Key Store password is set


