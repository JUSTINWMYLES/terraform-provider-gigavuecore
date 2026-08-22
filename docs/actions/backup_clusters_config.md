---
page_title: "gigavuecore_backup_clusters_config Action - gigavuecore"
subcategory: ""
description: |-
  Backup of multiple clusters can be requested
---

# gigavuecore_backup_clusters_config Action

Backup of multiple clusters can be requested

## Example Usage

```terraform
action "gigavuecore_backup_clusters_config" "example" {
  config {
    clusters_config_specs = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `clusters_config_specs` (List(Dynamic), required) - list of Clusters config backup references
