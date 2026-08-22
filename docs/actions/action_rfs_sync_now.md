---
page_title: "gigavuecore_action_rfs_sync_now Action - gigavuecore"
subcategory: ""
description: |-
  Force Sync Now
---

# gigavuecore_action_rfs_sync_now Action

Force Sync Now

## Example Usage

```terraform
action "gigavuecore_action_rfs_sync_now" "example" {
  config {
    alias = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.
