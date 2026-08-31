---
page_title: "gigavuecore_delete_ssl_decryption_key_mapping Action - gigavuecore"
subcategory: ""
description: |-
  Delete an Endpoint-to-Key mapping
---

# gigavuecore_delete_ssl_decryption_key_mapping Action

Delete an Endpoint-to-Key mapping

## Example Usage

```terraform
action "gigavuecore_delete_ssl_decryption_key_mapping" "example" {
  config {
    alias          = "example"
    cluster_id     = "example"
    endpoint_alias = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target KeyMap
* `cluster_id` (String, required) - Target Cluster ID
* `endpoint_alias` (String, required) - alias of the target SSL Decryption Endpoint


