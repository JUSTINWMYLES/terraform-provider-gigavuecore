---
page_title: "gigavuecore_delete_map_groups Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Map Groups
---

# gigavuecore_delete_map_groups Action

Delete all Map Groups

## Example Usage

```terraform
action "gigavuecore_delete_map_groups" "example" {
  config {
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
