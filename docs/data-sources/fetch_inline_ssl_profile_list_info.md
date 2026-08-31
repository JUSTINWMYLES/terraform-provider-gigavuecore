---
page_title: "gigavuecore_fetch_inline_ssl_profile_list_info Data Source - gigavuecore"
subcategory: ""
description: |-
  Fetch inline SSL profile nodecryptlist or decryptlist information
---

# gigavuecore_fetch_inline_ssl_profile_list_info Data Source

Fetch inline SSL profile nodecryptlist or decryptlist information

## Example Usage

```terraform
data "gigavuecore_fetch_inline_ssl_profile_list_info" "example" {
  alias     = "example"
  list_type = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `list_type` (String, required) - specify a nodecryptlist or decryptlist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `num_entries` (Number, computed) - the number of entries in the nodecryptlist or decryptlist


