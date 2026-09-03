---
page_title: "gigavuecore_renew_bindings Action - gigavuecore"
subcategory: ""
description: |-
  Renew eligible licenses
---

# gigavuecore_renew_bindings Action

Renew eligible licenses

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_renew_bindings" "example" {
  config {
    renewee_id = "example"
    renewer_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `renewee_id` (String, required) - Renewee ID
* `renewer_id` (String, required) - Renewer ID


