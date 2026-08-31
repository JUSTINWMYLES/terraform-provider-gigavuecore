---
page_title: "gigavuecore_delete_search_domain Action - gigavuecore"
subcategory: ""
description: |-
  Delete search domain server by name
---

# gigavuecore_delete_search_domain Action

Delete search domain server by name

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_search_domain" "example" {
  config {
    domain_name = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `domain_name` (String, required) - name of the search domain to be deleted


