---
page_title: "gigavuecore_upload_inline_ssl_trust_store Action - gigavuecore"
subcategory: ""
description: |-
  Upload inline SSL trust-store
---

# gigavuecore_upload_inline_ssl_trust_store Action

Upload inline SSL trust-store

## Example Usage

```terraform
action "gigavuecore_upload_inline_ssl_trust_store" "example" {
  config {
    cluster_id = "example"
    file = "example"
    file_source = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `file` (String, optional) - The contents of the file. Mutually exclusive with 'fileSource'
* `file_source` (Dynamic, optional) - Remote file source or destination
