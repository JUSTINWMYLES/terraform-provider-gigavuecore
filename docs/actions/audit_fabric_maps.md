---
page_title: "gigavuecore_audit_fabric_maps Action - gigavuecore"
subcategory: ""
description: |-
  Audit all user-defined fabric maps
---

# gigavuecore_audit_fabric_maps Action

Audit all user-defined fabric maps

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_audit_fabric_maps" "example" {
  config {
  }
}
```
