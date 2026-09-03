---
page_title: "gigavuecore_upgrade_cluster_nodes_image Action - gigavuecore"
subcategory: ""
description: |-
  upgrade cluster nodes image
---

# gigavuecore_upgrade_cluster_nodes_image Action

upgrade cluster nodes image

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_upgrade_cluster_nodes_image" "example" {
  config {
    async       = true
    cluster_ids = ["example"]
    image_file_specs = [{
      device_model = "HD8"
      file_path    = "example"
      file_type    = "image"
      target_slot  = "2"
      update_uboot = true
    }]
    image_server               = "example"
    node_ids                   = ["example"]
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
* `image_file_specs` (Attributes List, required) - lists image files to use (see [below for nested schema](#nestedatt--image_file_specs))
* `image_server` (String, required) - alias of an image file server. has to reference one of the existing image file server profiles
* `node_ids` (List of String, optional) - ids of individual nodes to upgrade image for
* `reboot` (Boolean, required) - indicates whether nodes should reboot after image upgrade
* `skip_not_reachable_devices` (Boolean, optional) - indicates whether not-reachable nodes(if any) should be skipped and continue with image upgrade

<a id="nestedatt--image_file_specs"></a>
### Nested Schema for `image_file_specs`

Required:

* `device_model` (String) - Gigamon physical device models
* `file_path` (String) - image file path on the image server

Optional:

* `file_type` (String) - image file type. 'web' and 'gigasmart' are only applicable for G-series devices
* `target_slot` (String) - only applicable for 'gigasmart' images for 2404 G-series devices. Indicates the target GS card to upgrade
* `update_uboot` (Boolean) - only applicable for H-series nodes. indicates whether the uboot is to be upgraded as part of the image upgrade

