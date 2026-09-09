---
page_title: "gigavuecore_delete_hsm_from_hsm_group Action - gigavuecore"
subcategory: ""
description: |-
  Delete HSM alias from HSM Group
---

# gigavuecore_delete_hsm_from_hsm_group Action

Delete HSM alias from HSM Group

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_hsm_from_hsm_group" "example" {
  config {
    alias     = "example"
    hsm_alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `hsm_alias` (String, required) - HSM alias


