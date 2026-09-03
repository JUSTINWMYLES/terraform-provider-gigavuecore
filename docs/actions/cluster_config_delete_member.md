---
page_title: "gigavuecore_cluster_config_delete_member Action - gigavuecore"
subcategory: ""
description: |-
  removes the specified member from the cluster
---

# gigavuecore_cluster_config_delete_member Action

removes the specified member from the cluster

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_cluster_config_delete_member" "example" {
  config {
    box_id     = "example"
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, required) - box id of the device that needs to be removed from the cluster
* `cluster_id` (String, required) - cluster id to which the device belongs


