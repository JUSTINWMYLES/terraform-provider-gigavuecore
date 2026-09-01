---
page_title: "gigavuecore_re_apply_all_fabric_map Action - gigavuecore"
subcategory: ""
description: |-
  Re-apply the failed or pending configuration operations for all fabric maps in partial success, failed or pending state
---

# gigavuecore_re_apply_all_fabric_map Action

Re-apply the failed or pending configuration operations for all fabric maps in partial success, failed or pending state

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_re_apply_all_fabric_map" "example" {
  config {
  }
}
```
