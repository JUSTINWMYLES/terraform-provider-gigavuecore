---
page_title: "gigavuecore_delete_keystore_keys Action - gigavuecore"
subcategory: ""
description: |-
  delete all keystore keys
---

# gigavuecore_delete_keystore_keys Action

delete all keystore keys

## Example Usage

```terraform
action "gigavuecore_delete_keystore_keys" "example" {
  config {
    cluster_id = "example"
    ecdsa_key = true
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `ecdsa_key` (Bool, optional) - If set to True, delete all ECDSA keys. If set to False, delete all RSA keys. If query param not specified, delete all keys
