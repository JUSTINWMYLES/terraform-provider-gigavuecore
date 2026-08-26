---
page_title: "gigavuecore_upgrade_cluster_nodes_image Action - gigavuecore"
subcategory: ""
description: |-
  upgrade cluster nodes image
---

# gigavuecore_upgrade_cluster_nodes_image Action

upgrade cluster nodes image

## Example Usage

```terraform
action "gigavuecore_upgrade_cluster_nodes_image" "example" {
  config {
    async                      = true
    cluster_ids                = [ "example" ]
    image_file_specs           = "example"
    image_server               = "example"
    node_ids                   = [ "example" ]
    reboot                     = true
    skip_not_reachable_devices = true
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `async` (Boolean, optional) - if provided, the call returns immediately with the \[202 Accepted\] HTTP status code, and the image upgrade will complete in the background
* `cluster_ids` (List of String, required) - ids of clusters to upgrade image for. Every node in these clusters is upgraded
* `image_file_specs` (List of Dynamic, required) - lists image files to use
* `image_server` (String, required) - alias of an image file server. has to reference one of the existing image file server profiles
* `node_ids` (List of String, optional) - ids of individual nodes to upgrade image for
* `reboot` (Boolean, required) - indicates whether nodes should reboot after image upgrade
* `skip_not_reachable_devices` (Boolean, optional) - indicates whether not-reachable nodes(if any) should be skipped and continue with image upgrade


