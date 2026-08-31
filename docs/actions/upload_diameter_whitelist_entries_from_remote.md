---
page_title: "gigavuecore_upload_diameter_whitelist_entries_from_remote Action - gigavuecore"
subcategory: ""
description: |-
  Upload Diameter Whitelist Entries from Remote
---

# gigavuecore_upload_diameter_whitelist_entries_from_remote Action

Upload Diameter Whitelist Entries from Remote

## Example Usage

```terraform
action "gigavuecore_upload_diameter_whitelist_entries_from_remote" "example" {
  config {
    alias = "example"
    remote = {
      hostname = "example"
      password = "example"
      path     = "example"
      protocol = "scp"
      username = "example"
    }
    usage           = "add"
    whitelist_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target Diameter Whitelist
* `remote` (Attributes, required) - Remote file source or destination (see [below for nested schema](#nestedatt--remote))
* `usage` (String, optional) - When 'delete' is specified, the uploaded list is treated as a delete request and all the entries in the list should be removed from the existing whitelist identified by the 'alias' parameter
* `whitelist_alias` (String, required) - alias of the target Diameter Whitelist

<a id="nestedatt--remote"></a>
### Nested Schema for `remote`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file

Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

