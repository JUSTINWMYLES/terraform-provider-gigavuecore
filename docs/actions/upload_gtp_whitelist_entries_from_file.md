---
page_title: "gigavuecore_upload_gtp_whitelist_entries_from_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload GTP Whitelist Entries from file
---

# gigavuecore_upload_gtp_whitelist_entries_from_file Action

Upload GTP Whitelist Entries from file

## Example Usage

```terraform
action "gigavuecore_upload_gtp_whitelist_entries_from_file" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    entries = "example"
    usage = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GTP Whitelist
* `cluster_id` (String, required) - Target Cluster ID
* `entries` (String, required) - File containing GTP Whitelist Entries (one entry per line)
* `usage` (String, required) - When 'delete' is specified, the uploaded list is treated as a delete request and all the entries in the list should be removed from the existing whitelist identified by the 'alias' parameter
