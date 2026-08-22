---
page_title: "gigavuecore_load_map_alias_chain Data Source - gigavuecore"
subcategory: ""
description: |-
  Load aliased map chain by source ports or mapChain ID
---

# gigavuecore_load_map_alias_chain Data Source

Load aliased map chain by source ports or mapChain ID

## Example Usage

```terraform
data "gigavuecore_load_map_alias_chain" "example" {
  cluster_id = null
  id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - id of the defining cluster
* `id` (String, required) - srcPortsAsId or mapChainId for which the map chain ordered by map alias is to be displayed

### Attributes

In addition to all arguments above, the following attributes are exported:

* `collector` (String, computed) - alias of the chain collector map. must reference a map of matching type and  'collector' subtype
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (List(Object({message, severity, traffic_health_state_computation_type})), computed)
* `map_chain_id` (String, computed) - mapChain ID - replaces srcPortsAsId - should be used anywhere that requires srcPortsAsId
* `ordered_map_aliases` (Set(String), computed) - aliases of the maps in the chain
* `src_ports` (List(String), computed) - network ports for the maps chain
* `src_ports_as_id` (String, computed) - (Deprecated - use mapChainId instead) mapChain ID

