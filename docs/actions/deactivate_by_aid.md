---
page_title: "gigavuecore_deactivate_by_aid Action - gigavuecore"
subcategory: ""
description: |-
  Deactivate a license (floating or VBL) by its Activation ID; floating licenses still assigned to at least one card or chassis will be skipped
---

# gigavuecore_deactivate_by_aid Action

Deactivate a license (floating or VBL) by its Activation ID; floating licenses still assigned to at least one card or chassis will be skipped

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Note:** This action is not yet wired to a remote API endpoint. Invoking it fails with an explicit "not wired" diagnostic instead of calling the API. The OpenAPI operations it was inferred from could not be resolved into a complete mapping; consult the eidos generation warnings for the exact cause.

## Example Usage

```terraform
action "gigavuecore_deactivate_by_aid" "example" {
  config {
    aid = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `aid` (String, required) - Activation ID (created when license is generated) of the feature activation


