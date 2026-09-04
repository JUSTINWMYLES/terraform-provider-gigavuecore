---
page_title: "gigavuecore_reset_global_template_defaults Action - gigavuecore"
subcategory: ""
description: |-
  new in FM 5.14
---

# gigavuecore_reset_global_template_defaults Action

new in FM 5.14

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_reset_global_template_defaults" "example" {
  config {
    config_type = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `config_type` (String, required) - configType of the fm template


