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
    alias      = "example"
    cluster_id = "example"
    port_maps = [{
      in_port  = 0
      out_port = 0
      rule_id  = 0
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of the inline SSL profile
* `cluster_id` (String, required) - Target Cluster ID
* `port_maps` (Attributes List, required) (see [below for nested schema](#nestedatt--port_maps))

<a id="nestedatt--port_maps"></a>
### Nested Schema for `port_maps`

Required:

* `in_port` (Number) - ingress port for decryption port map
* `out_port` (Number) - egress port for decryption port map

Optional:

* `rule_id` (Number)

