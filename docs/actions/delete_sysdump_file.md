---
page_title: "gigavuecore_delete_sysdump_file Action - gigavuecore"
subcategory: ""
description: |-
  Delete Sysdump File
---

# gigavuecore_delete_sysdump_file Action

Delete Sysdump File

## Example Usage

```terraform
action "gigavuecore_delete_sysdump_file" "example" {
  config {
    box_id = "example"
    cluster_id = "example"
    filename = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `box_id` (String, optional) - Box ID range from 1 to 64(inclusive)
* `cluster_id` (String, required) - Target Cluster ID
* `filename` (String, required) - Sysdump filename to delete
