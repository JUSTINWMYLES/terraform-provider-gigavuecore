---
page_title: "gigavuecore_replace_profile_decrypt_port_map Action - gigavuecore"
subcategory: ""
description: |-
  Replace port map in the profile
---

# gigavuecore_replace_profile_decrypt_port_map Action

Replace port map in the profile

-> **Note:** This action requires Terraform 1.14 or later. Standalone actions are invoked with `terraform apply -invoke=action.<type>.<name>` (or attached to a resource lifecycle `action_trigger`); a plain `terraform apply` does not invoke a standalone action block.

## Example Usage

```terraform
action "gigavuecore_replace_profile_decrypt_port_map" "example" {
  config {
    alias      = "example"
    cluster_id = "example"
    port_maps = [{
      in_port  = 1
      out_port = 1
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

