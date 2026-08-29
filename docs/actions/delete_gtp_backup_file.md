---
page_title: "gigavuecore_delete_gtp_backup_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete a Gtp backup file
---

# gigavuecore_delete_gtp_backup_file Action

Delete a Gtp backup file

## Example Usage

```terraform
action "gigavuecore_delete_gtp_backup_file" "example" {
  config {
    cluster_id = "example"
    filename   = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `filename` (String, required) - filename of Gtp backup file


