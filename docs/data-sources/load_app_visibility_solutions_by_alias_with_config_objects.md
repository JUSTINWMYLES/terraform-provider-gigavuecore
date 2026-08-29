---
page_title: "gigavuecore_load_app_visibility_solutions_by_alias_with_config_objects Data Source - gigavuecore"
subcategory: ""
description: |-
  Load an Application Intelligence Solution by alias with config objects
---

# gigavuecore_load_app_visibility_solutions_by_alias_with_config_objects Data Source

Load an Application Intelligence Solution by alias with config objects

## Example Usage

```terraform
data "gigavuecore_load_app_visibility_solutions_by_alias_with_config_objects" "example" {
  solution_alias = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `solution_alias` (String, required) - Alias of the solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `app_env_config` (Attributes, computed) (see [below for nested schema](#nestedatt--app_env_config))
* `app_env_id` (String, computed) - ID of the Application Environment
* `app_exporter_config` (Attributes, computed) (see [below for nested schema](#nestedatt--app_exporter_config))
* `app_filter_config` (Attributes, computed) (see [below for nested schema](#nestedatt--app_filter_config))
* `config_status` (String, computed) - configuration status
* `config_status_reason` (String, computed) - configuration status reasons
* `conn_id` (String, computed) - ID of the Connection Domain
* `dedup` (Attributes, computed) (see [below for nested schema](#nestedatt--dedup))
* `dynamic_scale_unit` (Boolean, computed) - When set as true, FM automatically sets the optimal scaleUnit
* `env_id` (String, computed) - ID of the Environment
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `ingress_traffic_configs` (Attributes List, computed) (see [below for nested schema](#nestedatt--ingress_traffic_configs))
* `monitor_solution_config` (Attributes, computed) (see [below for nested schema](#nestedatt--monitor_solution_config))
* `scale_unit` (Number, computed) - Number of units of memory needed for each Application
* `solution_created_timestamp` (String, computed)
* `solution_desc` (String, computed) - Description of the solution
* `solution_modified_timestamp` (String, computed)
* `traffic_policy_graph_alias` (String, computed) - Alias of the Traffic Policy Graph

<a id="nestedatt--app_env_config"></a>
### Nested Schema for `app_env_config`

Read-Only:

* `app_info` (String)
* `argv` (String)
* `name` (String)
* `nr_cores` (Number)
* `nr_s_cores` (Number)
<a id="nestedatt--app_exporter_config"></a>
### Nested Schema for `app_exporter_config`

Read-Only:

* `app_instance_name` (String)
* `app_metadata_tiered_batch_id` (String) - ID of the appExporterConfig tiered batch
* `cache_config` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config))
* `cache_config_alias` (String) - Alias of the cacheConfig
* `destination_configs` (Attributes List) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs))
<a id="nestedatt--app_exporter_config--cache_config"></a>
### Nested Schema for `app_exporter_config.cache_config`

Read-Only:

* `advance_hash` (Boolean) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
* `alias` (String)
* `description` (String)
* `dpi_inject_limit` (Number)
* `event` (String)
* `exporters` (List of String) - alias of metadata exporters to attach this cache
* `flow_behavior` (String) - direction for flow identification
* `match` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match))
* `multi_collect` (Boolean) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
* `network_profiles` (List of String) - alias of metadata network profiles to attach this cache
* `observation_domain_id` (Number)
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--sampling))
* `size` (Attributes) - size of the flows (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--size))
* `timeout` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--timeout))
<a id="nestedatt--app_exporter_config--cache_config--match"></a>
### Nested Schema for `app_exporter_config.cache_config.match`

Read-Only:

* `datalink` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--datalink))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv6))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--transport))
<a id="nestedatt--app_exporter_config--cache_config--match--datalink"></a>
### Nested Schema for `app_exporter_config.cache_config.match.datalink`

Read-Only:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--match--interface"></a>
### Nested Schema for `app_exporter_config.cache_config.match.interface`

Read-Only:

* `in_name_width` (Number)
* `in_physical_width` (Number)
<a id="nestedatt--app_exporter_config--cache_config--match--ip"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ip`

Read-Only:

* `version` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--match--ipv4"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv4`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--match--ipv4--destination"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv4.destination`

Read-Only:

* `prefix_min_mask` (String) - ipv4 destination prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_exporter_config--cache_config--match--ipv4--fragmentation"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv4.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--match--ipv4--section"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv4.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_exporter_config--cache_config--match--ipv4--source"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv4.source`

