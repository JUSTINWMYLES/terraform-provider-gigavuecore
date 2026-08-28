---
page_title: "gigavuecore_update_inline_ssl_config Action - gigavuecore"
subcategory: ""
description: |-
  Update inline SSL global configuration
---

# gigavuecore_update_inline_ssl_config Action

Update inline SSL global configuration

## Example Usage

```terraform
action "gigavuecore_update_inline_ssl_config" "example" {
  config {
    body_cluster_id = "example"
    caching         = null
    cluster_id      = "example"
    dhe_ciphersuit  = "example"
    monitor         = null
    resumption      = null
    ssl_versions    = null
    start_tls       = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `body_cluster_id` (String, optional) - id of the defining cluster
* `caching` (Dynamic, optional)
* `cluster_id` (String, required) - Target Cluster ID
* `dhe_ciphersuit` (String, optional)
* `monitor` (Dynamic, optional)
* `resumption` (Dynamic, optional)
* `ssl_versions` (Dynamic, optional)
* `start_tls` (Dynamic, optional)


