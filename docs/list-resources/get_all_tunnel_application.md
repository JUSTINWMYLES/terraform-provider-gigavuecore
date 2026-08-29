---
page_title: "gigavuecore_get_all_tunnel_application List Resource - gigavuecore"
subcategory: ""
description: |-
  Get all Tunnel Apps
---

# gigavuecore_get_all_tunnel_application List Resource

Get all Tunnel Apps

## Example Usage

```terraform
list "gigavuecore_get_all_tunnel_application" "example" {
  provider = gigavuecore
  limit    = 100
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


### Identity Attributes

The following identity attributes are exported for each matching result:

* `alias` (String, computed)


