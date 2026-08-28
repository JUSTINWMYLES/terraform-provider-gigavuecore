---
page_title: "gigavuecore_delete_profile_decrypt_port_map Action - gigavuecore"
subcategory: ""
description: |-
  Delete all port map entries from the profile
---

# gigavuecore_delete_profile_decrypt_port_map Action

Delete all port map entries from the profile

## Example Usage

```terraform
action "gigavuecore_delete_profile_decrypt_port_map" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID


