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
    clusters_config_specs = [{
      cluster_id = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `clusters_config_specs` (Attributes List, required) - list of Clusters config backup references (see [below for nested schema](#nestedatt--clusters_config_specs))

<a id="nestedatt--clusters_config_specs"></a>
### Nested Schema for `clusters_config_specs`

Required:

* `cluster_id` (String) - Cluster Id

