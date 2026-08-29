---
page_title: "gigavuecore_add_ssl_trust_store_cert Action - gigavuecore"
subcategory: ""
description: |-
  Add certificate to the inline SSL trust-store
---

# gigavuecore_add_ssl_trust_store_cert Action

Add certificate to the inline SSL trust-store

## Example Usage

```terraform
action "gigavuecore_add_ssl_trust_store_cert" "example" {
  config {
    file = "example"
    file_source = {
      hostname = "example"
      password = "example"
      path     = "example"
      protocol = "example"
      username = "example"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Attributes, optional) - Remote file source or destination (see [below for nested schema](#nestedatt--file_source))

<a id="nestedatt--file_source"></a>
### Nested Schema for `file_source`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file
Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

