---
page_title: "gigavuecore_bulk_update_search_domain Action - gigavuecore"
subcategory: ""
description: |-
  update search domains
---

# gigavuecore_bulk_update_search_domain Action

update search domains

## Example Usage

```terraform
action "gigavuecore_bulk_update_search_domain" "example" {
  config {
    domain_names = [ "example" ]
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `domain_names` (List(String), required) - List of search domains
