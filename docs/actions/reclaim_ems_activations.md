---
page_title: "gigavuecore_reclaim_ems_activations Action - gigavuecore"
subcategory: ""
description: |-
  Reclaim a license activation, return to entitlement the associated quantity(number of licenses)
---

# gigavuecore_reclaim_ems_activations Action

Reclaim a license activation, return to entitlement the associated quantity(number of licenses)

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_reclaim_ems_activations" "example" {
  config {
    aid = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `aid` (String, required) - Activation ID


