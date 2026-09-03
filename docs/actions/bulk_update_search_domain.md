---
page_title: "gigavuecore_bulk_update_search_domain Action - gigavuecore"
subcategory: ""
description: |-
  update search domains
---

# gigavuecore_bulk_update_search_domain Action

update search domains

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_bulk_update_search_domain" "example" {
  config {
    domain_names = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `domain_names` (List of String, required) - List of search domains


