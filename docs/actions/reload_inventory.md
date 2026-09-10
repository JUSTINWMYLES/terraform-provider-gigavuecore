---
page_title: "gigavuecore_reload_inventory Action - gigavuecore"
subcategory: ""
description: |-
  Reload unified resource connection artifacts
---

# gigavuecore_reload_inventory Action

Reload unified resource connection artifacts

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_reload_inventory" "example" {
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


