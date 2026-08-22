---
page_title: "gigavuecore_cluster_config_add_member Action - gigavuecore"
subcategory: ""
description: |-
  adds the device to the specified cluster id
---

# gigavuecore_cluster_config_add_member Action

adds the device to the specified cluster id

## Example Usage

```terraform
action "gigavuecore_cluster_config_add_member" "example" {
  config {
    bulk = true
    cluster_id = "example"
    cluster_member_specs = null
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `bulk` (Bool, optional) - if more than one member needs to be added then set it to true, otherwise false
* `cluster_id` (String, required) - cluster id to which the member(s) to be added
* `cluster_member_specs` (List(Dynamic), optional)
