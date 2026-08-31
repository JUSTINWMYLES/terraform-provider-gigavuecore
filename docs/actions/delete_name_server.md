---
page_title: "gigavuecore_delete_name_server Action - gigavuecore"
subcategory: ""
description: |-
  Delete name server by address
---

# gigavuecore_delete_name_server Action

Delete name server by address

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_name_server" "example" {
  config {
    server_ip = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `server_ip` (String, required) - address of the nameserver to be deleted


