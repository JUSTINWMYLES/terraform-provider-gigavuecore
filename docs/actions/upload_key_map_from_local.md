---
page_title: "gigavuecore_upload_key_map_from_local Action - gigavuecore"
subcategory: ""
description: |-
  Upload Key Map file from local
---

# gigavuecore_upload_key_map_from_local Action

Upload Key Map file from local

## Example Usage

```terraform
action "gigavuecore_upload_key_map_from_local" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    file = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.
* `file` (String, required) - key map file to upload
