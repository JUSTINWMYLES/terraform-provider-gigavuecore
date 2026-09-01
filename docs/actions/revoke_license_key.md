---
page_title: "gigavuecore_revoke_license_key Action - gigavuecore"
subcategory: ""
description: |-
  Revoke Node-Locked Device License
---

# gigavuecore_revoke_license_key Action

Revoke Node-Locked Device License

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_revoke_license_key" "example" {
  config {
    activation_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `activation_id` (String, required) - ID of the feature activation


