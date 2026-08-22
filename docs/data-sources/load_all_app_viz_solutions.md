---
page_title: "gigavuecore_load_all_app_viz_solutions Data Source - gigavuecore"
subcategory: ""
description: |-
  Load all Apps Visibility Solutions
---

# gigavuecore_load_all_app_viz_solutions Data Source

Load all Apps Visibility Solutions

## Example Usage

```terraform
data "gigavuecore_load_all_app_viz_solutions" "example" {
  cluster_id = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `cluster_id` (String, required) - clusterId

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (List(Object({app_export_config, app_filter_config, associated_monitor_solution_alias, cluster_id, config_status, delete_monitor_sol, egress_map_aliases_to_delete, exporter_aliases_to_delete, health_state, health_state_reasons, ingress_map_aliases_to_delete, ingress_traffic_configs, monitor_solution_config, solution_alias, solution_desc, solution_status, solution_type})), computed)

