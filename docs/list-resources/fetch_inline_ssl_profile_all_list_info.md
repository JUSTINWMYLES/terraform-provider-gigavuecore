---
page_title: "gigavuecore_fetch_inline_ssl_profile_all_list_info List Resource - gigavuecore"
subcategory: ""
description: |-
  Fetch inline SSL profile nodecryptlist and decryptlist information
---

# gigavuecore_fetch_inline_ssl_profile_all_list_info List Resource

Fetch inline SSL profile nodecryptlist and decryptlist information

## Example Usage

```terraform
list "gigavuecore_fetch_inline_ssl_profile_all_list_info" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile


