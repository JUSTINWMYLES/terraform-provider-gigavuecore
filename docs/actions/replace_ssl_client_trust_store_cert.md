---
page_title: "gigavuecore_replace_ssl_client_trust_store_cert Action - gigavuecore"
subcategory: ""
description: |-
  Replace certificate to the SSL client trust-store
---

# gigavuecore_replace_ssl_client_trust_store_cert Action

Replace certificate to the SSL client trust-store

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

~> **Warning:** This action accepts attributes whose names indicate secrets (password), but action schemas cannot mark attributes Sensitive, so their values are stored and displayed in plain text. Avoid passing real secrets through these attributes.

## Example Usage

```terraform
action "gigavuecore_replace_ssl_client_trust_store_cert" "example" {
  config {
    alias      = "example"
    body_alias = "example"
    cluster_id = "example"
    file       = "example"
    file_source = {
      hostname = "example"
      password = "example"
      path     = "example"
      protocol = "scp"
      username = "example"
    }
    type = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Client Trust Store alias
* `body_alias` (String, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Attributes, optional) - Remote file source or destination (see [below for nested schema](#nestedatt--file_source))
* `type` (String, optional)

<a id="nestedatt--file_source"></a>
### Nested Schema for `file_source`

Required:

* `hostname` (String) - server address
* `path` (String) - file path on server
* `protocol` (String) - Protocol to access server. HTTP and HTTPS only applicable for retrieving file

Optional:

* `password` (String) - password to use for server login
* `username` (String) - user name to use for server login

