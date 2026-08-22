---
page_title: "gigavuecore_get_policy_nrt_stats Data Source - gigavuecore"
subcategory: ""
description: |-
  Get near real-time statistics for a traffic flow
---

# gigavuecore_get_policy_nrt_stats Data Source

Get near real-time statistics for a traffic flow

## Example Usage

```terraform
data "gigavuecore_get_policy_nrt_stats" "example" {
  alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - Traffic flow alias

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({alias, cluster_name, component_type, health_state, health_state_reasons, nrt_alias, reason, stats_data, traffic_health_state, traffic_health_state_reasons, update_time})), computed)

