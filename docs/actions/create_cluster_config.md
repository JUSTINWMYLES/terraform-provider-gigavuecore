---
page_title: "gigavuecore_create_cluster_config Action - gigavuecore"
subcategory: ""
description: |-
  creates the cluster with the given device members list and cluster parameters
---

# gigavuecore_create_cluster_config Action

creates the cluster with the given device members list and cluster parameters

## Example Usage

```terraform
action "gigavuecore_create_cluster_config" "example" {
  config {
    cluster_member_specs = null
    cluster_params = null
    cluster_type = "example"
    seed_node_mgmt_address = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_member_specs` (List(Dynamic), optional)
* `cluster_params` (Dynamic, optional) - provides clustering parameters
* `cluster_type` (String, optional) - Type of the cluster eg. oob. It cannot be null
* `seed_node_mgmt_address` (String, optional) - seed node management address
