---
page_title: "gigavuecore_fetch_inline_ssl_profile_list Action - gigavuecore"
subcategory: ""
description: |-
  Fetch inline SSL profile nodecryptlist or decryptlist
---

# gigavuecore_fetch_inline_ssl_profile_list Action

Fetch inline SSL profile nodecryptlist or decryptlist

## Example Usage

```terraform
action "gigavuecore_fetch_inline_ssl_profile_list" "example" {
  config {
    alias       = "example"
    cluster_id  = "example"
    file_source = null
    list        = "example"
    list_type   = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `file_source` (Dynamic, optional) - Remote file source or destination
* `list` (String, optional) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'
* `list_type` (String, required) - specify a nodecryptlist or decryptlist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)


