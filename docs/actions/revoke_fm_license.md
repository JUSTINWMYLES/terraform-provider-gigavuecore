---
page_title: "gigavuecore_revoke_fm_license Action - gigavuecore"
subcategory: ""
description: |-
  Revoke FM License
---

# gigavuecore_revoke_fm_license Action

Revoke FM License

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_revoke_fm_license" "example" {
  config {
    license_key = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `license_key` (String, required) - FM License Key to revoke


