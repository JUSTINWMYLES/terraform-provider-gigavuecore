---
page_title: "gigavuecore_get_overlap_components Data Source - gigavuecore"
subcategory: ""
description: |-
  Retrieve overlap components for a cluster and vport
---

# gigavuecore_get_overlap_components Data Source

Retrieve overlap components for a cluster and vport

## Example Usage

```terraform
data "gigavuecore_get_overlap_components" "example" {
  cluster_id = null
  vport_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - Target Cluster ID
* `vport_alias` (String, required) - VPort alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `cluster_name` (String, computed)
* `comment` (String, computed)
* `components` (List(Object({alias, health_state, health_state_reasons, source_alias, traffic_health_state, traffic_health_state_reasons, type})), computed)
* `map_group_alias` (String, computed)

