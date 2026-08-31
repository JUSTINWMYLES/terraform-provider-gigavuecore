---
page_title: "gigavuecore_register_or_de_register_nrt Action - gigavuecore"
subcategory: ""
description: |-
  Register or deregister near real-time statistics for a traffic flow
---

# gigavuecore_register_or_de_register_nrt Action

Register or deregister near real-time statistics for a traffic flow

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_register_or_de_register_nrt" "example" {
  config {
    alias     = "example"
    operation = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias
* `operation` (String, optional) - Operation to perform (add or delete)


