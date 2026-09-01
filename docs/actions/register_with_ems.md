---
page_title: "gigavuecore_register_with_ems Action - gigavuecore"
subcategory: ""
description: |-
  Register FM with EMS
---

# gigavuecore_register_with_ems Action

Register FM with EMS

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (token), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_register_with_ems" "example" {
  config {
    token = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `token` (String, optional)


