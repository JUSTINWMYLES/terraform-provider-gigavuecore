---
page_title: "gigavuecore_replace_ssl_client_trust_store_cert Action - gigavuecore"
subcategory: ""
description: |-
  Replace certificate to the SSL client trust-store
---

# gigavuecore_replace_ssl_client_trust_store_cert Action

Replace certificate to the SSL client trust-store

## Example Usage

```terraform
action "gigavuecore_replace_ssl_client_trust_store_cert" "example" {
  config {
    alias       = "example"
    body_alias  = "example"
    cluster_id  = "example"
    file        = "example"
    file_source = null
    type        = "example"
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
* `file_source` (Dynamic, optional) - Remote file source or destination
* `type` (String, optional)


