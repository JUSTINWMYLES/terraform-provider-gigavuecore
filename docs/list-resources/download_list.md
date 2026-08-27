---
page_title: "gigavuecore_download_list List Resource - gigavuecore"
subcategory: ""
description: |-
  Download a nodecryptlist or decryptlist
---

# gigavuecore_download_list List Resource

Download a nodecryptlist or decryptlist

## Example Usage

```terraform
list "gigavuecore_download_list" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    alias      = "example"
    cluster_id = "example"
    list_type  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `list_type` (String, required) - specify a nodecryptlist or decryptlist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)


