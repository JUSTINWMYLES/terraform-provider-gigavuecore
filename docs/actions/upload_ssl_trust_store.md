---
page_title: "gigavuecore_upload_ssl_trust_store Action - gigavuecore"
subcategory: ""
description: |-
  Upload inline SSL trust-store
---

# gigavuecore_upload_ssl_trust_store Action

Upload inline SSL trust-store

## Example Usage

```terraform
action "gigavuecore_upload_ssl_trust_store" "example" {
  config {
    cluster_id = "example"
    file       = "example"
    file_source = {
      hostname = "example"
      password = "example"
      path     = "example"
      protocol = "scp"
      username = "example"
    }
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
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

