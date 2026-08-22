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
    cluster_specs = "example"
    image_server = "example"
    skip_not_reachable_devices = true
    task_name = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `cluster_specs` (List(Dynamic), required) - Cluster wise gs image spec
* `image_server` (String, required) - alias of an image file server. it has to reference one of the existing image file server profiles
* `skip_not_reachable_devices` (Bool, optional) - indicates whether the selected GigaSMART cards which belong to not-reachable nodes(if any) should be skipped and continue with image upgrade
* `task_name` (String, required) - user provided task name for the image upgrade
