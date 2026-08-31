---
page_title: "gigavuecore_upload_inline_ssl_signing Action - gigavuecore"
subcategory: ""
description: |-
  Upload inline SSL signing
---

# gigavuecore_upload_inline_ssl_signing Action

Upload inline SSL signing

## Example Usage

```terraform
action "gigavuecore_upload_inline_ssl_signing" "example" {
  config {
    cluster_id  = "example"
    key_alias   = "example"
    signing_for = "primary"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `key_alias` (String, required) - Keystore key alias (from apps/keystore/keys)
* `signing_for` (String, required)


