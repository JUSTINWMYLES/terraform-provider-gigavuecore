---
page_title: "gigavuecore_delete_env Action - gigavuecore"
subcategory: ""
description: |-
  Delete unified resource environment
---

# gigavuecore_delete_env Action

Delete unified resource environment

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_env" "example" {
  config {
    env_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - envId is the unified resource environment id


