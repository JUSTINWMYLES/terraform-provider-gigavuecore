---
page_title: "gigavuecore_load_mobility_sam_node Data Source - gigavuecore"
subcategory: ""
description: |-
  Load mobility intent sam node by alias
---

# gigavuecore_load_mobility_sam_node Data Source

Load mobility intent sam node by alias

## Example Usage

```terraform
data "gigavuecore_load_mobility_sam_node" "example" {
  samnode_alias  = "example"
  solution_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `samnode_alias` (String, required) - Alias of the sam node
* `solution_alias` (String, required) - Alias of the mobility solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias` (String, computed) - Alias of the SAM Exporter node
* `app_profile_config` (Attributes, computed) (see [below for nested schema](#nestedatt--app_profile_config))
* `config_status` (String, computed)
* `config_status_reasons` (String, computed)
* `control_plane_setting` (Attributes, computed) (see [below for nested schema](#nestedatt--control_plane_setting))
* `deployed` (Boolean, computed) - True when the SAM node is attempted for deployment
* `deployment_details` (Attributes List, computed) (see [below for nested schema](#nestedatt--deployment_details))
* `engine_meta_data_cache_configs` (Attributes List, computed) (see [below for nested schema](#nestedatt--engine_meta_data_cache_configs))
* `engine_source_mappings` (Attributes List, computed) (see [below for nested schema](#nestedatt--engine_source_mappings))
* `exporter_config` (Attributes, computed) (see [below for nested schema](#nestedatt--exporter_config))
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `ip_interface_alias` (String, computed) - Alias of the ip interface used by SAM Exporter Node
* `location` (Attributes, computed) - Location of the engine port (see [below for nested schema](#nestedatt--location))
* `node_override_network_ports` (List of String, computed) - Network ports for the SAM Exporter node. Its a list of ports of format cluster:port
* `node_type` (String, computed) - Type of the SAM Exporter Node
* `param_configs` (Attributes List, computed) (see [below for nested schema](#nestedatt--param_configs))
* `smaf_details` (Attributes List, computed) (see [below for nested schema](#nestedatt--smaf_details))
* `tags` (Attributes List, computed) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--tags))
* `traffic_sources` (Attributes List, computed) - List of all traffic sources for the SAM Exporter node (see [below for nested schema](#nestedatt--traffic_sources))

<a id="nestedatt--app_profile_config"></a>
### Nested Schema for `app_profile_config`

Read-Only:

* `application_id` (Boolean)
* `applications` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--counter))
* `flow` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--gtpu))
* `inner_ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--inner_ipv4))
* `inner_ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--inner_ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--app_profile_config--transport))
<a id="nestedatt--app_profile_config--applications"></a>
### Nested Schema for `app_profile_config.applications`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_profile_config--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name
<a id="nestedatt--app_profile_config--applications--attributes"></a>
### Nested Schema for `app_profile_config.applications.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value
<a id="nestedatt--app_profile_config--counter"></a>
### Nested Schema for `app_profile_config.counter`

Read-Only:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)
<a id="nestedatt--app_profile_config--flow"></a>
### Nested Schema for `app_profile_config.flow`

Read-Only:

* `end_reason` (Boolean)
<a id="nestedatt--app_profile_config--gtpu"></a>
### Nested Schema for `app_profile_config.gtpu`

Read-Only:

* `qfi` (Boolean)
* `teid` (Boolean)
<a id="nestedatt--app_profile_config--inner_ipv4"></a>
### Nested Schema for `app_profile_config.inner_ipv4`

Read-Only:

* `destination` (Boolean)
* `protocol` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_profile_config--inner_ipv6"></a>
### Nested Schema for `app_profile_config.inner_ipv6`

Read-Only:

* `destination` (Boolean)
* `next_header` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_profile_config--outer_ipv4"></a>
### Nested Schema for `app_profile_config.outer_ipv4`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_profile_config--outer_ipv6"></a>
### Nested Schema for `app_profile_config.outer_ipv6`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_profile_config--timestamp"></a>
### Nested Schema for `app_profile_config.timestamp`

Read-Only:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
<a id="nestedatt--app_profile_config--transport"></a>
### Nested Schema for `app_profile_config.transport`

Read-Only:

* `dst_port` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--control_plane_setting"></a>
### Nested Schema for `control_plane_setting`

Read-Only:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--control_plane_setting--event_enable))
* `trigger` (String)
<a id="nestedatt--control_plane_setting--event_enable"></a>
### Nested Schema for `control_plane_setting.event_enable`

Read-Only:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--deployment_details"></a>
### Nested Schema for `deployment_details`

Read-Only:

* `engine_port` (String)
* `sam_node_alias` (String)
<a id="nestedatt--engine_meta_data_cache_configs"></a>
### Nested Schema for `engine_meta_data_cache_configs`

Read-Only:

* `engine_port` (String)
* `event` (String)
* `flow_behavior` (String)
* `flows_size` (Number)
* `idle_timeout` (Number)
* `match` (Attributes) (see [below for nested schema](#nestedatt--engine_meta_data_cache_configs--match))
* `observation_domain_id` (Number)
<a id="nestedatt--engine_meta_data_cache_configs--match"></a>
### Nested Schema for `engine_meta_data_cache_configs.match`

Read-Only:

* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--engine_meta_data_cache_configs--match--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--engine_meta_data_cache_configs--match--ipv6))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--engine_meta_data_cache_configs--match--transport))
<a id="nestedatt--engine_meta_data_cache_configs--match--ipv4"></a>
### Nested Schema for `engine_meta_data_cache_configs.match.ipv4`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--engine_meta_data_cache_configs--match--ipv4--destination))
* `protocol` (Boolean)
* `source` (Attributes) (see [below for nested schema](#nestedatt--engine_meta_data_cache_configs--match--ipv4--source))
<a id="nestedatt--engine_meta_data_cache_configs--match--ipv4--destination"></a>
### Nested Schema for `engine_meta_data_cache_configs.match.ipv4.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--engine_meta_data_cache_configs--match--ipv4--source"></a>
### Nested Schema for `engine_meta_data_cache_configs.match.ipv4.source`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--engine_meta_data_cache_configs--match--ipv6"></a>
### Nested Schema for `engine_meta_data_cache_configs.match.ipv6`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--engine_meta_data_cache_configs--match--ipv6--destination))
* `next_header` (Boolean)
* `source` (Attributes) (see [below for nested schema](#nestedatt--engine_meta_data_cache_configs--match--ipv6--source))
<a id="nestedatt--engine_meta_data_cache_configs--match--ipv6--destination"></a>
### Nested Schema for `engine_meta_data_cache_configs.match.ipv6.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--engine_meta_data_cache_configs--match--ipv6--source"></a>
### Nested Schema for `engine_meta_data_cache_configs.match.ipv6.source`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--engine_meta_data_cache_configs--match--transport"></a>
### Nested Schema for `engine_meta_data_cache_configs.match.transport`

Read-Only:

* `dst_port` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--engine_source_mappings"></a>
### Nested Schema for `engine_source_mappings`

Read-Only:

* `engine_port` (String)
* `network_source` (String)
<a id="nestedatt--exporter_config"></a>
### Nested Schema for `exporter_config`

Read-Only:

* `active_timeout` (String)
* `inactive_timeout` (Number)
* `record_type` (String)
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--location"></a>
### Nested Schema for `location`

Read-Only:

* `cluster_id` (String)
* `engine_ports` (List of String)
<a id="nestedatt--param_configs"></a>
### Nested Schema for `param_configs`

Read-Only:

* `engine_port` (String)
* `resource_metadata` (Number)
<a id="nestedatt--smaf_details"></a>
### Nested Schema for `smaf_details`

Read-Only:

* `control_application` (Attributes) (see [below for nested schema](#nestedatt--smaf_details--control_application))
* `management_address` (String)
* `user_application` (Attributes) (see [below for nested schema](#nestedatt--smaf_details--user_application))
<a id="nestedatt--smaf_details--control_application"></a>
### Nested Schema for `smaf_details.control_application`

Read-Only:

* `interface_address` (String)
* `port` (Number)
* `protocol` (String)
<a id="nestedatt--smaf_details--user_application"></a>
### Nested Schema for `smaf_details.user_application`

Read-Only:

* `interface_address` (String)
* `port` (Number)
* `protocol` (String)
<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_sources"></a>
### Nested Schema for `traffic_sources`

Read-Only:

* `comment` (String)
* `expand_port_identifier` (Boolean) - if enabled two rules added will be created with source port and destination port along with the other identifiers.
* `group_interfaces` (Boolean) - Enable if the user needs to combine rules in tunnelIdentifiers into single ingress map
* `ip4_frag_rule_type` (String) - if provided, the ip frag rule of given type would be appended to the rules.
* `network_function_interfaces` (Attributes List) (see [below for nested schema](#nestedatt--traffic_sources--network_function_interfaces))
* `network_function_name` (String)
* `network_function_type` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--traffic_sources--tags))
<a id="nestedatt--traffic_sources--network_function_interfaces"></a>
### Nested Schema for `traffic_sources.network_function_interfaces`

Read-Only:

* `comment` (String)
* `interface_type` (String)
* `source_override_dst_ports` (List of String) - Tool ports which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `source_override_network_ports` (List of String) - Network ports for the network element. Its a list of ports of format cluster:port
* `tunnel_identifiers` (List of Dynamic)
<a id="nestedatt--traffic_sources--tags"></a>
### Nested Schema for `traffic_sources.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag

