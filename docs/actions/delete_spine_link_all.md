---
page_title: "gigavuecore_delete_spine_link_all Action - gigavuecore"
subcategory: ""
description: |-
  Delete all spine-link configurations
---

# gigavuecore_delete_spine_link_all Action

Delete all spine-link configurations

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_spine_link_all" "example" {
  config {
  }
}
```
