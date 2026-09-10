---
page_title: "gigavuecore_delete_permitted_address Action - gigavuecore"
subcategory: ""
description: |-
  Delete Permitted Address
---

# gigavuecore_delete_permitted_address Action

Delete Permitted Address

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_permitted_address" "example" {
  config {
    address = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `address` (String, required) - Permitted Address


