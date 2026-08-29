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
    alias      = "example"
    cluster_id = "example"
    key_maps = [{
      hostname = "example"
      key      = "example"
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
* `key_maps` (Attributes List, required) (see [below for nested schema](#nestedatt--key_maps))

<a id="nestedatt--key_maps"></a>
### Nested Schema for `key_maps`

Required:

* `hostname` (String) - hostname or IP address
* `key` (String) - SSL key alias
Optional:

* `rule_id` (Number)

