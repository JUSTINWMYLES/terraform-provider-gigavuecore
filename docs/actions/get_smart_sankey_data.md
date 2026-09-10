---
page_title: "gigavuecore_get_smart_sankey_data Action - gigavuecore"
subcategory: ""
description: |-
  FM's best guess sankey constructed based on nodes' neighborship count
---

# gigavuecore_get_smart_sankey_data Action

FM's best guess sankey constructed based on nodes' neighborship count

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_get_smart_sankey_data" "example" {
  config {
  }
}
```
