---
page_title: "gigavuecore_delete_acme_server_details Action - gigavuecore"
subcategory: ""
description: |-
  Delete ACME server details
---

# gigavuecore_delete_acme_server_details Action

Delete ACME server details

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_acme_server_details" "example" {
  config {
    alias = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the acme server


