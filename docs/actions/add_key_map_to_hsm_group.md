---
page_title: "gigavuecore_add_key_map_to_hsm_group Action - gigavuecore"
subcategory: ""
description: |-
  Add keymap to HSM Group
---

# gigavuecore_add_key_map_to_hsm_group Action

Add keymap to HSM Group

## Example Usage

```terraform
action "gigavuecore_add_key_map_to_hsm_group" "example" {
  config {
    alias        = "example"
    cluster_id   = "example"
    hsm_key_maps = null
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.
* `hsm_key_maps` (List of Dynamic, optional)


