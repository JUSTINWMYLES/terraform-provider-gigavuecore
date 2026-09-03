---
page_title: "gigavuecore_delete_all_apps_exporter_group Action - gigavuecore"
subcategory: ""
description: |-
  Delete all Apps Exporter Group
---

# gigavuecore_delete_all_apps_exporter_group Action

Delete all Apps Exporter Group

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_apps_exporter_group" "example" {
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


