---
page_title: "gigavuecore_enable_crypto_mode Action - gigavuecore"
subcategory: ""
description: |-
  Enable crypto mode.
---

# gigavuecore_enable_crypto_mode Action

Enable crypto mode.

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_enable_crypto_mode" "example" {
  config {
  }
}
```
