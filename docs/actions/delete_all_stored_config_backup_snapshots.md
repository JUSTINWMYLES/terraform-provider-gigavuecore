---
page_title: "gigavuecore_delete_all_stored_config_backup_snapshots Action - gigavuecore"
subcategory: ""
description: |-
  Delete all config backup snapshots associated with a cluster
---

# gigavuecore_delete_all_stored_config_backup_snapshots Action

Delete all config backup snapshots associated with a cluster

## Example Usage

```terraform
action "gigavuecore_delete_all_stored_config_backup_snapshots" "example" {
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - ID of the target cluster