Read-Only:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_exporter_config--cache_config--match--ipv6"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv6`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--match--ipv6--destination"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv6.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--app_exporter_config--cache_config--match--ipv6--fragmentation"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv6.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--match--ipv6--length"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv6.length`

Read-Only:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--match--ipv6--section"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv6.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_exporter_config--cache_config--match--ipv6--source"></a>
### Nested Schema for `app_exporter_config.cache_config.match.ipv6.source`

Read-Only:

* `prefix_min_mask` (String) - ipv6 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_exporter_config--cache_config--match--transport"></a>
### Nested Schema for `app_exporter_config.cache_config.match.transport`

Read-Only:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--cache_config--match--transport--udp))
<a id="nestedatt--app_exporter_config--cache_config--match--transport--icmp"></a>
### Nested Schema for `app_exporter_config.cache_config.match.transport.icmp`

Read-Only:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--match--transport--tcp"></a>
### Nested Schema for `app_exporter_config.cache_config.match.transport.tcp`

Read-Only:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--match--transport--udp"></a>
### Nested Schema for `app_exporter_config.cache_config.match.transport.udp`

Read-Only:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--app_exporter_config--cache_config--sampling"></a>
### Nested Schema for `app_exporter_config.cache_config.sampling`

Read-Only:

* `mode` (String)
* `single_sampling_rate` (Number) - Packet interval window size. Valid values: 10-16000 (in packets)
<a id="nestedatt--app_exporter_config--cache_config--size"></a>
### Nested Schema for `app_exporter_config.cache_config.size`

Read-Only:

* `flows` (Number) - size of flows in millions
<a id="nestedatt--app_exporter_config--cache_config--timeout"></a>
### Nested Schema for `app_exporter_config.cache_config.timeout`

Read-Only:

* `idle` (Number) - idle timeout in seconds. max value 7days. default 30 min
<a id="nestedatt--app_exporter_config--destination_configs"></a>
### Nested Schema for `app_exporter_config.destination_configs`

Read-Only:

* `application_names` (Attributes List) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--application_names))
* `destination_name` (String)
* `export_meta_app_profile` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile))
* `export_meta_app_profile_alias` (String)
* `exporter_alias` (String)
* `exporter_config` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--exporter_config))
* `iface` (String) - Interface Mapping for Egress Tunnel
<a id="nestedatt--app_exporter_config--destination_configs--application_names"></a>
### Nested Schema for `app_exporter_config.destination_configs.application_names`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name
<a id="nestedatt--app_exporter_config--destination_configs--application_names--attributes"></a>
### Nested Schema for `app_exporter_config.destination_configs.application_names.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile`

Read-Only:

* `alias` (String) - application profile alias
* `application_id` (Boolean) - only valid with 'export' type
* `applications` (Attributes List) - application and attributes. (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--counter))
* `datalink` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--datalink))
* `description` (String)
* `flow` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--gtpu))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--transport))
* `type` (String)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--applications"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.applications`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--applications--attributes"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.applications.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--counter"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.counter`

Read-Only:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `inner_byte` (Boolean)
* `inner_byte_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--datalink"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.datalink`

Read-Only:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--flow"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.flow`

Read-Only:

* `end_reason` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--gtpu"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.gtpu`

Read-Only:

* `qfi` (Boolean)
* `teid` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--interface"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.interface`

Read-Only:

* `in_name_width` (Number)
* `in_physical_width` (Number)
* `out_physical_width` (Number)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ip"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ip`

Read-Only:

* `version` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv4`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4--destination"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv4.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4--fragmentation"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv4.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4--section"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv4.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv4--source"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv4.source`

Read-Only:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv6`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--destination"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv6.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--fragmentation"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv6.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--length"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv6.length`

Read-Only:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--section"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv6.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--ipv6--source"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.ipv6.source`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--outer_ipv4"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.outer_ipv4`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--outer_ipv6"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.outer_ipv6`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--timestamp"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.timestamp`

Read-Only:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
* `sys_up_time_first` (Boolean)
* `sys_up_time_last` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--transport"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.transport`

Read-Only:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--transport--udp))
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--transport--icmp"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.transport.icmp`

Read-Only:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--transport--tcp"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.transport.tcp`

Read-Only:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--export_meta_app_profile--transport--udp"></a>
### Nested Schema for `app_exporter_config.destination_configs.export_meta_app_profile.transport.udp`

Read-Only:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--exporter_config"></a>
### Nested Schema for `app_exporter_config.destination_configs.exporter_config`

Read-Only:

