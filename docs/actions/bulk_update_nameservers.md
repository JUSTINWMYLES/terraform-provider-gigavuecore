---
page_title: "gigavuecore_bulk_update_nameservers Action - gigavuecore"
subcategory: ""
description: |-
  update name servers
---

# gigavuecore_bulk_update_nameservers Action

update name servers

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_bulk_update_nameservers" "example" {
  config {
    body_interface_name = "example"
    interface_name      = "example"
    servers             = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `body_interface_name` (String, required) - Interface in which the nameserver resides
* `interface_name` (String, required) - interfaceName for which name servers are updated
* `servers` (List of String, required) - List of name server address


