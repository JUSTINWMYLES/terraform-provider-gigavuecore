---
page_title: "gigavuecore_run_integrity_check Action - gigavuecore"
subcategory: ""
description: |-
  Run FM Integrity check.
---

# gigavuecore_run_integrity_check Action

Run FM Integrity check.

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_run_integrity_check" "example" {
  config {
  }
}
```
