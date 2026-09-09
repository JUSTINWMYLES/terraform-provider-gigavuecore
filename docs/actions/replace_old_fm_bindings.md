---
page_title: "gigavuecore_replace_old_fm_bindings Action - gigavuecore"
subcategory: ""
description: |-
  Replaces old FM's bindings on a new FM
---

# gigavuecore_replace_old_fm_bindings Action

Replaces old FM's bindings on a new FM

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_replace_old_fm_bindings" "example" {
  config {
    old_fm_vmac = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `old_fm_vmac` (String, required) - The VMAC of old FM


