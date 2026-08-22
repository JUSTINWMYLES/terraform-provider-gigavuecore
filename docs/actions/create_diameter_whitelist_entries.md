---
page_title: "gigavuecore_create_diameter_whitelist_entries Action - gigavuecore"
subcategory: ""
description: |-
  Create Diameter Whitelist Entries
---

# gigavuecore_create_diameter_whitelist_entries Action

Create Diameter Whitelist Entries

## Example Usage

```terraform
action "gigavuecore_create_diameter_whitelist_entries" "example" {
  config {
    alias = "example"
    entries = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Diameter Whitelist
* `entries` (List(Dynamic), required)
