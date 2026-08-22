---
page_title: "gigavuecore_delete_all_tunnel_endpoint Action - gigavuecore"
subcategory: ""
description: |-
  Delete all tunnel endpoints
---

# gigavuecore_delete_all_tunnel_endpoint Action

Delete all tunnel endpoints

## Example Usage

```terraform
action "gigavuecore_delete_all_tunnel_endpoint" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
