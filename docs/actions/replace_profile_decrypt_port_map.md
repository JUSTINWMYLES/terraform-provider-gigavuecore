---
page_title: "gigavuecore_replace_profile_decrypt_port_map Action - gigavuecore"
subcategory: ""
description: |-
  Replace port map in the profile
---

# gigavuecore_replace_profile_decrypt_port_map Action

Replace port map in the profile

## Example Usage

```terraform
action "gigavuecore_replace_profile_decrypt_port_map" "example" {
  config {
    alias = "example"
    cluster_id = "example"
    port_maps = "example"
  }
}

```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `port_maps` (List(Dynamic), required)
