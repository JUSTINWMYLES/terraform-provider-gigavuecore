---
page_title: "gigavuecore_update_map_priority Action - gigavuecore"
subcategory: ""
description: |-
  Update Map priority within its Chain
---

# gigavuecore_update_map_priority Action

Update Map priority within its Chain

## Example Usage

```terraform
action "gigavuecore_update_map_priority" "example" {
  config {
    cluster_id      = "example"
    id              = "example"
    map_alias       = "example"
    map_chain_id    = "example"
    priority_type   = "highest"
    ref_map_alias   = "example"
    src_ports_as_id = "example"
  }
}

```
## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `id` (String, required) - srcPortsAsId or mapChainId for which the map order is to be redefined. Note that mapChainID works only for classic maps.
* `map_alias` (String, required) - alias of the map(either cluster map or fabric map) to update priority for
* `map_chain_id` (String, optional) - mapChain ID - replaces srcPortsAsId - should be used anywhere that requires srcPortsAsId
* `priority_type` (String, required) - map priority update type within its map chain
* `ref_map_alias` (String, optional) - when 'priorityType' is 'before' or 'after', this field specifies the reference map (either cluster map or fabric map) for the update
* `src_ports_as_id` (String, optional) - (Deprecated - use mapChainId instead) mapChain ID


