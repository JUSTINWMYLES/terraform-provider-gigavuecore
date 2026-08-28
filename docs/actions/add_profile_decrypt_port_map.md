---
page_title: "gigavuecore_add_profile_decrypt_port_map Action - gigavuecore"
subcategory: ""
description: |-
  Add a new port map entry to the profile
---

# gigavuecore_add_profile_decrypt_port_map Action

Add a new port map entry to the profile

## Example Usage

```terraform
action "gigavuecore_add_profile_decrypt_port_map" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    in_port    = 1
    out_port   = 1
    rule_id    = 1
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `in_port` (Number, required) - ingress port for decryption port map
* `out_port` (Number, required) - egress port for decryption port map
* `rule_id` (Number, optional)


