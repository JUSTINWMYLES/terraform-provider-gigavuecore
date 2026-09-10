---
page_title: "gigavuecore_delete_deploy Action - gigavuecore"
subcategory: ""
description: |-
  Delete unified resource deployment by environment and unified resource id
---

# gigavuecore_delete_deploy Action

Delete unified resource deployment by environment and unified resource id

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_deploy" "example" {
  config {
    env_id   = "example"
    unify_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `env_id` (String, required) - unified environment identifier
* `unify_id` (String, required) - unified resource identifier


