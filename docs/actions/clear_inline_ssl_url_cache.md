---
page_title: "gigavuecore_clear_inline_ssl_url_cache Action - gigavuecore"
subcategory: ""
description: |-
  Clear Inline SSL URL cache
---

# gigavuecore_clear_inline_ssl_url_cache Action

Clear Inline SSL URL cache

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_clear_inline_ssl_url_cache" "example" {
  config {
  }
}
```