* `alias` (String)
* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--exporter_config--cef))
* `description` (String)
* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--exporter_config--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--exporter_config--mobility_sam))
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--exporter_config--monitor))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--exporter_config--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--exporter_config--snmp))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--exporter_config--source))
* `type` (String)
<a id="nestedatt--app_exporter_config--destination_configs--exporter_config--cef"></a>
### Nested Schema for `app_exporter_config.destination_configs.exporter_config.cef`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
<a id="nestedatt--app_exporter_config--destination_configs--exporter_config--destination"></a>
### Nested Schema for `app_exporter_config.destination_configs.exporter_config.destination`

Read-Only:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)
<a id="nestedatt--app_exporter_config--destination_configs--exporter_config--mobility_sam"></a>
### Nested Schema for `app_exporter_config.destination_configs.exporter_config.mobility_sam`

Read-Only:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs--exporter_config--mobility_sam--event_enable))
* `trigger` (String)
<a id="nestedatt--app_exporter_config--destination_configs--exporter_config--mobility_sam--event_enable"></a>
### Nested Schema for `app_exporter_config.destination_configs.exporter_config.mobility_sam.event_enable`

Read-Only:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--app_exporter_config--destination_configs--exporter_config--monitor"></a>
### Nested Schema for `app_exporter_config.destination_configs.exporter_config.monitor`

Read-Only:

* `timeout` (Number) - how often to export in seconds
<a id="nestedatt--app_exporter_config--destination_configs--exporter_config--netflow"></a>
### Nested Schema for `app_exporter_config.destination_configs.exporter_config.netflow`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)
<a id="nestedatt--app_exporter_config--destination_configs--exporter_config--snmp"></a>
### Nested Schema for `app_exporter_config.destination_configs.exporter_config.snmp`

Read-Only:

* `enabled` (Boolean) - snmp reverse lookup enable/disable
<a id="nestedatt--app_exporter_config--destination_configs--exporter_config--source"></a>
### Nested Schema for `app_exporter_config.destination_configs.exporter_config.source`

Read-Only:

* `ip_interface` (String)
<a id="nestedatt--app_filter_config"></a>
### Nested Schema for `app_filter_config`

Read-Only:

* `app_filter_tiered_batch_id` (String) - ID of the appFilterConfig tiered batch
* `egress_traffic_configs` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs))
* `map_alias` (String) - Alias of the appFiltering map
* `sapf_profile` (Attributes) - Session-Aware APF (see [below for nested schema](#nestedatt--app_filter_config--sapf_profile))
* `sapf_profile_alias` (String) - Alias of the SAPF Profile
<a id="nestedatt--app_filter_config--egress_traffic_configs"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs`

Read-Only:

* `drop_app_profile_alias` (String)
* `drop_application_names` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--drop_application_names))
* `pass_app_profile_alias` (String)
* `pass_application_names` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--pass_application_names))
* `priority` (Number)
* `rule_sets` (Attributes) - Flow Map Rule Sets (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--rule_sets))
* `tunnel_aliases` (List of String)
* `tunnel_interface_mappings` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--tunnel_interface_mappings))
<a id="nestedatt--app_filter_config--egress_traffic_configs--drop_application_names"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.drop_application_names`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--drop_application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name
<a id="nestedatt--app_filter_config--egress_traffic_configs--drop_application_names--attributes"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.drop_application_names.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value
<a id="nestedatt--app_filter_config--egress_traffic_configs--pass_application_names"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.pass_application_names`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--pass_application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name
<a id="nestedatt--app_filter_config--egress_traffic_configs--pass_application_names--attributes"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.pass_application_names.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value
<a id="nestedatt--app_filter_config--egress_traffic_configs--rule_sets"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.rule_sets`

Read-Only:

* `aep_id` (Number) - Traffic is forwarded to this AEP if it does not match any drop rules and matches at least 1 pass rule.
* `drop_rules` (Attributes List) - Flow Map Drop Rules. Only used in traffic and exclusion maps. Not relevant for inclusion maps. (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--rule_sets--drop_rules))
* `pass_rules` (Attributes List) - Flow Map Pass Rules. Only used in traffic and inclusion maps. Not relevant for exclusion maps. (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--rule_sets--pass_rules))
* `priority` (Number) - The priority of the rule set. The lower the number the higher the precedence
* `rule_set_id` (String) - The id of the rule set. Each rule set must have a unique ruleSetId
<a id="nestedatt--app_filter_config--egress_traffic_configs--rule_sets--drop_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.rule_sets.drop_rules`

Read-Only:

