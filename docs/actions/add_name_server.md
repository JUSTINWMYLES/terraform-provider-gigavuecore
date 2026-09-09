---
page_title: "gigavuecore_add_name_server Action - gigavuecore"
subcategory: ""
description: |-
  Add name server
---

# gigavuecore_add_name_server Action

Add name server

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_add_name_server" "example" {
  config {
    interface_name = "example"
    servers        = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `interface_name` (String, required) - Interface in which the nameserver resides
* `servers` (List of String, required) - List of name server address


