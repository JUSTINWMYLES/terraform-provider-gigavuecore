---
page_title: "gigavuecore_replace_profile_key_maps Action - gigavuecore"
subcategory: ""
description: |-
  Replace key map entries in the profile
---

# gigavuecore_replace_profile_key_maps Action

Replace key map entries in the profile

## Example Usage

```terraform
action "gigavuecore_replace_profile_key_maps" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    key_maps = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `key_maps` (List(Dynamic), required)
