---
page_title: "gigavuecore_upgrade_cluster_gs_card_image Action - gigavuecore"
subcategory: ""
description: |-
  upgrade the GS cards for selected clusters, it is an async only
---

# gigavuecore_upgrade_cluster_gs_card_image Action

upgrade the GS cards for selected clusters, it is an async only

## Example Usage

```terraform
action "gigavuecore_upgrade_cluster_gs_card_image" "example" {
  config {
    cluster_specs = [{
      cluster_id = "example"
      image_file_spec = [{
        box_id       = "example"
        file_path    = "example"
        product_code = "example"
        target_slots = [ "example" ]
      }]
    }]
    image_server               = "example"
    skip_not_reachable_devices = true
    task_name                  = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_specs` (Attributes List, required) - Cluster wise gs image spec (see [below for nested schema](#nestedatt--cluster_specs))
* `image_server` (String, required) - alias of an image file server. it has to reference one of the existing image file server profiles
* `skip_not_reachable_devices` (Boolean, optional) - indicates whether the selected GigaSMART cards which belong to not-reachable nodes(if any) should be skipped and continue with image upgrade
* `task_name` (String, required) - user provided task name for the image upgrade

<a id="nestedatt--cluster_specs"></a>
### Nested Schema for `cluster_specs`

Required:

* `cluster_id` (String) - clusterId
* `image_file_spec` (Attributes List) - list of upgrade spec for the cluster (see [below for nested schema](#nestedatt--cluster_specs--image_file_spec))

<a id="nestedatt--cluster_specs--image_file_spec"></a>
### Nested Schema for `cluster_specs.image_file_spec`

Required:

* `file_path` (String) - path of the image file
* `product_code` (String) - product code of gs card to be upgraded

Optional:

* `box_id` (String) - boxId of the gs card to be upgraded
* `target_slots` (List of String) - slotIds of the box to be to upgraded

