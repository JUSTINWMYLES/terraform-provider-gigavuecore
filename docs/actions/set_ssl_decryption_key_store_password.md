---
page_title: "gigavuecore_set_ssl_decryption_key_store_password Action - gigavuecore"
subcategory: ""
description: |-
  Set keystore keychain password
---

# gigavuecore_set_ssl_decryption_key_store_password Action

Set keystore keychain password

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (ksPassword), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_set_ssl_decryption_key_store_password" "example" {
  config {
    cluster_id  = "example"
    ks_password = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `ks_password` (String, required) - Only valid for 'set' operations


