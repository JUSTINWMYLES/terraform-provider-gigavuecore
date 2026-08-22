---
page_title: "gigavuecore_upload_diameter_whitelist_entries_from_remote Action - gigavuecore"
subcategory: ""
description: |-
  Upload Diameter Whitelist Entries from Remote
---

# gigavuecore_upload_diameter_whitelist_entries_from_remote Action

Upload Diameter Whitelist Entries from Remote

## Example Usage

```terraform
action "gigavuecore_upload_diameter_whitelist_entries_from_remote" "example" {
  config {
    alias = "example"
    remote = "example"
    usage = "example"
    whitelist_alias = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Diameter Whitelist
* `remote` (Dynamic, required) - Remote file source or destination
* `usage` (String, optional) - When 'delete' is specified, the uploaded list is treated as a delete request and all the entries in the list should be removed from the existing whitelist identified by the 'alias' parameter
* `whitelist_alias` (String, required) - alias of the target Diameter Whitelist
