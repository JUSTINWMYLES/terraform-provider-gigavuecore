---
page_title: "gigavuecore_delete_gs_dump_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete Gigasmart dump File
---

# gigavuecore_delete_gs_dump_file Action

Delete Gigasmart dump File

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_gs_dump_file" "example" {
  config {
    cluster_id = "example"
    filename   = "example"
    hostname   = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster Id
* `filename` (String, required) - Gigasmart dump filename to delete
* `hostname` (String, required) - Target hostname


