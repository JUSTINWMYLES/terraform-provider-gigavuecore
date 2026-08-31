---
page_title: "gigavuecore_convert_to_fabric_map Action - gigavuecore"
subcategory: ""
description: |-
  convert the afm maps in the cluster to fabric maps
---

# gigavuecore_convert_to_fabric_map Action

convert the afm maps in the cluster to fabric maps

## Example Usage

```terraform
action "gigavuecore_convert_to_fabric_map" "example" {
  config {
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, optional) - indicate the clusterId of afm maps will be converted, if not specified, convert afm maps of the entire system


