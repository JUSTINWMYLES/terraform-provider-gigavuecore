---
page_title: "gigavuecore_undeploy_policy Action - gigavuecore"
subcategory: ""
description: |-
  Undeploy policy
---

# gigavuecore_undeploy_policy Action

Undeploy policy

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_undeploy_policy" "example" {
  config {
    name = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `name` (String, required) - policy name


