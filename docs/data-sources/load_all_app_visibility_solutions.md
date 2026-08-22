---
page_title: "gigavuecore_load_all_app_visibility_solutions Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Application Intelligence Solutions deployed on V Series
---

# gigavuecore_load_all_app_visibility_solutions Data Source

Load all Application Intelligence Solutions deployed on V Series

## Example Usage

```terraform
data "gigavuecore_load_all_app_visibility_solutions" "example" {
}
```

## Schema

### Arguments

The following arguments are supported:


### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({app_env_id, app_exporter_config, app_filter_config, config_status, config_status_reasons, conn_id, dedup, dynamic_scale_unit, env_id, health_state, health_state_reasons, ingress_traffic_configs, monitor_solution_config, scale_unit, solution_alias, solution_created_timestamp, solution_desc, solution_modified_timestamp, traffic_policy_graph_alias})), computed)

