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

* `blacklist` (Attributes, computed) (see [below for nested schema](#nestedatt--blacklist))
* `decryptlist` (Attributes, computed) (see [below for nested schema](#nestedatt--decryptlist))
* `nodecryptlist` (Attributes, computed) (see [below for nested schema](#nestedatt--nodecryptlist))
* `whitelist` (Attributes, computed) (see [below for nested schema](#nestedatt--whitelist))

<a id="nestedatt--blacklist"></a>
### Nested Schema for `blacklist`

Read-Only:

* `num_entries` (Number) - the number of entries in the whitelist or blacklist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)
<a id="nestedatt--decryptlist"></a>
### Nested Schema for `decryptlist`

Read-Only:

* `num_entries` (Number) - the number of entries in the nodecryptlist or decryptlist
<a id="nestedatt--nodecryptlist"></a>
### Nested Schema for `nodecryptlist`

Read-Only:

* `num_entries` (Number) - the number of entries in the nodecryptlist or decryptlist
<a id="nestedatt--whitelist"></a>
### Nested Schema for `whitelist`

Read-Only:

* `num_entries` (Number) - the number of entries in the whitelist or blacklist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)

