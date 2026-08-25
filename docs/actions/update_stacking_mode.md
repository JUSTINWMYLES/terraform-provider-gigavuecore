---
page_title: "gigavuecore_update_stacking_mode Action - gigavuecore"
subcategory: ""
description: |-
  updating the stacking mode on the cluster
---

# gigavuecore_update_stacking_mode Action

updating the stacking mode on the cluster

## Example Usage

```terraform
action "gigavuecore_update_stacking_mode" "example" {
  config {
    cluster_id    = "example"
    stacking_mode = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Cluster Id of the cluster
* `stacking_mode` (String, required) - stacking mode to be updated into cluster


