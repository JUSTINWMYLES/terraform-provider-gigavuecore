---
page_title: "gigavuecore_delete_all_apps_exporter Action - gigavuecore"
subcategory: ""
description: |-
  new in H 5.8
---

# gigavuecore_delete_all_apps_exporter Action

new in H 5.8

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_delete_all_apps_exporter" "example" {
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID


