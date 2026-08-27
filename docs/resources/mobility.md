---
page_title: "gigavuecore_mobility Resource - gigavuecore"
subcategory: ""
description: |-
  Load mobility solutions by alias
---

# gigavuecore_mobility Resource

Load mobility solutions by alias

## Example Usage

```terraform
resource "gigavuecore_mobility" "example" {
  deployed         = null
  health_state     = null
  site_name        = null
  site_tag         = null
  sites            = []
  solution_alias   = null
  solution_type    = null
  tags             = []
  traffic_policies = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `deployed` (Boolean, optional) - Mobility solution with gtpNodes, cpNodes and upNodes matching the requested deployed state is/are returned along with site configurations and trafficPolicies
* `health_state` (String, optional) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `site_name` (String, optional) - mobility solution with requested site name is returned along with trafficPolicies
* `site_tag` (String, optional) - Mobility solution with sites matching the requested tag values is/are returned along with trafficPolicies
* `sites` (Attributes List, required) (see [below for nested schema](#nestedatt--sites))
* `solution_alias` (String, required) - Alias of the solution
* `solution_type` (String, required) - Type of the solution
* `tags` (Attributes List, optional) - RBAC Tags (see [below for nested schema](#nestedatt--tags))
* `traffic_policies` (Attributes, optional) - Forwarding policies for the processing nodes (see [below for nested schema](#nestedatt--traffic_policies))

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `deployed` (Boolean, computed) - Mobility solution with gtpNodes, cpNodes and upNodes matching the requested deployed state is/are returned along with site configurations and trafficPolicies
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `tags` (Attributes List, computed) - RBAC Tags (see [below for nested schema](#nestedatt--tags))
* `traffic_policies` (Attributes, computed) - Forwarding policies for the processing nodes (see [below for nested schema](#nestedatt--traffic_policies))

<a id="nestedatt--sites"></a>
### Nested Schema for `sites`

Required:

* `alias` (String) - Alias of the Site
* `tool_bindings` (Attributes List) - List of all tools used in the forwarding policies (see [below for nested schema](#nestedatt--sites--tool_bindings))
Optional:

* `cp_nodes` (Attributes List) (see [below for nested schema](#nestedatt--sites--cp_nodes))
* `gtp_nodes` (Attributes List) (see [below for nested schema](#nestedatt--sites--gtp_nodes))
* `network_ports` (List of String) - Network ports common to all the processing nodes. It is of the format 'cluster:port'.
* `sam_exporter_nodes` (Attributes List) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes))
* `site_override_of_policy_arrangements` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements))
* `skip_deployment` (Boolean) - When enabled, the deployment of the Site will be skipped
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--tags))
* `up_nodes` (Attributes List) (see [below for nested schema](#nestedatt--sites--up_nodes))
<a id="nestedatt--sites--tool_bindings"></a>
### Nested Schema for `sites.tool_bindings`

Required:

* `alias` (String) - Alias of the tool
* `tool_cluster_id` (String) - ClusterId in which the tool resource is present
* `tool_resource_type` (String) - Type of the tool resource
Optional:

* `meta_data_exporter_alias` (String) - Alias of the metadata exporter used for SAM Generation
* `tool_resource_id` (String) - ID of the tool resource
<a id="nestedatt--sites--cp_nodes"></a>
### Nested Schema for `sites.cp_nodes`

Required:

* `alias` (String) - Alias of the Control Node
* `location` (Attributes) - Location of the engine port (see [below for nested schema](#nestedatt--sites--cp_nodes--location))
* `node_type` (String) - Type of the Control Node
* `traffic_sources` (Attributes List) - List of all traffic sources for the Control node (see [below for nested schema](#nestedatt--sites--cp_nodes--traffic_sources))
Optional:

* `additional_tool_ports` (List of String) - Additional Tool ports for the Control node which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `app5_g_http2_ports` (List of Number) - List of TCP ports
* `app_tcp` (Attributes) - TCP Loadbalancing properties for Control 5G node (PCPN\_5G) (see [below for nested schema](#nestedatt--sites--cp_nodes--app_tcp))
* `collector_tools` (List of String) - Collector Tool ports. Its a list of ports of format cluster:port
* `control_metadata_ip_interface_alias` (String) - Alias of the ip interface used for exporting control json records
* `gtp_control_sample` (Boolean) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.
* `gtp_random_sampling` (Attributes) - GsGroup Gtp Random Sampling Parameters (see [below for nested schema](#nestedatt--sites--cp_nodes--gtp_random_sampling))
* `ip_interface_alias` (String) - Alias of the ip interface used by Control Node
* `node_override_network_ports` (List of String) - Network ports for the Control node. Its a list of ports of format cluster:port
* `number_of5_g_sessions` (Number) - Number of 5G sessions to allocate for Control 5G Node (PCPN\_5G)
* `number_of_lte_sessions` (Number) - Number of LTE sessions to allocate for Control LTE Node (PCPN\_LTE)
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--cp_nodes--tags))
Read-Only:

* `config_status` (String)
* `config_status_reasons` (String)
* `deployed` (Boolean) - True when the Control node attempted for deployment
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--sites--cp_nodes--health_state_reasons))
<a id="nestedatt--sites--cp_nodes--location"></a>
### Nested Schema for `sites.cp_nodes.location`

Required:

* `cluster_id` (String)
* `engine_ports` (List of String)
<a id="nestedatt--sites--cp_nodes--traffic_sources"></a>
### Nested Schema for `sites.cp_nodes.traffic_sources`

Required:

* `network_function_interfaces` (Attributes List) (see [below for nested schema](#nestedatt--sites--cp_nodes--traffic_sources--network_function_interfaces))
* `network_function_name` (String)
* `network_function_type` (String)
Optional:

* `comment` (String)
* `expand_port_identifier` (Boolean) - if enabled two rules added will be created with source port and destination port along with the other identifiers.
* `group_interfaces` (Boolean) - Enable if the user needs to combine rules in tunnelIdentifiers into single ingress map
* `ip4_frag_rule_type` (String) - if provided, the ip frag rule of given type would be appended to the rules.
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--cp_nodes--traffic_sources--tags))
<a id="nestedatt--sites--cp_nodes--traffic_sources--network_function_interfaces"></a>
### Nested Schema for `sites.cp_nodes.traffic_sources.network_function_interfaces`

Optional:

* `comment` (String)
* `interface_type` (String)
* `source_override_dst_ports` (List of String) - Tool ports which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `source_override_network_ports` (List of String) - Network ports for the network element. Its a list of ports of format cluster:port
* `tunnel_identifiers` (List of Dynamic)
<a id="nestedatt--sites--cp_nodes--traffic_sources--tags"></a>
### Nested Schema for `sites.cp_nodes.traffic_sources.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--cp_nodes--app_tcp"></a>
### Nested Schema for `sites.cp_nodes.app_tcp`

Optional:

* `application` (String) - To choose the action on Unknown Application Data
* `load_balance` (Boolean) - When true it enables TCP loadbalancing on the Tool Ports
* `tcp_control` (String) - To choose the action on TCP Control messages
<a id="nestedatt--sites--cp_nodes--gtp_random_sampling"></a>
### Nested Schema for `sites.cp_nodes.gtp_random_sampling`

Optional:

* `enabled` (Boolean) - When enabled, sampling of subscriber's sessions happens in random fashion
* `interval` (Number) - Rotation Interval in multiples of 12 (hrs)
<a id="nestedatt--sites--cp_nodes--tags"></a>
### Nested Schema for `sites.cp_nodes.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--cp_nodes--health_state_reasons"></a>
### Nested Schema for `sites.cp_nodes.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--sites--gtp_nodes"></a>
### Nested Schema for `sites.gtp_nodes`

Required:

* `alias` (String) - Alias of the GTP node
* `location` (Attributes) - Location of the engine port (see [below for nested schema](#nestedatt--sites--gtp_nodes--location))
* `node_type` (String) - Type of the GTP Node
* `traffic_sources` (Attributes List) - List of all traffic sources for the GTP node (see [below for nested schema](#nestedatt--sites--gtp_nodes--traffic_sources))
Optional:

* `additional_tool_ports` (List of String) - Additional Tool ports for the GTP node which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `collector_tools` (List of String) - Collector Tool ports. Its a list of ports of format cluster:port
* `gtp_random_sampling` (Attributes) - GsGroup Gtp Random Sampling Parameters (see [below for nested schema](#nestedatt--sites--gtp_nodes--gtp_random_sampling))
* `node_override_network_ports` (List of String) - Network ports for the GTP node. Its a list of ports of format cluster:port
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--gtp_nodes--tags))
Read-Only:

* `config_status` (String)
* `config_status_reasons` (String)
* `deployed` (Boolean) - True when the GTP node is attempted for deployment
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--sites--gtp_nodes--health_state_reasons))
<a id="nestedatt--sites--gtp_nodes--location"></a>
### Nested Schema for `sites.gtp_nodes.location`

Required:

* `cluster_id` (String)
* `engine_ports` (List of String)
<a id="nestedatt--sites--gtp_nodes--traffic_sources"></a>
### Nested Schema for `sites.gtp_nodes.traffic_sources`

Required:

* `network_function_interfaces` (Attributes List) (see [below for nested schema](#nestedatt--sites--gtp_nodes--traffic_sources--network_function_interfaces))
* `network_function_name` (String)
* `network_function_type` (String)
Optional:

* `comment` (String)
* `expand_port_identifier` (Boolean) - if enabled two rules added will be created with source port and destination port along with the other identifiers.
* `group_interfaces` (Boolean) - Enable if the user needs to combine rules in tunnelIdentifiers into single ingress map
* `ip4_frag_rule_type` (String) - if provided, the ip frag rule of given type would be appended to the rules.
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--gtp_nodes--traffic_sources--tags))
<a id="nestedatt--sites--gtp_nodes--traffic_sources--network_function_interfaces"></a>
### Nested Schema for `sites.gtp_nodes.traffic_sources.network_function_interfaces`

Optional:

* `comment` (String)
* `interface_type` (String)
* `source_override_dst_ports` (List of String) - Tool ports which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `source_override_network_ports` (List of String) - Network ports for the network element. Its a list of ports of format cluster:port
* `tunnel_identifiers` (List of Dynamic)
<a id="nestedatt--sites--gtp_nodes--traffic_sources--tags"></a>
### Nested Schema for `sites.gtp_nodes.traffic_sources.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--gtp_nodes--gtp_random_sampling"></a>
### Nested Schema for `sites.gtp_nodes.gtp_random_sampling`

Optional:

* `enabled` (Boolean) - When enabled, sampling of subscriber's sessions happens in random fashion
* `interval` (Number) - Rotation Interval in multiples of 12 (hrs)
<a id="nestedatt--sites--gtp_nodes--tags"></a>
### Nested Schema for `sites.gtp_nodes.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--gtp_nodes--health_state_reasons"></a>
### Nested Schema for `sites.gtp_nodes.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--sites--sam_exporter_nodes"></a>
### Nested Schema for `sites.sam_exporter_nodes`

Required:

* `alias` (String) - Alias of the SAM Exporter node
* `location` (Attributes) - Location of the engine port (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--location))
* `node_type` (String) - Type of the SAM Exporter Node
Optional:

* `smaf_details` (Attributes List) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--smaf_details))
* `app_profile_config` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config))
* `control_plane_setting` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--control_plane_setting))
* `engine_meta_data_cache_configs` (Attributes List) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs))
* `engine_source_mappings` (Attributes List) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_source_mappings))
* `exporter_config` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--exporter_config))
* `ip_interface_alias` (String) - Alias of the ip interface used by SAM Exporter Node
* `node_override_network_ports` (List of String) - Network ports for the SAM Exporter node. Its a list of ports of format cluster:port
* `param_configs` (Attributes List) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--param_configs))
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--tags))
* `traffic_sources` (Attributes List) - List of all traffic sources for the SAM Exporter node (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--traffic_sources))
Read-Only:

* `config_status` (String)
* `config_status_reasons` (String)
* `deployed` (Boolean) - True when the SAM node is attempted for deployment
* `deployment_details` (Attributes List) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--deployment_details))
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--health_state_reasons))
<a id="nestedatt--sites--sam_exporter_nodes--location"></a>
### Nested Schema for `sites.sam_exporter_nodes.location`

Required:

* `cluster_id` (String)
* `engine_ports` (List of String)
<a id="nestedatt--sites--sam_exporter_nodes--smaf_details"></a>
### Nested Schema for `sites.sam_exporter_nodes.smaf_details`

Optional:

* `control_application` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--smaf_details--control_application))
* `management_address` (String)
* `user_application` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--smaf_details--user_application))
<a id="nestedatt--sites--sam_exporter_nodes--smaf_details--control_application"></a>
### Nested Schema for `sites.sam_exporter_nodes.smaf_details.control_application`

Optional:

* `interface_address` (String)
* `port` (Number)
* `protocol` (String)
<a id="nestedatt--sites--sam_exporter_nodes--smaf_details--user_application"></a>
### Nested Schema for `sites.sam_exporter_nodes.smaf_details.user_application`

Optional:

* `interface_address` (String)
* `port` (Number)
* `protocol` (String)
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config`

Optional:

* `application_id` (Boolean)
* `applications` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--counter))
* `flow` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--gtpu))
* `inner_ipv4` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--inner_ipv4))
* `inner_ipv6` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--inner_ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--transport))
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--applications"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.applications`

Required:

* `name` (String) - application name
Optional:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--app_profile_config--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--applications--attributes"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.applications.attributes`

Required:

* `name` (String) - attribute name
Optional:

* `value` (String) - application's attribute value
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--counter"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.counter`

Optional:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--flow"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.flow`

Optional:

* `end_reason` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--gtpu"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.gtpu`

Optional:

* `qfi` (Boolean)
* `teid` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--inner_ipv4"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.inner_ipv4`

Optional:

* `destination` (Boolean)
* `protocol` (Boolean)
* `source` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--inner_ipv6"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.inner_ipv6`

Optional:

* `destination` (Boolean)
* `next_header` (Boolean)
* `source` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--outer_ipv4"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.outer_ipv4`

Optional:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--outer_ipv6"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.outer_ipv6`

Optional:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--timestamp"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.timestamp`

Optional:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--app_profile_config--transport"></a>
### Nested Schema for `sites.sam_exporter_nodes.app_profile_config.transport`

Optional:

* `dst_port` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--control_plane_setting"></a>
### Nested Schema for `sites.sam_exporter_nodes.control_plane_setting`

Optional:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--control_plane_setting--event_enable))
* `trigger` (String)
<a id="nestedatt--sites--sam_exporter_nodes--control_plane_setting--event_enable"></a>
### Nested Schema for `sites.sam_exporter_nodes.control_plane_setting.event_enable`

Optional:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_meta_data_cache_configs`

Optional:

* `engine_port` (String)
* `event` (String)
* `flow_behavior` (String)
* `flows_size` (Number)
* `idle_timeout` (Number)
* `match` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match))
* `observation_domain_id` (Number)
<a id="nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_meta_data_cache_configs.match`

Optional:

* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv6))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--transport))
<a id="nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv4"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_meta_data_cache_configs.match.ipv4`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv4--destination))
* `protocol` (Boolean)
* `source` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv4--source))
<a id="nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv4--destination"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_meta_data_cache_configs.match.ipv4.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv4--source"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_meta_data_cache_configs.match.ipv4.source`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv6"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_meta_data_cache_configs.match.ipv6`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv6--destination))
* `next_header` (Boolean)
* `source` (Attributes) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv6--source))
<a id="nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv6--destination"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_meta_data_cache_configs.match.ipv6.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--ipv6--source"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_meta_data_cache_configs.match.ipv6.source`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--sites--sam_exporter_nodes--engine_meta_data_cache_configs--match--transport"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_meta_data_cache_configs.match.transport`

Optional:

* `dst_port` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--sites--sam_exporter_nodes--engine_source_mappings"></a>
### Nested Schema for `sites.sam_exporter_nodes.engine_source_mappings`

Required:

* `engine_port` (String)
* `network_source` (String)
<a id="nestedatt--sites--sam_exporter_nodes--exporter_config"></a>
### Nested Schema for `sites.sam_exporter_nodes.exporter_config`

Optional:

* `active_timeout` (String)
* `inactive_timeout` (Number)
* `record_type` (String)
<a id="nestedatt--sites--sam_exporter_nodes--param_configs"></a>
### Nested Schema for `sites.sam_exporter_nodes.param_configs`

Optional:

* `engine_port` (String)
* `resource_metadata` (Number)
<a id="nestedatt--sites--sam_exporter_nodes--tags"></a>
### Nested Schema for `sites.sam_exporter_nodes.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--sam_exporter_nodes--traffic_sources"></a>
### Nested Schema for `sites.sam_exporter_nodes.traffic_sources`

Required:

* `network_function_interfaces` (Attributes List) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--traffic_sources--network_function_interfaces))
* `network_function_name` (String)
* `network_function_type` (String)
Optional:

* `comment` (String)
* `expand_port_identifier` (Boolean) - if enabled two rules added will be created with source port and destination port along with the other identifiers.
* `group_interfaces` (Boolean) - Enable if the user needs to combine rules in tunnelIdentifiers into single ingress map
* `ip4_frag_rule_type` (String) - if provided, the ip frag rule of given type would be appended to the rules.
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--sam_exporter_nodes--traffic_sources--tags))
<a id="nestedatt--sites--sam_exporter_nodes--traffic_sources--network_function_interfaces"></a>
### Nested Schema for `sites.sam_exporter_nodes.traffic_sources.network_function_interfaces`

Optional:

* `comment` (String)
* `interface_type` (String)
* `source_override_dst_ports` (List of String) - Tool ports which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `source_override_network_ports` (List of String) - Network ports for the network element. Its a list of ports of format cluster:port
* `tunnel_identifiers` (List of Dynamic)
<a id="nestedatt--sites--sam_exporter_nodes--traffic_sources--tags"></a>
### Nested Schema for `sites.sam_exporter_nodes.traffic_sources.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--sam_exporter_nodes--deployment_details"></a>
### Nested Schema for `sites.sam_exporter_nodes.deployment_details`

Read-Only:

* `engine_port` (String)
* `sam_node_alias` (String)
<a id="nestedatt--sites--sam_exporter_nodes--health_state_reasons"></a>
### Nested Schema for `sites.sam_exporter_nodes.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--sites--site_override_of_policy_arrangements"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements`

Optional:

* `for5_g` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g))
* `for_lte` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte))
* `for_non_cups_lte` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g`

Optional:

* `gtp_flow_timeout` (Number)
* `gtp_persistence` (Attributes) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--gtp_persistence))
* `load_balancing` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--load_balancing))
* `overlap_mode` (Boolean) - When enabled flow-filtering cannot be configured
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--sampling))
* `whitelisting` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--whitelisting))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--gtp_persistence"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.gtp_persistence`

Optional:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--load_balancing"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.load_balancing`

Required:

* `app_type` (String)
* `hashing_key` (String)
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--sampling"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.sampling`

Optional:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--sampling--flow_maps))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--sampling--flow_maps"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.sampling.flow_maps`

Required:

* `alias` (String) - Alias of the sampling map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--sampling--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--sampling--flow_maps--tags))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--sampling--flow_maps--rules"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.sampling.flow_maps.rules`

Required:

* `control_plane_percentage` (Number)
* `user_plane_percentage` (Number)
Optional:

* `comment` (String)
* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--sampling--flow_maps--tags"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.sampling.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--whitelisting"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.whitelisting`

Required:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--whitelisting--flow_maps))
Optional:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
* `white_list_alias` (String) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--whitelisting--flow_maps"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.whitelisting.flow_maps`

Required:

* `alias` (String) - Alias of the whitelist map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--whitelisting--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for5_g--whitelisting--flow_maps--tags))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--whitelisting--flow_maps--rules"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.whitelisting.flow_maps.rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--sites--site_override_of_policy_arrangements--for5_g--whitelisting--flow_maps--tags"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for5_g.whitelisting.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte`

Optional:

* `gtp_flow_timeout` (Number)
* `gtp_persistence` (Attributes) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte--gtp_persistence))
* `overlap_mode` (Boolean) - When enabled flow-filtering cannot be configured
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte--sampling))
* `whitelisting` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte--whitelisting))
Read-Only:

* `load_balancing_lte` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte--load_balancing_lte))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte--gtp_persistence"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte.gtp_persistence`

Optional:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte--sampling"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte.sampling`

Optional:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte--sampling--flow_maps))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte--sampling--flow_maps"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte.sampling.flow_maps`

Required:

* `alias` (String) - Alias of the sampling map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte--sampling--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte--sampling--flow_maps--tags))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte--sampling--flow_maps--rules"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte.sampling.flow_maps.rules`

Required:

* `control_plane_percentage` (Number)
* `user_plane_percentage` (Number)
Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `comment` (String)
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte--sampling--flow_maps--tags"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte.sampling.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte--whitelisting"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte.whitelisting`

Required:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte--whitelisting--flow_maps))
Optional:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
* `white_list_alias` (String) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte--whitelisting--flow_maps"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte.whitelisting.flow_maps`

Required:

* `alias` (String) - Alias of the whitelist map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_lte--whitelisting--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte--whitelisting--flow_maps--rules"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte.whitelisting.flow_maps.rules`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_lte--load_balancing_lte"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_lte.load_balancing_lte`

Read-Only:

* `app_type` (String)
* `hashing_key` (String)
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte`

Optional:

* `flowfiltering` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering))
* `gtp_flow_timeout` (Number)
* `gtp_persistence` (Attributes) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--gtp_persistence))
* `load_balancing` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--load_balancing))
* `overlap_mode` (Boolean) - When enabled flow-filtering cannot be configured
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--sampling))
* `whitelisting` (Attributes) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--whitelisting))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.flowfiltering`

Optional:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering--flow_maps))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering--flow_maps"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.flowfiltering.flow_maps`

Required:

* `alias` (String) - Alias of the flow-filtering map
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering--flow_maps--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering--flow_maps--pass_rules))
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering--flow_maps--tags))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering--flow_maps--drop_rules"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.flowfiltering.flow_maps.drop_rules`

Optional:

* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering--flow_maps--pass_rules"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.flowfiltering.flow_maps.pass_rules`

Optional:

* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--flowfiltering--flow_maps--tags"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.flowfiltering.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--gtp_persistence"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.gtp_persistence`

Optional:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--load_balancing"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.load_balancing`

Required:

* `app_type` (String)
* `hashing_key` (String)
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--sampling"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.sampling`

Optional:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--sampling--flow_maps))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--sampling--flow_maps"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.sampling.flow_maps`

Required:

* `alias` (String) - Alias of the sampling map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--sampling--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--sampling--flow_maps--tags))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--sampling--flow_maps--rules"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.sampling.flow_maps.rules`

Required:

* `gtp_sample_percentage` (Number)
Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `comment` (String)
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--sampling--flow_maps--tags"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.sampling.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--whitelisting"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.whitelisting`

Required:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--whitelisting--flow_maps))
Optional:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
* `white_list_alias` (String) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--whitelisting--flow_maps"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.whitelisting.flow_maps`

Required:

* `alias` (String) - Alias of the whitelist map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--whitelisting--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--whitelisting--flow_maps--tags))
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--whitelisting--flow_maps--rules"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.whitelisting.flow_maps.rules`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--sites--site_override_of_policy_arrangements--for_non_cups_lte--whitelisting--flow_maps--tags"></a>
### Nested Schema for `sites.site_override_of_policy_arrangements.for_non_cups_lte.whitelisting.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--tags"></a>
### Nested Schema for `sites.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--up_nodes"></a>
### Nested Schema for `sites.up_nodes`

Required:

* `alias` (String) - Alias of the User node
* `location` (Attributes) - Location of the engine port (see [below for nested schema](#nestedatt--sites--up_nodes--location))
* `node_type` (String) - Type of the User Node
* `traffic_sources` (Attributes List) - List of all traffic sources for the User node (see [below for nested schema](#nestedatt--sites--up_nodes--traffic_sources))
Optional:

* `additional_tool_ports` (List of String) - Additional Tool ports for the User node which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `collector_tools` (List of String) - Collector Tool ports. Its a list of ports of format cluster:port
* `gtp_control_sample` (Boolean) - When enabled, GTP Control plane traffic for the subscribers contained within the subscriber sample defined in GTP flowsample maps will be sent to tools. When disabled, all GTP Control plane traffic for all subscribers in GTP flowsample maps will be sent to tools.
* `gtp_random_sampling` (Attributes) - GsGroup Gtp Random Sampling Parameters (see [below for nested schema](#nestedatt--sites--up_nodes--gtp_random_sampling))
* `ip_interface_alias` (String) - Alias of the ip interface used by User Node
* `node_override_network_ports` (List of String) - Network ports for the User node. Its a list of ports of format cluster:port
* `number_of_lte_sessions` (Number) - Number of LTE sessions to allocate for UPN Node
* `stand_alone_mode` (Boolean) - Enable or Disable standAloneMode in User Nodes
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--up_nodes--tags))
Read-Only:

* `config_status` (String)
* `config_status_reasons` (String)
* `deployed` (Boolean) - True when the User node is attempted for deployment
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--sites--up_nodes--health_state_reasons))
<a id="nestedatt--sites--up_nodes--location"></a>
### Nested Schema for `sites.up_nodes.location`

Required:

* `cluster_id` (String)
* `engine_ports` (List of String)
<a id="nestedatt--sites--up_nodes--traffic_sources"></a>
### Nested Schema for `sites.up_nodes.traffic_sources`

Required:

* `network_function_interfaces` (Attributes List) (see [below for nested schema](#nestedatt--sites--up_nodes--traffic_sources--network_function_interfaces))
* `network_function_name` (String)
* `network_function_type` (String)
Optional:

* `comment` (String)
* `expand_port_identifier` (Boolean) - if enabled two rules added will be created with source port and destination port along with the other identifiers.
* `group_interfaces` (Boolean) - Enable if the user needs to combine rules in tunnelIdentifiers into single ingress map
* `ip4_frag_rule_type` (String) - if provided, the ip frag rule of given type would be appended to the rules.
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--sites--up_nodes--traffic_sources--tags))
<a id="nestedatt--sites--up_nodes--traffic_sources--network_function_interfaces"></a>
### Nested Schema for `sites.up_nodes.traffic_sources.network_function_interfaces`

Optional:

* `comment` (String)
* `interface_type` (String)
* `source_override_dst_ports` (List of String) - Tool ports which takes a copy of traffic along with vport that is added into the first level maps. Its a list of ports of format cluster:port
* `source_override_network_ports` (List of String) - Network ports for the network element. Its a list of ports of format cluster:port
* `tunnel_identifiers` (List of Dynamic)
<a id="nestedatt--sites--up_nodes--traffic_sources--tags"></a>
### Nested Schema for `sites.up_nodes.traffic_sources.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--up_nodes--gtp_random_sampling"></a>
### Nested Schema for `sites.up_nodes.gtp_random_sampling`

Optional:

* `enabled` (Boolean) - When enabled, sampling of subscriber's sessions happens in random fashion
* `interval` (Number) - Rotation Interval in multiples of 12 (hrs)
<a id="nestedatt--sites--up_nodes--tags"></a>
### Nested Schema for `sites.up_nodes.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--sites--up_nodes--health_state_reasons"></a>
### Nested Schema for `sites.up_nodes.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_policies"></a>
### Nested Schema for `traffic_policies`

Optional:

* `for5_g` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for5_g))
* `for_lte` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for_lte))
* `for_non_cups_lte` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte))
<a id="nestedatt--traffic_policies--for5_g"></a>
### Nested Schema for `traffic_policies.for5_g`

Optional:

* `gtp_flow_timeout` (Number)
* `gtp_persistence` (Attributes) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--traffic_policies--for5_g--gtp_persistence))
* `load_balancing` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for5_g--load_balancing))
* `overlap_mode` (Boolean) - When enabled flow-filtering cannot be configured
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for5_g--sampling))
* `whitelisting` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for5_g--whitelisting))
<a id="nestedatt--traffic_policies--for5_g--gtp_persistence"></a>
### Nested Schema for `traffic_policies.for5_g.gtp_persistence`

Optional:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)
<a id="nestedatt--traffic_policies--for5_g--load_balancing"></a>
### Nested Schema for `traffic_policies.for5_g.load_balancing`

Required:

* `app_type` (String)
* `hashing_key` (String)
<a id="nestedatt--traffic_policies--for5_g--sampling"></a>
### Nested Schema for `traffic_policies.for5_g.sampling`

Optional:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for5_g--sampling--flow_maps))
<a id="nestedatt--traffic_policies--for5_g--sampling--flow_maps"></a>
### Nested Schema for `traffic_policies.for5_g.sampling.flow_maps`

Required:

* `alias` (String) - Alias of the sampling map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for5_g--sampling--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--traffic_policies--for5_g--sampling--flow_maps--tags))
<a id="nestedatt--traffic_policies--for5_g--sampling--flow_maps--rules"></a>
### Nested Schema for `traffic_policies.for5_g.sampling.flow_maps.rules`

Required:

* `control_plane_percentage` (Number)
* `user_plane_percentage` (Number)
Optional:

* `comment` (String)
* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--traffic_policies--for5_g--sampling--flow_maps--tags"></a>
### Nested Schema for `traffic_policies.for5_g.sampling.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_policies--for5_g--whitelisting"></a>
### Nested Schema for `traffic_policies.for5_g.whitelisting`

Required:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for5_g--whitelisting--flow_maps))
Optional:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
* `white_list_alias` (String) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
<a id="nestedatt--traffic_policies--for5_g--whitelisting--flow_maps"></a>
### Nested Schema for `traffic_policies.for5_g.whitelisting.flow_maps`

Required:

* `alias` (String) - Alias of the whitelist map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for5_g--whitelisting--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--traffic_policies--for5_g--whitelisting--flow_maps--tags))
<a id="nestedatt--traffic_policies--for5_g--whitelisting--flow_maps--rules"></a>
### Nested Schema for `traffic_policies.for5_g.whitelisting.flow_maps.rules`

Optional:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--traffic_policies--for5_g--whitelisting--flow_maps--tags"></a>
### Nested Schema for `traffic_policies.for5_g.whitelisting.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_policies--for_lte"></a>
### Nested Schema for `traffic_policies.for_lte`

Optional:

* `gtp_flow_timeout` (Number)
* `gtp_persistence` (Attributes) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--traffic_policies--for_lte--gtp_persistence))
* `overlap_mode` (Boolean) - When enabled flow-filtering cannot be configured
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for_lte--sampling))
* `whitelisting` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for_lte--whitelisting))
Read-Only:

* `load_balancing_lte` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for_lte--load_balancing_lte))
<a id="nestedatt--traffic_policies--for_lte--gtp_persistence"></a>
### Nested Schema for `traffic_policies.for_lte.gtp_persistence`

Optional:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)
<a id="nestedatt--traffic_policies--for_lte--sampling"></a>
### Nested Schema for `traffic_policies.for_lte.sampling`

Optional:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_lte--sampling--flow_maps))
<a id="nestedatt--traffic_policies--for_lte--sampling--flow_maps"></a>
### Nested Schema for `traffic_policies.for_lte.sampling.flow_maps`

Required:

* `alias` (String) - Alias of the sampling map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_lte--sampling--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--traffic_policies--for_lte--sampling--flow_maps--tags))
<a id="nestedatt--traffic_policies--for_lte--sampling--flow_maps--rules"></a>
### Nested Schema for `traffic_policies.for_lte.sampling.flow_maps.rules`

Required:

* `control_plane_percentage` (Number)
* `user_plane_percentage` (Number)
Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `comment` (String)
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--traffic_policies--for_lte--sampling--flow_maps--tags"></a>
### Nested Schema for `traffic_policies.for_lte.sampling.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_policies--for_lte--whitelisting"></a>
### Nested Schema for `traffic_policies.for_lte.whitelisting`

Required:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_lte--whitelisting--flow_maps))
Optional:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
* `white_list_alias` (String) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
<a id="nestedatt--traffic_policies--for_lte--whitelisting--flow_maps"></a>
### Nested Schema for `traffic_policies.for_lte.whitelisting.flow_maps`

