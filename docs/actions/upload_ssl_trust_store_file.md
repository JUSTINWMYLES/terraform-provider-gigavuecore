---
page_title: "gigavuecore_upload_ssl_trust_store_file Action - gigavuecore"
subcategory: ""
description: |-
  Upload inline SSL trust-store from local file
---

# gigavuecore_upload_ssl_trust_store_file Action

Upload inline SSL trust-store from local file

## Example Usage

```terraform
action "gigavuecore_upload_ssl_trust_store_file" "example" {
  config {
    cluster_id = "example"
    file       = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `file` (String, required) - file to upload to device


