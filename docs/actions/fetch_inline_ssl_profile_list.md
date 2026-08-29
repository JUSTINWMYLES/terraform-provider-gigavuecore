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
    alias      = "example"
    cluster_id = "example"
    file_source = {
      hostname = "example"
      password = "example"
      path     = "example"
      protocol = "example"
      username = "example"
    }
    list      = "example"
    list_type = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `file_source` (Attributes, optional) - Remote file source or destination (see [below for nested schema](#nestedatt--file_source))
* `list` (String, optional) - The nodecryptlist or decryptlist . Mutually exclusive with 'fileSource'
* `list_type` (String, required) - specify a nodecryptlist or decryptlist. (deprecated: use nodecryptlist and decryptlist instead of whitelist and blacklist)

<a id="nestedatt--file_source"></a>
### Nested Schema for `file_source`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file

Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

