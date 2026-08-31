---
page_title: "gigavuecore_restore_cluater_config_backup_snapshot Action - gigavuecore"
subcategory: ""
description: |-
  Restore selected config backup snapshot for a given cluster
---

# gigavuecore_restore_cluater_config_backup_snapshot Action

Restore selected config backup snapshot for a given cluster

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_restore_cluater_config_backup_snapshot" "example" {
  config {
    activate   = true
    backup_id  = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `activate` (Boolean, optional) - indicates whether this restored config should also be made active upon restore
* `backup_id` (String, required) - Reference to cluster config backup to be restored
* `cluster_id` (String, required) - Cluster Id to restore cluster config for


