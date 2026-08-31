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

### Attributes

In addition to all arguments above, the following attributes are exported:

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `app_env_id` (String) - ID of the Application Environment
* `app_exporter_config` (Attributes) (see [below for nested schema](#nestedatt--items--app_exporter_config))
* `app_filter_config` (Attributes) (see [below for nested schema](#nestedatt--items--app_filter_config))
* `config_status` (String) - configuration status
* `config_status_reasons` (String) - configuration status reasons
* `conn_id` (String) - ID of the Connection Domain
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--items--dedup))
* `dynamic_scale_unit` (Boolean) - When set as true, FM automatically sets the optimal scaleUnit
* `env_id` (String) - ID of the Environment
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--items--health_state_reasons))
* `ingress_traffic_configs` (Attributes List) (see [below for nested schema](#nestedatt--items--ingress_traffic_configs))
* `monitor_solution_config` (Attributes) (see [below for nested schema](#nestedatt--items--monitor_solution_config))
* `scale_unit` (Number) - Number of units of memory needed for each Application
* `solution_alias` (String) - Alias of the solution
* `solution_created_timestamp` (String)
* `solution_desc` (String) - Description of the solution
* `solution_modified_timestamp` (String)
* `traffic_policy_graph_alias` (String) - Alias of the Traffic Policy Graph

<a id="nestedatt--items--app_exporter_config"></a>
### Nested Schema for `items.app_exporter_config`

Read-Only:

* `app_instance_name` (String) - The name of the AppMetadata app instance
* `app_metadata_tiered_batch_id` (String) - ID of the appExporterConfig tiered batch
* `cache_config_alias` (String) - Alias of the cacheConfig
* `destination_configs` (Attributes List) (see [below for nested schema](#nestedatt--items--app_exporter_config--destination_configs))

<a id="nestedatt--items--app_exporter_config--destination_configs"></a>
### Nested Schema for `items.app_exporter_config.destination_configs`

Read-Only:

* `destination_name` (String) - Alias of the destination tool
* `exporter_alias` (String) - Alias of the exporter
* `iface` (String) - Interface Mapping for Egress Tunnel
* `template_name` (String) - Name of the App Profile Template Used

<a id="nestedatt--items--app_filter_config"></a>
### Nested Schema for `items.app_filter_config`

Read-Only:

* `app_filter_tiered_batch_id` (String) - ID of the appFilterConfig tiered batch
* `egress_traffic_configs` (Attributes List) (see [below for nested schema](#nestedatt--items--app_filter_config--egress_traffic_configs))
* `map_alias` (String) - Alias of the appFiltering map
* `sapf_profile_alias` (String) - Alias of the SAPF Profile

<a id="nestedatt--items--app_filter_config--egress_traffic_configs"></a>
### Nested Schema for `items.app_filter_config.egress_traffic_configs`

Read-Only:

* `drop_application_names` (Attributes List) (see [below for nested schema](#nestedatt--items--app_filter_config--egress_traffic_configs--drop_application_names))
* `pass_application_names` (Attributes List) (see [below for nested schema](#nestedatt--items--app_filter_config--egress_traffic_configs--pass_application_names))
* `priority` (Number)
* `tunnel_aliases` (List of String)

<a id="nestedatt--items--app_filter_config--egress_traffic_configs--drop_application_names"></a>
### Nested Schema for `items.app_filter_config.egress_traffic_configs.drop_application_names`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--items--app_filter_config--egress_traffic_configs--drop_application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name

<a id="nestedatt--items--app_filter_config--egress_traffic_configs--drop_application_names--attributes"></a>
### Nested Schema for `items.app_filter_config.egress_traffic_configs.drop_application_names.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value

<a id="nestedatt--items--app_filter_config--egress_traffic_configs--pass_application_names"></a>
### Nested Schema for `items.app_filter_config.egress_traffic_configs.pass_application_names`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--items--app_filter_config--egress_traffic_configs--pass_application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name

<a id="nestedatt--items--app_filter_config--egress_traffic_configs--pass_application_names--attributes"></a>
### Nested Schema for `items.app_filter_config.egress_traffic_configs.pass_application_names.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value

<a id="nestedatt--items--dedup"></a>
### Nested Schema for `items.dedup`

Read-Only:

* `app_config_id` (String)
* `app_instance_name` (String)
* `dedup_tiered_batch_id` (String)
* `enabled` (Boolean)
* `gs_params_name` (String)

<a id="nestedatt--items--health_state_reasons"></a>
### Nested Schema for `items.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

<a id="nestedatt--items--ingress_traffic_configs"></a>
### Nested Schema for `items.ingress_traffic_configs`

Read-Only:

* `source_selector_aliases` (List of String)
* `src_raw_end_point_interfaces` (List of String)
* `tunnel_aliases` (List of String)
* `tunnel_interface_mappings` (Attributes List) (see [below for nested schema](#nestedatt--items--ingress_traffic_configs--tunnel_interface_mappings))

<a id="nestedatt--items--ingress_traffic_configs--tunnel_interface_mappings"></a>
### Nested Schema for `items.ingress_traffic_configs.tunnel_interface_mappings`

Read-Only:

* `iface` (String)
* `tunnel_alias` (String)

<a id="nestedatt--items--monitor_solution_config"></a>
### Nested Schema for `items.monitor_solution_config`

Read-Only:

* `app_instance_config_id` (String)
* `app_instance_name` (String)
* `app_viz_tiered_batch_id` (String) - ID of the monitoringSolutionConfig tiered batch
* `destination_config` (Attributes) (see [below for nested schema](#nestedatt--items--monitor_solution_config--destination_config))
* `gs_param_config_id` (String) - Gs param name on which the solution is applied
* `mgmt_interface` (String) - Pass this value if you would like to use the management interface of the Node to export App Monitoring metadata to FM
* `user_defined_application_profile` (Dynamic)
* `user_defined_applications` (List of List of String)

<a id="nestedatt--items--monitor_solution_config--destination_config"></a>
### Nested Schema for `items.monitor_solution_config.destination_config`

Read-Only:

* `destination_name` (String)
* `exporter_alias` (String)
* `iface` (String) - Interface Mapping for Egress Tunnel

