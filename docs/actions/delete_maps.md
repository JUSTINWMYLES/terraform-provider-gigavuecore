---
page_title: "gigavuecore_delete_maps Action - gigavuecore"
subcategory: ""
description: |-
  Delete all maps
---

# gigavuecore_delete_maps Action

Delete all maps

## Example Usage

```terraform
action "gigavuecore_delete_maps" "example" {
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


