---
page_title: "gigavuecore_upload_keystore_key Action - gigavuecore"
subcategory: ""
description: |-
  upload a key to the keystore from local file
---

# gigavuecore_upload_keystore_key Action

upload a key to the keystore from local file

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

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


