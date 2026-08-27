---
page_title: "gigavuecore_create_ssl_client_trust_store Action - gigavuecore"
subcategory: ""
description: |-
  Create the SSL Client trust-store
---

# gigavuecore_create_ssl_client_trust_store Action

Create the SSL Client trust-store

## Example Usage

```terraform
action "gigavuecore_create_ssl_client_trust_store" "example" {
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


