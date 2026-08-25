---
page_title: "gigavuecore_upload_ssl_client_trust_store_file Action - gigavuecore"
subcategory: ""
description: |-
  Create/Replace the SSL Client trust-store from local file
---

# gigavuecore_upload_ssl_client_trust_store_file Action

Create/Replace the SSL Client trust-store from local file

## Example Usage

```terraform
action "gigavuecore_upload_ssl_client_trust_store_file" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    file       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Client Trust Store alias
* `cluster_id` (String, required) - Target Cluster ID
* `file` (String, required) - file to upload to device


