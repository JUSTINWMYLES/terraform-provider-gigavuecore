---
page_title: "gigavuecore_clear_inline_ssl_cert_valid_cache Action - gigavuecore"
subcategory: ""
description: |-
  Clear Inline SSL certificate validation cache
---

# gigavuecore_clear_inline_ssl_cert_valid_cache Action

Clear Inline SSL certificate validation cache

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_inline_ssl_cert_valid_cache" "example" {
  config {
  }
}
```
