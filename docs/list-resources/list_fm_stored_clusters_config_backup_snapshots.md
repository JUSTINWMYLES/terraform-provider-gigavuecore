---
page_title: "gigavuecore_list_fm_stored_clusters_config_backup_snapshots List Resource - gigavuecore"
subcategory: ""
description: |-
  Lists config backup snapshots for multiple clusters
---

# gigavuecore_list_fm_stored_clusters_config_backup_snapshots List Resource

Lists config backup snapshots for multiple clusters

## Example Usage

```terraform
list "gigavuecore_list_fm_stored_clusters_config_backup_snapshots" "example" {
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

* `cluster_id` (String, optional) - if provided, only backups associated with this cluster will be returned


### Identity Attributes

The following identity attributes are exported for each matching result:

* `cluster_id` (String, computed)
* `backup_id` (String, computed)


