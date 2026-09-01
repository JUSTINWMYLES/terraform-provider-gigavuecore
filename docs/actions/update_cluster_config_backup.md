---
page_title: "gigavuecore_update_cluster_config_backup Action - gigavuecore"
subcategory: ""
description: |-
  Update meta data of selected config backup snapshot for a given cluster
---

# gigavuecore_update_cluster_config_backup Action

Update meta data of selected config backup snapshot for a given cluster

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_update_cluster_config_backup" "example" {
  config {
    alias        = "example"
    backup_id    = "example"
    cluster_id   = "example"
    comment      = "example"
    do_not_purge = true
    tags         = ["example"]
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, optional) - User-entered alias for this snapshot
* `backup_id` (String, required) - Unique identifier of the cluster config snapshot for a given cluster
* `cluster_id` (String, required) - Cluster Id to restore cluster config for
* `comment` (String, optional) - User-entered comments for this snapshot
* `do_not_purge` (Boolean, optional) - Indicates whether this snapshot should not be auto-aged/purged
* `tags` (List of String, optional) - Tags associated with this file


