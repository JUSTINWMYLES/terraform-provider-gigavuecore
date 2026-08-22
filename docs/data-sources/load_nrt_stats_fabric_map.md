---
page_title: "gigavuecore_load_nrt_stats_fabric_map Data Source - gigavuecore"
subcategory: ""
description: |-
  Get NRT stats for the user-defined fabric map
---

# gigavuecore_load_nrt_stats_fabric_map Data Source

Get NRT stats for the user-defined fabric map

## Example Usage

```terraform
data "gigavuecore_load_nrt_stats_fabric_map" "example" {
  fabric_map_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `fabric_map_alias` (String, required) - alias of the fabric map

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, cluster_name, component_type, health_state, health_state_reasons, nrt_alias, reason, stats_data, traffic_health_state, traffic_health_state_reasons, update_time})), computed)

