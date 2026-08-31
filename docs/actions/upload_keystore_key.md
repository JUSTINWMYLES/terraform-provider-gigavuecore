---
page_title: "gigavuecore_upload_keystore_key Action - gigavuecore"
subcategory: ""
description: |-
  upload a key to the keystore from local file
---

# gigavuecore_upload_keystore_key Action

upload a key to the keystore from local file

## Example Usage

```terraform
action "gigavuecore_upload_keystore_key" "example" {
  config {
    alias      = "example"
    body_alias = "example"
    cluster_id = "example"
    comment    = "example"
    file       = "example"
    passphrase = "example"
    type       = "private"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias for the key
* `body_alias` (String, required) - alias for the key
* `cluster_id` (String, required) - Target Cluster ID
* `comment` (String, optional) - key comment
* `file` (String, required) - file to upload to device
* `passphrase` (String, optional) - passphrase applicable for pkcs12 only
* `type` (String, required) - key type


