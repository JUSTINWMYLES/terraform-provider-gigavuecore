---
page_title: "gigavuecore_add_ssl_decryption_key_mapping Action - gigavuecore"
subcategory: ""
description: |-
  Add an Endpoint-to-Key mapping
---

# gigavuecore_add_ssl_decryption_key_mapping Action

Add an Endpoint-to-Key mapping

## Example Usage

```terraform
action "gigavuecore_add_ssl_decryption_key_mapping" "example" {
  config {
    alias          = "example"
    cluster_id     = "example"
    endpoint_alias = "example"
    key_alias      = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the target KeyMap
* `cluster_id` (String, required) - Target Cluster ID
* `endpoint_alias` (String, required) - Alias of referenced SSL Endpoint
* `key_alias` (String, required) - Alias of referenced Decryption Key associated with this endpoint


