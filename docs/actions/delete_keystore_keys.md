---
page_title: "gigavuecore_delete_keystore_keys Action - gigavuecore"
subcategory: ""
description: |-
  delete all keystore keys
---

# gigavuecore_delete_keystore_keys Action

delete all keystore keys

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_keystore_keys" "example" {
  config {
    cluster_id = "example"
    ecdsa_key  = true
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `ecdsa_key` (Boolean, optional) - If set to True, delete all ECDSA keys. If set to False, delete all RSA keys. If query param not specified, delete all keys


