---
page_title: "gigavuecore_unlock_ssl_decryption_key_store Action - gigavuecore"
subcategory: ""
description: |-
  Unlock keystore keychain
---

# gigavuecore_unlock_ssl_decryption_key_store Action

Unlock keystore keychain

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (ksPassword), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_unlock_ssl_decryption_key_store" "example" {
  config {
    auto_login  = true
    cluster_id  = "example"
    ks_password = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `auto_login` (Boolean, optional) - If enabled, On every node reboot FM would auto login into device with the password configured.
* `cluster_id` (String, required) - Target Cluster ID
* `ks_password` (String, required) - Only valid for 'set' operations


