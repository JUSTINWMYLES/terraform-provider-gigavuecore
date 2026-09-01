---
page_title: "gigavuecore_remove_clusters_from_domain Action - gigavuecore"
subcategory: ""
description: |-
  Removes all managed clusters and standalone nodes from FM management
---

# gigavuecore_remove_clusters_from_domain Action

Removes all managed clusters and standalone nodes from FM management

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_remove_clusters_from_domain" "example" {
  config {
    cluster_id = "example"
  }
}
```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - if provided, only requested cluster is removed. For Standalone nodes, removes that one node (identified by its nodeId). For clustered nodes, removes entire cluster


