---
page_title: "gigavuecore_add_search_domain Action - gigavuecore"
subcategory: ""
description: |-
  Add search domain
---

# gigavuecore_add_search_domain Action

Add search domain

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_add_search_domain" "example" {
  config {
    domain_names = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `domain_names` (List of String, required) - List of search domains


