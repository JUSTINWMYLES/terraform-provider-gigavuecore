---
page_title: "gigavuecore_upload_diameter_whitelist_entries_from_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload Diameter Whitelist Entries from file
---

# gigavuecore_upload_diameter_whitelist_entries_from_file Action

Upload Diameter Whitelist Entries from file

## Example Usage

```terraform
action "gigavuecore_upload_diameter_whitelist_entries_from_file" "example" {
  config {
    alias   = "example"
    entries = "example"
    usage   = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Diameter Whitelist
* `entries` (String, required) - File containing Diameter Whitelist Entries (one entry per line)
* `usage` (String, required) - When 'delete' is specified, the uploaded list is treated as a delete request and all the entries in the list should be removed from the existing whitelist identifed by the 'alias' parameter