Required:

* `alias` (String) - Alias of the whitelist map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_lte--whitelisting--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
<a id="nestedatt--traffic_policies--for_lte--whitelisting--flow_maps--rules"></a>
### Nested Schema for `traffic_policies.for_lte.whitelisting.flow_maps.rules`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--traffic_policies--for_lte--load_balancing_lte"></a>
### Nested Schema for `traffic_policies.for_lte.load_balancing_lte`

Read-Only:

* `app_type` (String)
* `hashing_key` (String)
<a id="nestedatt--traffic_policies--for_non_cups_lte"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte`

Optional:

* `flowfiltering` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--flowfiltering))
* `gtp_flow_timeout` (Number)
* `gtp_persistence` (Attributes) - GsGroup Gtp Persistence Parameters (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--gtp_persistence))
* `load_balancing` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--load_balancing))
* `overlap_mode` (Boolean) - When enabled flow-filtering cannot be configured
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--sampling))
* `whitelisting` (Attributes) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--whitelisting))
<a id="nestedatt--traffic_policies--for_non_cups_lte--flowfiltering"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.flowfiltering`

Optional:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--flowfiltering--flow_maps))
<a id="nestedatt--traffic_policies--for_non_cups_lte--flowfiltering--flow_maps"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.flowfiltering.flow_maps`

Required:

* `alias` (String) - Alias of the flow-filtering map
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--flowfiltering--flow_maps--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--flowfiltering--flow_maps--pass_rules))
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--flowfiltering--flow_maps--tags))
<a id="nestedatt--traffic_policies--for_non_cups_lte--flowfiltering--flow_maps--drop_rules"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.flowfiltering.flow_maps.drop_rules`

Optional:

* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--traffic_policies--for_non_cups_lte--flowfiltering--flow_maps--pass_rules"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.flowfiltering.flow_maps.pass_rules`

Optional:

* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--traffic_policies--for_non_cups_lte--flowfiltering--flow_maps--tags"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.flowfiltering.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_policies--for_non_cups_lte--gtp_persistence"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.gtp_persistence`

Optional:

* `enabled` (Boolean) - GTP Persistence Status
* `file_age_timeout` (Number) - GTP Persistence File Age Timeout(mins)
* `interval` (Number) - GTP Persistence Interval(mins)
* `restart_age_time` (Number) - GTP Persistence Restart Age Time(mins)
<a id="nestedatt--traffic_policies--for_non_cups_lte--load_balancing"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.load_balancing`

Required:

* `app_type` (String)
* `hashing_key` (String)
<a id="nestedatt--traffic_policies--for_non_cups_lte--sampling"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.sampling`

Optional:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--sampling--flow_maps))
<a id="nestedatt--traffic_policies--for_non_cups_lte--sampling--flow_maps"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.sampling.flow_maps`

Required:

* `alias` (String) - Alias of the sampling map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--sampling--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--sampling--flow_maps--tags))
<a id="nestedatt--traffic_policies--for_non_cups_lte--sampling--flow_maps--rules"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.sampling.flow_maps.rules`

Required:

* `gtp_sample_percentage` (Number)
Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `comment` (String)
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--traffic_policies--for_non_cups_lte--sampling--flow_maps--tags"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.sampling.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--traffic_policies--for_non_cups_lte--whitelisting"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.whitelisting`

Required:

* `flow_maps` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--whitelisting--flow_maps))
Optional:

* `multi_whitelists` (List of String) - Alias/Aliases of referenced GTP Whitelists.Use either whiteListAlias or multiWhitelists.
* `white_list_alias` (String) - Alias of referenced GTP Whitelist.Use either whiteListAlias or multiWhitelists.
<a id="nestedatt--traffic_policies--for_non_cups_lte--whitelisting--flow_maps"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.whitelisting.flow_maps`

Required:

* `alias` (String) - Alias of the whitelist map
* `rules` (Attributes List) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--whitelisting--flow_maps--rules))
* `tool` (String) - Alias of referenced tool
Optional:

* `comment` (String)
* `source_group_id` (String) - if provided a vport with this sourceGroupId would be created else default vport would be used
* `tags` (Attributes List) - User defined tags (Aggregation tags) (see [below for nested schema](#nestedatt--traffic_policies--for_non_cups_lte--whitelisting--flow_maps--tags))
<a id="nestedatt--traffic_policies--for_non_cups_lte--whitelisting--flow_maps--rules"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.whitelisting.flow_maps.rules`

Optional:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--traffic_policies--for_non_cups_lte--whitelisting--flow_maps--tags"></a>
### Nested Schema for `traffic_policies.for_non_cups_lte.whitelisting.flow_maps.tags`

Required:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_mobility.example {solution_alias}
```
