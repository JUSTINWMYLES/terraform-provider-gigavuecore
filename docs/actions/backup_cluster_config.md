---
page_title: "gigavuecore_backup_cluster_config Action - gigavuecore"
subcategory: ""
description: |-
  Backup of cluster can be requested
---

# gigavuecore_backup_cluster_config Action

Backup of cluster can be requested

## Example Usage

```terraform
action "gigavuecore_backup_cluster_config" "example" {
  config {
    body_cluster_id = "example"
    cluster_id = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `body_cluster_id` (String, required) - Cluster Id
* `cluster_id` (String, required) - ID of the target device
