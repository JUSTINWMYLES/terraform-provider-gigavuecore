---
page_title: "gigavuecore_reset_ssl_decryption_keys Action - gigavuecore"
subcategory: ""
description: |-
  Reset keystore keychain password and invalidate all keys
---

# gigavuecore_reset_ssl_decryption_keys Action

Reset keystore keychain password and invalidate all keys

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (ksPassword), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_reset_ssl_decryption_keys" "example" {
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


