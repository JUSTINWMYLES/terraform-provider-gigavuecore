---
page_title: "gigavuecore_add_search_domain Action - gigavuecore"
subcategory: ""
description: |-
  Add search domain
---

# gigavuecore_add_search_domain Action

Add search domain

## Example Usage

```terraform
action "gigavuecore_add_search_domain" "example" {
  config {
    domain_names = [ "example" ]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `domain_names` (List of String, required) - List of search domains


