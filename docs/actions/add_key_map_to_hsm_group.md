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
    alias      = "example"
    cluster_id = "example"
    hsm_key_maps = [{
      address    = "example"
      cluster_id = "example"
      key_name   = "example"
      key_token  = "example"
      port       = 0
      rfs_match  = "example"
      rule_id    = "example"
    }]
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - alias of target HSM Group
* `cluster_id` (String, required) - Target cluster ID.
* `hsm_key_maps` (Attributes List, optional) (see [below for nested schema](#nestedatt--hsm_key_maps))

<a id="nestedatt--hsm_key_maps"></a>
### Nested Schema for `hsm_key_maps`

Required:

* `address` (String) - IPv4 address of the SSL endpoint (server)

Optional:

* `cluster_id` (String) - id of the defining cluster
* `key_name` (String) - key-name
* `key_token` (String) - key-token
* `port` (Number) - Port of the SSL endpoint (server). Value of 0 indicates any port
* `rfs_match` (String) - Is there a matching RFS key, yes or no
* `rule_id` (String) - Rule Id

