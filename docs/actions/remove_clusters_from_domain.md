---
page_title: "gigavuecore_remove_clusters_from_domain Action - gigavuecore"
subcategory: ""
description: |-
  Removes all managed clusters and standalone nodes from FM management
---

# gigavuecore_remove_clusters_from_domain Action

Removes all managed clusters and standalone nodes from FM management

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
