---
page_title: "gigavuecore_fetch_inline_ssl_profile_all_list_info Data Source - gigavuecore"
subcategory: ""
description: |-
  Fetch inline SSL profile nodecryptlist and decryptlist information
---

# gigavuecore_fetch_inline_ssl_profile_all_list_info Data Source

Fetch inline SSL profile nodecryptlist and decryptlist information

## Example Usage

```terraform
data "gigavuecore_fetch_inline_ssl_profile_all_list_info" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile

### Attributes

In addition to all arguments above, the following attributes are exported:

* `blacklist` (Object({num_entries}), computed)
  * `num_entries` (Number, computed) - the number of entries in the whitelist or blacklist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)
* `decryptlist` (Object({num_entries}), computed)
  * `num_entries` (Number, computed) - the number of entries in the nodecryptlist or decryptlist
* `nodecryptlist` (Object({num_entries}), computed)
  * `num_entries` (Number, computed) - the number of entries in the nodecryptlist or decryptlist
* `whitelist` (Object({num_entries}), computed)
  * `num_entries` (Number, computed) - the number of entries in the whitelist or blacklist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)

