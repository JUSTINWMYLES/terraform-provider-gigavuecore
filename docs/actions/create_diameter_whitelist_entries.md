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
    entries = [{
      active_sessions = 0
      user_name       = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Diameter Whitelist
* `entries` (Attributes List, required) (see [below for nested schema](#nestedatt--entries))

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Optional:

* `active_sessions` (Number) - Number of active sessions
* `user_name` (String)

