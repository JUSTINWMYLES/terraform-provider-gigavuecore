---
page_title: "gigavuecore_delete_stored_cluster_config_backup_snapshot Action - gigavuecore"
subcategory: ""
description: |-
  Delete a labeled config backup snapshot
---

# gigavuecore_delete_stored_cluster_config_backup_snapshot Action

Delete a labeled config backup snapshot

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_stored_cluster_config_backup_snapshot" "example" {
  config {
    backup_id  = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `backup_id` (String, required) - id of the config backup snapshot to delete
* `cluster_id` (String, required) - ID of the target cluster


