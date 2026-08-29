---
page_title: "gigavuecore_generate_keystore_key Action - gigavuecore"
subcategory: ""
description: |-
  generate a self-signed key in the keystore
---

# gigavuecore_generate_keystore_key Action

generate a self-signed key in the keystore

## Example Usage

```terraform
action "gigavuecore_generate_keystore_key" "example" {
  config {
    alias       = "example"
    cluster_id  = "example"
    common_name = "example"
    country     = "example"
    days        = 0
    hash_type   = "example"
    keysize     = 0
    org_name    = "example"
    org_unit    = "example"
    state       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required)
* `cluster_id` (String, required) - Target Cluster ID
* `common_name` (String, required)
* `country` (String, optional)
* `days` (Number, optional)
* `hash_type` (String, optional)
* `keysize` (Number, optional) - ecdsa sizes: 256, 384, 521; rsa sizes: 1024, 2048, 4096
* `org_name` (String, required)
* `org_unit` (String, optional)
* `state` (String, optional)


