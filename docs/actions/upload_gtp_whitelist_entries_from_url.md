---
page_title: "gigavuecore_upload_gtp_whitelist_entries_from_url Action - gigavuecore"
subcategory: ""
description: |-
  Upload GTP Whitelist Entries from URL
---

# gigavuecore_upload_gtp_whitelist_entries_from_url Action

Upload GTP Whitelist Entries from URL

## Example Usage

```terraform
action "gigavuecore_upload_gtp_whitelist_entries_from_url" "example" {
  config {
    alias           = "example"
    cluster_id      = "example"
    url             = "example"
    usage           = "example"
    whitelist_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target GTP Whitelist
* `cluster_id` (String, required) - Target Cluster ID
* `url` (String, required) - URL of the Whitelist file. format: protocol://\[username\[:password\]\]@hostname\[:port\]/path/filename
* `usage` (String, optional) - When 'delete' is specified, the uploaded list is treated as a delete request and all the entries in the list should be removed from the existing whitelist identified by the 'alias' parameter
* `whitelist_alias` (String, required) - alias of the target GTP Whitelist