* `matches` (List of Dynamic) - A set of rule's matching elements. Within a rule, matching elements of the the same type may be repeated. However, their matching positions must be unique.
* `rule_id` (Number) - Rule Id
<a id="nestedatt--app_filter_config--egress_traffic_configs--rule_sets--pass_rules"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.rule_sets.pass_rules`

Read-Only:

* `matches` (List of Dynamic) - A set of rule's matching elements. Within a rule, matching elements of the the same type may be repeated. However, their matching positions must be unique.
* `rule_id` (Number) - Rule Id
<a id="nestedatt--app_filter_config--egress_traffic_configs--tunnel_interface_mappings"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.tunnel_interface_mappings`

Read-Only:

* `iface` (String)
* `tunnel_alias` (String)
<a id="nestedatt--app_filter_config--sapf_profile"></a>
### Nested Schema for `app_filter_config.sapf_profile`

Read-Only:

* `alias` (String)
* `bidi` (Boolean) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
* `buffering` (Attributes) - Session-Aware APF Buffereing settings (see [below for nested schema](#nestedatt--app_filter_config--sapf_profile--buffering))
* `session_fields` (Attributes Set) - A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20 (see [below for nested schema](#nestedatt--app_filter_config--sapf_profile--session_fields))
* `timeout` (Number) - in seconds
<a id="nestedatt--app_filter_config--sapf_profile--buffering"></a>
### Nested Schema for `app_filter_config.sapf_profile.buffering`

Read-Only:

* `buffer_count_before_match` (Number) - Maximum number of packets BSAPF will buffer per session before APF match
* `enabled` (Boolean)
* `protocol` (String) - When buffering is enabled, changing protocol requires reboot before new value takes effect.
<a id="nestedatt--app_filter_config--sapf_profile--session_fields"></a>
### Nested Schema for `app_filter_config.sapf_profile.session_fields`

Read-Only:

* `pos` (Number) - Value of 1 also an alias for 'outer'. Value of 2 also an alias for 'inner'. Not applicable for 'gtpuTeid'. For 'fiveTuple' 'outer' is not supported
* `type` (String)
<a id="nestedatt--dedup"></a>
### Nested Schema for `dedup`

Read-Only:

* `app_config_id` (String)
* `app_instance_name` (String)
* `dedup_tiered_batch_id` (String)
* `enabled` (Boolean)
* `gs_params_name` (String)
<a id="nestedatt--health_state_reasons"></a>
### Nested Schema for `health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--ingress_traffic_configs"></a>
### Nested Schema for `ingress_traffic_configs`

Read-Only:

* `source_selector_aliases` (List of String)
* `src_raw_end_point_interfaces` (List of String)
* `tunnel_aliases` (List of String)
* `tunnel_interface_mappings` (Attributes List) (see [below for nested schema](#nestedatt--ingress_traffic_configs--tunnel_interface_mappings))
<a id="nestedatt--ingress_traffic_configs--tunnel_interface_mappings"></a>
### Nested Schema for `ingress_traffic_configs.tunnel_interface_mappings`

Read-Only:

* `iface` (String)
* `tunnel_alias` (String)
<a id="nestedatt--monitor_solution_config"></a>
### Nested Schema for `monitor_solution_config`

Read-Only:

* `app_instance_config` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--app_instance_config))
* `app_instance_config_id` (String)
* `app_instance_name` (String)
* `app_viz_tiered_batch_id` (String) - ID of the monitoringSolutionConfig tiered batch
* `destination_config` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--destination_config))
* `mgmt_interface` (String) - Pass this value if you would like to use the management interface of the Node to export App Monitoring metadata to FM
<a id="nestedatt--monitor_solution_config--app_instance_config"></a>
### Nested Schema for `monitor_solution_config.app_instance_config`

Read-Only:

* `action` (Boolean)
* `exporter` (String)
<a id="nestedatt--monitor_solution_config--destination_config"></a>
### Nested Schema for `monitor_solution_config.destination_config`

Read-Only:

* `destination_name` (String)
* `exporter_alias` (String)
* `exporter_config` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--destination_config--exporter_config))
<a id="nestedatt--monitor_solution_config--destination_config--exporter_config"></a>
### Nested Schema for `monitor_solution_config.destination_config.exporter_config`

Read-Only:

* `alias` (String)
* `destination` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--destination_config--exporter_config--destination))
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--destination_config--exporter_config--monitor))
* `type` (String)
<a id="nestedatt--monitor_solution_config--destination_config--exporter_config--destination"></a>
### Nested Schema for `monitor_solution_config.destination_config.exporter_config.destination`

Read-Only:

* `ipv4_address` (String)
* `ipv6_address` (String)
<a id="nestedatt--monitor_solution_config--destination_config--exporter_config--monitor"></a>
### Nested Schema for `monitor_solution_config.destination_config.exporter_config.monitor`

Read-Only:

* `timeout` (Number)

