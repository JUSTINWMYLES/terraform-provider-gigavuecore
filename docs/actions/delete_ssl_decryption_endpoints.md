---
page_title: "gigavuecore_delete_ssl_decryption_endpoints Action - gigavuecore"
subcategory: ""
description: |-
  Delete SSL Decryption Endpoints
---

# gigavuecore_delete_ssl_decryption_endpoints Action

Delete SSL Decryption Endpoints

## Example Usage

```terraform
action "gigavuecore_delete_ssl_decryption_endpoints" "example" {
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


