---
page_title: "gigavuecore_load_mobility_sam_node_configs Data Source - gigavuecore"
subcategory: ""
description: |-
  Load mobility sam node configurations by alias
---

# gigavuecore_load_mobility_sam_node_configs Data Source

Load mobility sam node configurations by alias

## Example Usage

```terraform
data "gigavuecore_load_mobility_sam_node_configs" "example" {
  samnode_alias  = null
  solution_alias = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `samnode_alias` (String, required) - Alias of the sam node
* `solution_alias` (String, required) - Alias of the configured solution

### Attributes

In addition to all arguments above, the following attributes are exported:

* `alias` (String, computed) - Alias of the Sam node
* `engine_level_app_vz_configs` (Attributes List, computed) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs))
* `ip_interface_config` (Attributes, computed) (see [below for nested schema](#nestedatt--ip_interface_config))
* `location` (Attributes, computed) - Location of the engine port (see [below for nested schema](#nestedatt--location))
* `site_name` (String, computed) - Name of the site where the node is deployed

<a id="nestedatt--engine_level_app_vz_configs"></a>
### Nested Schema for `engine_level_app_vz_configs`

Read-Only:

* `app_vz_config` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config))
* `engine_port` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config`

Read-Only:

* `app_export_config` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config))
* `app_filter_config` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config))
* `associated_monitor_solution_alias` (String) - monitor object alias
* `cluster_id` (String) - cluster id for which solution is getting created
* `config_status` (String) - configuration status
* `delete_monitor_sol` (Boolean)
* `egress_map_aliases_to_delete` (List of String)
* `exporter_aliases_to_delete` (List of String)
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--health_state_reasons))
* `ingress_map_aliases_to_delete` (List of String)
* `ingress_traffic_configs` (Attributes List) - ingress traffic configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs))
* `monitor_solution_config` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config))
* `solution_alias` (String) - user defined application visibility solution alias
* `solution_desc` (String) - user defined application visibility solution description
* `solution_status` (String) - solution status
* `solution_type` (String) - user intent solution type
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config`

Read-Only:

* `cache_config` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config))
* `cache_config_alias` (String)
* `destination_configs` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs))
* `gsop_alias` (String)
* `gsop_config` (Attributes) - GigaSMART Operation (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config`

Read-Only:

* `advance_hash` (Boolean) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
* `alias` (String)
* `description` (String)
* `dpi_inject_limit` (Number)
* `event` (String)
* `exporters` (List of String) - alias of metadata exporters to attach this cache
* `flow_behavior` (String) - direction for flow identification
* `match` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match))
* `multi_collect` (Boolean) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
* `network_profiles` (List of String) - alias of metadata network profiles to attach this cache
* `observation_domain_id` (Number)
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--sampling))
* `size` (Attributes) - size of the flows (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--size))
* `timeout` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--timeout))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match`

Read-Only:

* `datalink` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--datalink))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--transport))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--datalink"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.datalink`

Read-Only:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--interface"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.interface`

Read-Only:

* `in_name_width` (Number)
* `in_physical_width` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ip"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ip`

Read-Only:

* `version` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv4`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4--destination"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv4.destination`

Read-Only:

* `prefix_min_mask` (String) - ipv4 destination prefix minimum-mask - netmask or mask length
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4--fragmentation"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv4.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4--section"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv4.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv4--source"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv4.source`

Read-Only:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv6`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--destination"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv6.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--fragmentation"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv6.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--length"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv6.length`

Read-Only:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--section"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv6.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--ipv6--source"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.ipv6.source`

Read-Only:

* `prefix_min_mask` (String) - ipv6 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--transport"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.transport`

Read-Only:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--transport--udp))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--transport--icmp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.transport.icmp`

Read-Only:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--transport--tcp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.transport.tcp`

Read-Only:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--match--transport--udp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.match.transport.udp`

Read-Only:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--sampling"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.sampling`

Read-Only:

* `mode` (String)
* `single_sampling_rate` (Number) - Packet interval window size. Valid values: 10-16000 (in packets)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--size"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.size`

Read-Only:

* `flows` (Number) - size of flows in millions
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--cache_config--timeout"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.cache_config.timeout`

Read-Only:

* `idle` (Number) - idle timeout in seconds. max value 7days. default 30 min
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs`

Read-Only:

* `application_names` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--application_names))
* `destination_name` (String)
* `export_ip_interface` (String)
* `export_meta_app_profile` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile))
* `export_meta_app_profile_alias` (String)
* `exporter_alias` (String)
* `exporter_config` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--application_names"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.application_names`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--application_names--attributes"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.application_names.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile`

Read-Only:

* `alias` (String) - application profile alias
* `application_id` (Boolean) - only valid with 'export' type
* `applications` (Attributes List) - application and attributes. (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--counter))
* `datalink` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--datalink))
* `description` (String)
* `flow` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--gtpu))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--transport))
* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--applications"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.applications`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--applications--attributes"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.applications.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--counter"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.counter`

Read-Only:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `inner_byte` (Boolean)
* `inner_byte_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--datalink"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.datalink`

Read-Only:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--flow"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.flow`

Read-Only:

* `end_reason` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--gtpu"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.gtpu`

Read-Only:

* `qfi` (Boolean)
* `teid` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--interface"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.interface`

Read-Only:

* `in_name_width` (Number)
* `in_physical_width` (Number)
* `out_physical_width` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ip"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ip`

Read-Only:

* `version` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv4`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4--destination"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv4.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4--fragmentation"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv4.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4--section"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv4.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv4--source"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv4.source`

Read-Only:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv6`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--destination"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv6.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--fragmentation"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv6.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--length"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv6.length`

Read-Only:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--section"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv6.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--ipv6--source"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.ipv6.source`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--outer_ipv4"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.outer_ipv4`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--outer_ipv6"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.outer_ipv6`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--timestamp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.timestamp`

Read-Only:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
* `sys_up_time_first` (Boolean)
* `sys_up_time_last` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--transport"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.transport`

Read-Only:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--transport--udp))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--transport--icmp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.transport.icmp`

Read-Only:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--transport--tcp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.transport.tcp`

Read-Only:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--export_meta_app_profile--transport--udp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.export_meta_app_profile.transport.udp`

Read-Only:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.exporter_config`

Read-Only:

* `alias` (String)
* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--cef))
* `description` (String)
* `destination` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--mobility_sam))
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--monitor))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--snmp))
* `source` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--source))
* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--cef"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.exporter_config.cef`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--destination"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.exporter_config.destination`

Read-Only:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--mobility_sam"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.exporter_config.mobility_sam`

Read-Only:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--mobility_sam--event_enable))
* `trigger` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--mobility_sam--event_enable"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.exporter_config.mobility_sam.event_enable`

Read-Only:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--monitor"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.exporter_config.monitor`

Read-Only:

* `timeout` (Number) - how often to export in seconds
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--netflow"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.exporter_config.netflow`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--snmp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.exporter_config.snmp`

Read-Only:

* `enabled` (Boolean) - snmp reverse lookup enable/disable
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--destination_configs--exporter_config--source"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.destination_configs.exporter_config.source`

Read-Only:

* `ip_interface` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config`

Read-Only:

* `alias` (String)
* `cluster_id` (String) - id of the defining cluster
* `gs_apps` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps))
* `gs_group` (String) - Alias of referenced managing GsGroup
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--health_state_reasons))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps`

Read-Only:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--apf"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--dedup"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.dedup`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--diameter_whitelist"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.diameter_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--flow_filter"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.flow_filter`

Read-Only:

* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--flow_sampling"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.flow_sampling`

Read-Only:

* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_header_add"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.gseries_header_add`

Read-Only:

* `types` (Set of String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_header_remove"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.gseries_header_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_load_balance"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.gseries_load_balance`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_load_balance--variable_offset))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.gseries_load_balance.fixed_offset`

Read-Only:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_load_balance--variable_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.gseries_load_balance.variable_offset`

Read-Only:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_pattern_match"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.gseries_pattern_match`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_pattern_match--variable_offset))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.gseries_pattern_match.fixed_offset`

Read-Only:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.gseries_pattern_match.variable_offset`

Read-Only:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--gtp_whitelist"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.gtp_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--header_add"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.header_add`

Read-Only:

* `vlan` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--header_remove"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.header_remove`

Read-Only:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--icap"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.icap`

Read-Only:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--inline_ssl"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.inline_ssl`

Read-Only:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.load_balance`

Read-Only:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance--stateless))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance--enhanced"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.load_balance.enhanced`

Read-Only:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance--stateful"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.load_balance.stateful`

Read-Only:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.load_balance.stateful.diameter_key_multi_hash_type`

Read-Only:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
* `key` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--load_balance--stateless"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.load_balance.stateless`

Read-Only:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
* `hash_fields` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--masking"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.masking`

Read-Only:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
* `protocol` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--metadata_export"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.metadata_export`

Read-Only:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--netflow"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.netflow`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--sa_apf"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.sa_apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--sip_whitelist"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.sip_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--slicing"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.slicing`

Read-Only:

* `enhanced` (String) - enhanced-slicing apps alias
* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--ssl_decrypt"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.ssl_decrypt`

Read-Only:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--trailer_add"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.trailer_add`

Read-Only:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--trailer_remove"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.trailer_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_decap"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.tunnel_decap`

Read-Only:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_decap--tls_pcapng))
* `type` (String)
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_decap--vxlan))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_decap--custom"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.tunnel_decap.custom`

Read-Only:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.tunnel_decap.tls_pcapng`

Read-Only:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_decap--vxlan"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.tunnel_decap.vxlan`

Read-Only:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.tunnel_encap`

Read-Only:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap--tls_pcapng))
* `type` (String)
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap--vxlan_config))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap--gmip_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.tunnel_encap.gmip_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `src_port` (Number)
* `ttl` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.tunnel_encap.l2_gre_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `key` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.tunnel_encap.tls_pcapng`

Read-Only:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--gs_apps--tunnel_encap--vxlan_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.gs_apps.tunnel_encap.vxlan_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `ttl` (Number)
* `vni` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_export_config--gsop_config--health_state_reasons"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_export_config.gsop_config.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config`

Read-Only:

* `egress_traffic_config` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config))
* `sapf_profile` (Attributes) - Session-Aware APF (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--sapf_profile))
* `sapf_profile_alias` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config`

Read-Only:

* `drop_app_profile_alias` (String)
* `egress_traffic_map` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map))
* `gs_apps` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps))
* `gsop_alias` (String)
* `gsop_config` (Attributes) - GigaSMART Operation (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config))
* `map_alias` (String)
* `pass_app_profile_alias` (String)
* `priority` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map`

Read-Only:

* `alias` (String) - unique map alias
* `ap_rules` (Attributes) - pass and drop application profile rules (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules))
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `dst_ports` (List of String) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
* `egress_gigastream` (List of String) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Boolean) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String) - tunnel alias
* `flex_inline` (Attributes) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline))
* `flex_inline_failover` (String) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Attributes) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules))
* `fstype` (Attributes) - Type of Mobility Flowsampling & properties (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--fstype))
* `gs_rules` (Attributes) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules))
* `gsop` (String) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--ip_rewrite))
* `mod_time` (Number) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Boolean) - enabled when the dstPort is null
* `order` (Number) - relative order within per-source port map chain
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rewrite))
* `roles` (Attributes) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--roles))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules))
* `rx_cluster_ports` (List of String)
* `src_ports` (List of String) - list of the 'from' ports. Only port number is supported
* `sub_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List of String)
* `type` (String) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--vlan_tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.ap_rules`

Read-Only:

* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules--drop_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.ap_rules.drop_rules`

Read-Only:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--ap_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.ap_rules.pass_rules`

Read-Only:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline`

Read-Only:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--a_to_b"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.a_to_b`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--b_to_a"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.b_to_a`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--oob_copy"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.oob_copy`

Read-Only:

* `direction` (String)
* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists
* `tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--oob_copy--tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--oob_copy--tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.oob_copy.tag`

Read-Only:

* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flex_inline--tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flex_inline.tag`

Read-Only:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `type` (String)
* `vlan_id` (Number) - only applicable when type is 'vlan'
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--drop_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules.drop_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--drop_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules.drop_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules.pass_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--pass_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_rules.pass_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `priority` (Number)
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample5_g_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_diameter_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_diameter_rules.pass_rules`

Read-Only:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_diameter_rules.pass_rules.diameter`

Read-Only:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_overlap_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules`

Read-Only:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules.sip`

Read-Only:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--id_range))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_sample_sip_rules.pass_rules.sip.id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist5_g_overlap_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist5_g_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist5_g_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--sip))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_overlap_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--sip))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.flow_whitelist_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--fstype"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.fstype`

Read-Only:

* `offset` (Number) - Offset for Mobility Rotational Flowsampling
* `timer` (Number) - Timer for Mobility Rotational Flowsampling in minutes
* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.gs_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules--drop_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.gs_rules.drop_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--gs_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.gs_rules.pass_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--ip_rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--roles"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.roles`

Read-Only:

* `editors` (List of String)
* `listeners` (List of String)
* `owners` (List of String)
* `viewers` (List of String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rules.drop_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--vlan_tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rules.drop_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rules.drop_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rules.drop_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rules.pass_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--vlan_tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rules.pass_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rules.pass_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.rules.pass_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--egress_traffic_map--vlan_tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.egress_traffic_map.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps`

Read-Only:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--apf"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--dedup"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.dedup`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--diameter_whitelist"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.diameter_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--flow_filter"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.flow_filter`

Read-Only:

* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--flow_sampling"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.flow_sampling`

Read-Only:

* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_header_add"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.gseries_header_add`

Read-Only:

* `types` (Set of String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_header_remove"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.gseries_header_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.gseries_load_balance`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance--variable_offset))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.gseries_load_balance.fixed_offset`

Read-Only:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_load_balance--variable_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.gseries_load_balance.variable_offset`

Read-Only:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.gseries_pattern_match`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match--variable_offset))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.gseries_pattern_match.fixed_offset`

Read-Only:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.gseries_pattern_match.variable_offset`

Read-Only:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--gtp_whitelist"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.gtp_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--header_add"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.header_add`

Read-Only:

* `vlan` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--header_remove"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.header_remove`

Read-Only:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--icap"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.icap`

Read-Only:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--inline_ssl"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.inline_ssl`

Read-Only:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.load_balance`

Read-Only:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateless))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance--enhanced"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.load_balance.enhanced`

Read-Only:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateful"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.load_balance.stateful`

Read-Only:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.load_balance.stateful.diameter_key_multi_hash_type`

Read-Only:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
* `key` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--load_balance--stateless"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.load_balance.stateless`

Read-Only:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
* `hash_fields` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--masking"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.masking`

Read-Only:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
* `protocol` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--metadata_export"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.metadata_export`

Read-Only:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--netflow"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.netflow`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--sa_apf"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.sa_apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--sip_whitelist"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.sip_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--slicing"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.slicing`

Read-Only:

* `enhanced` (String) - enhanced-slicing apps alias
* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--ssl_decrypt"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.ssl_decrypt`

Read-Only:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--trailer_add"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.trailer_add`

Read-Only:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--trailer_remove"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.trailer_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.tunnel_decap`

Read-Only:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--tls_pcapng))
* `type` (String)
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--vxlan))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--custom"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.tunnel_decap.custom`

Read-Only:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.tunnel_decap.tls_pcapng`

Read-Only:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_decap--vxlan"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.tunnel_decap.vxlan`

Read-Only:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.tunnel_encap`

Read-Only:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--tls_pcapng))
* `type` (String)
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--vxlan_config))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--gmip_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.tunnel_encap.gmip_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `src_port` (Number)
* `ttl` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.tunnel_encap.l2_gre_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `key` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.tunnel_encap.tls_pcapng`

Read-Only:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gs_apps--tunnel_encap--vxlan_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gs_apps.tunnel_encap.vxlan_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `ttl` (Number)
* `vni` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config`

Read-Only:

* `alias` (String)
* `cluster_id` (String) - id of the defining cluster
* `gs_apps` (Attributes) - GigaSMART Applications for a GSOP. At least one GsApp must be defined for the config to be valid (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps))
* `gs_group` (String) - Alias of referenced managing GsGroup
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--health_state_reasons))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps`

Read-Only:

* `apf` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--apf))
* `dedup` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--dedup))
* `diameter_whitelist` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--diameter_whitelist))
* `flow_filter` (Attributes) - GigaSMART 'Flow Filter' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--flow_filter))
* `flow_sampling` (Attributes) - GigaSMART 'Flow Sampling' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--flow_sampling))
* `gseries_header_add` (Attributes) - Only applicable for G-series (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_header_add))
* `gseries_header_remove` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_header_remove))
* `gseries_load_balance` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance))
* `gseries_pattern_match` (Attributes) - Only applicable for G-series per-rule GSOP (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match))
* `gtp_whitelist` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gtp_whitelist))
* `header_add` (Attributes) - GigaSMART 'Add Header' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--header_add))
* `header_remove` (Attributes) - GigaSMART 'Remove Header' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--header_remove))
* `icap` (Attributes) - GigaSMART ICAP Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--icap))
* `inline_ssl` (Attributes) - GigaSMART Inline SSL Profile Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--inline_ssl))
* `load_balance` (Attributes) - GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance))
* `masking` (Attributes) - GigaSMART 'Masking' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--masking))
* `metadata_export` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--metadata_export))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--netflow))
* `sa_apf` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--sa_apf))
* `sip_whitelist` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--sip_whitelist))
* `slicing` (Attributes) - GigaSMART 'Slicing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--slicing))
* `ssl_decrypt` (Attributes) - GigaSMART 'SSL Decrypt' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--ssl_decrypt))
* `trailer_add` (Attributes) - GigaSMART 'Add Trailer' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--trailer_add))
* `trailer_remove` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--trailer_remove))
* `tunnel_decap` (Attributes) - GigaSMART 'Decapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap))
* `tunnel_encap` (Attributes) - GigaSMART 'Encapsulate Tunnel' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--apf"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--dedup"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.dedup`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--diameter_whitelist"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.diameter_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--flow_filter"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.flow_filter`

Read-Only:

* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--flow_sampling"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.flow_sampling`

Read-Only:

* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_header_add"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_header_add`

Read-Only:

* `types` (Set of String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_header_remove"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_header_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_load_balance`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Load Balancing config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance--variable_offset))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance--fixed_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_load_balance.fixed_offset`

Read-Only:

* `hash` (String)
* `length` (Number)
* `offset` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_load_balance--variable_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_load_balance.variable_offset`

Read-Only:

* `end_delim` (String)
* `hash` (String)
* `start_delim` (String)
* `start_field` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_pattern_match`

Read-Only:

* `fixed_offset` (Attributes) - Per-rule Fixed Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match--fixed_offset))
* `variable_offset` (Attributes) - Per-rule Variable Offset Pattern Match config for G-seres devices (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match--variable_offset))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match--fixed_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_pattern_match.fixed_offset`

Read-Only:

* `length` (Number)
* `offset` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gseries_pattern_match--variable_offset"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.gseries_pattern_match.variable_offset`

Read-Only:

* `end_delim` (String)
* `start_delim` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--gtp_whitelist"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.gtp_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--header_add"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.header_add`

Read-Only:

* `vlan` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--header_remove"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.header_remove`

Read-Only:

* `ah1` (String) - only valid and required for 'generic'. First anchor header, from which the header to be stripped is occurred.
* `ah2` (String) - only valid and required for 'generic'. next anchor header.
* `custom_len` (Number) - only valid for 'generic'. length of unknown header.
* `erspan_flow_id` (Number) - only applicable for 'erspan', specifies the flow id to strip. Value of 0 will strip all flow ids
* `fp_dst_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit destination switch id
* `fp_src_switch_id` (Number) - Only valid and required for 'fabricPath', 12 bit source switch id
* `header_count` (Number) - only valid for 'generic'. Number of headers to be stripped.
* `offset` (String) - only valid and required for 'generic'.'start': strip from start of ah1, 'end': strip from end of ah1 or any other offset in the range of length of ah1, 'offsetRangeValue' is required if offset is 'offsetRange'.
* `offset_range_value` (Number) - only valid and required when offset is 'offsetRange', integer within range of size of header
* `protocol` (String) - 'gre' and 'fabricPath' are only applicable for H-series
* `timestamp_format` (String) - Timestamp format. Only valid and required for 'fm6000Ts'
* `vlan_header` (String) - Only valid when protocol is 'vlan'. Specifies the target vlan header to strip. Defaults to 'outer'
* `vxlan_id` (Number) - 24-bit value. Only valid when protocol is 'vxlan'. Specifies the vxlan id to strip. Value of '0' will strip all vxlan ids
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--icap"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.icap`

Read-Only:

* `icap_profile` (String) - Alias of referenced ICAP Profile
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--inline_ssl"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.inline_ssl`

Read-Only:

* `inline_ssl_profile` (String) - Alias of referenced Inline SSL Profile
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance`

Read-Only:

* `enhanced` (Attributes) - Private class. Enhanced part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--enhanced))
* `stateful` (Attributes) - Private class. Stateful part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateful))
* `stateless` (Attributes) - Private class. Stateless part of the GigaSMART 'Load Balancing' Application Configuration (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateless))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--enhanced"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance.enhanced`

Read-Only:

* `elb_alias` (String) - elb app alias
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateful"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance.stateful`

Read-Only:

* `app_type` (String) - 'sapf' option is added in release H 4.3.01. 'sip' and 'tunnel' added in release H 5.1. 'diameter' is added in release H 5.5
* `diameter_key_hash_type` (String) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyHash'
* `diameter_key_multi_hash_type` (Attributes List) - required when 'appType' == 'diameter' and 'lbType' == 'diameterKeyMultiHash' (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type))
* `gtp_key_hash_type` (String) - required when 'appType' == 'gtp' and 'lbType' == 'gtpKeyHash'. ignored otherwise
* `lb_type` (String) - 'gtpKeyHash' is only applicable when 'appType' == 'gtp'. 'sipKeyHash' is only applicable when 'appType' == 'sip'
* `sip_key_hash_type` (String) - required when 'appType' == 'sip' and 'lbType' == 'sipKeyHash'. ignored otherwise
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateful--diameter_key_multi_hash_type"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance.stateful.diameter_key_multi_hash_type`

Read-Only:

* `avp_codevalue` (Number) - required when 'key' == 'avpCode'
* `key` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--load_balance--stateless"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.load_balance.stateless`

Read-Only:

* `field_location` (String) - Ignored when 'hashFields' == 'gtpuTeid'. Required otherwise
* `hash_fields` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--masking"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.masking`

Read-Only:

* `content_type` (String) - content type that will trigger masking, only valid and required for protocol 'sip'
* `length` (Number) - max length is 9600 on H-series; 9000 on G-series, not applicable for protocol 'sip' valid and required otherwise
* `offset` (Number) - min offset is 0 for 'none'; 1 for protocol-referenced offsets, not applicable for protocol 'sip' valid and required otherwise
* `pattern` (String) - 1-byte hex mask, not applicable for protocol 'sip' valid and required otherwise
* `protocol` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--metadata_export"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.metadata_export`

Read-Only:

* `cache` (String) - metadata cache alias. cache should have exporters defined, and the exporters should have applicationProfiles defined
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--netflow"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.netflow`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--sa_apf"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.sa_apf`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--sip_whitelist"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.sip_whitelist`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--slicing"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.slicing`

Read-Only:

* `enhanced` (String) - enhanced-slicing apps alias
* `offset` (Number) - min offset is 64 for 'none'; 4 for protocol-referenced offsets. required property till H 5.6
* `protocol` (String) - required property till H 5.6
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--ssl_decrypt"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.ssl_decrypt`

Read-Only:

* `in_port` (Number) - Port number of 0 represents 'any' port
* `out_port` (Number) - Port number of 0 represents 'auto' port
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--trailer_add"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.trailer_add`

Read-Only:

* `types` (Set of String) - 'crc' is not applicable for G-series
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--trailer_remove"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.trailer_remove`

Read-Only:

* `enabled` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_decap`

Read-Only:

* `custom` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--custom))
* `erspan_flow_id` (Number) - only applicable for 'erspan', A Flow ID of 0 decapsulates all ERSPAN tunnel traffic regardless of Flow ID
* `gmip_port` (Number) - only applicable for 'gmip', in which case it is required. Specifies the UDP port on which the Tunnel Network port on the receiving GigaVUE H Series is listening. Must match the configuration of the portdst configured on the sending end of the tunnel
* `l2_gre_key` (Number) - only applicable for 'l2gre', in which case it is required.
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--tls_pcapng))
* `type` (String)
* `vxlan` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--vxlan))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--custom"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_decap.custom`

Read-Only:

* `port_dst` (Number) - when specified as 0, no validation will be done in the packet.
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--tls_pcapng"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_decap.tls_pcapng`

Read-Only:

* `decap_key` (String)
* `listener` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_decap--vxlan"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_decap.vxlan`

Read-Only:

* `port_dst` (Number)
* `port_src` (Number) - when specified as 0, no validation will be done in the packet.
* `vni` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap`

Read-Only:

* `gmip_config` (Attributes) - Configuration for GMIP Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--gmip_config))
* `l2_gre_config` (Attributes) - Configuration for l2GRE Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--l2_gre_config))
* `tls_pcapng` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--tls_pcapng))
* `type` (String)
* `vxlan_config` (Attributes) - Configuration for Vxlan Tunnel Encapsulate GigaSmaprt App (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--vxlan_config))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--gmip_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap.gmip_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `flow_label` (Number)
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `src_port` (Number)
* `ttl` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--l2_gre_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap.l2_gre_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - ip destination, mutually exclusive with 'pgDst'. IPv4 or IPv6.
* `flow_label` (Number)
* `key` (Number)
* `pg_dst` (String) - port group destination alias, mutually exclusive with 'dstIp'
* `prec` (Number) - decimal Precedence value from 0-7 to be used in the ToS byte of the outer headers of tunneled packets
* `session_field` (String) - required with stateful loadBalance when 'appType' is 'tunnel'
* `session_pos` (String) - required if 'sessionField' is specified
* `ttl` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--tls_pcapng"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap.tls_pcapng`

Read-Only:

* `exporter` (String)
* `exporter_group` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--gs_apps--tunnel_encap--vxlan_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.gs_apps.tunnel_encap.vxlan_config`

Read-Only:

* `dscp` (Number)
* `dst_ip` (String) - IP Destination. IPv4 or IPv6.
* `dst_port` (Number)
* `src_port` (Number)
* `ttl` (Number)
* `vni` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--egress_traffic_config--gsop_config--health_state_reasons"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.egress_traffic_config.gsop_config.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--sapf_profile"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.sapf_profile`

Read-Only:

* `alias` (String)
* `bidi` (Boolean) - include reverse traffic for the selected sessionFields. Not applicable for 'mplsLabel', 'gtpuTeid', 'vlanId'
* `buffering` (Attributes) - Session-Aware APF Buffereing settings (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--sapf_profile--buffering))
* `cluster_id` (String) - id of the defining cluster
* `packet_count` (Number) - For each session, forward this number of packets. Valid range is 2-100; 0 disables packetCount.
* `session_fields` (Attributes Set) - A sessionField cannot contain overlapping session attribute, position pairs (e.g. ipv4-5tuple pos 1 and ipv4-src pos 1 is not allowed). Max number of session fields is 20 (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--sapf_profile--session_fields))
* `timeout` (Number) - in seconds
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--sapf_profile--buffering"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.sapf_profile.buffering`

Read-Only:

* `buffer_count_before_match` (Number) - Maximum number of packets BSAPF will buffer per session before APF match
* `enabled` (Boolean)
* `protocol` (String) - When buffering is enabled, changing protocol requires reboot before new value takes effect.
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--app_filter_config--sapf_profile--session_fields"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.app_filter_config.sapf_profile.session_fields`

Read-Only:

* `pos` (Number) - Value of 1 also an alias for 'outer'. Value of 2 also an alias for 'inner'. Not applicable for 'gtpuTeid'. For 'fiveTuple' 'outer' is not supported
* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--health_state_reasons"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs`

Read-Only:

* `ingress_traffic_map` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map))
* `map_alias` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map`

Read-Only:

* `alias` (String) - unique map alias
* `ap_rules` (Attributes) - pass and drop application profile rules (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--ap_rules))
* `cluster_id` (String) - id of the defining cluster
* `comment` (String)
* `dst_ports` (List of String) - List of the 'to' ports. Invalid if 'nullDstPort' is true'. Only port number is supported
* `egress_gigastream` (List of String) - Applicable to fabric map only. It is to specify gigastream(s) to use as egress when going across cluster. The format of array member is clusterID:gigastreamAlias
* `enable` (Boolean) - enable/disable map, applicable only to first level maps
* `encap_tunnel` (String) - tunnel alias
* `flex_inline` (Attributes) - When 'oobCopy' is configured at least one of 'aToB' or 'bToA' must be configured (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline))
* `flex_inline_failover` (String) - only valid for flexInline maps
* `flex_inline_vlan_id` (Number) - VLAN ID carried in the VLAN tag of packets coming from the inline-network port(s). valid and applicable only with 'flexInline' map type, 'srcPorts' should be either inline-network or ib-pathway or vport
* `flow_rules` (Attributes) - Map Flow Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules))
* `flow_sample5_g_overlap_rules` (Attributes) - Map Flow Sample 5g Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules))
* `flow_sample5_g_rules` (Attributes) - Map Flow Sample 5g Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules))
* `flow_sample_diameter_rules` (Attributes) - Map Flow Sample Diameter Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules))
* `flow_sample_overlap_rules` (Attributes) - Map Flow Sample Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules))
* `flow_sample_rules` (Attributes) - Map Flow Sample Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules))
* `flow_sample_sip_rules` (Attributes) - Map Flow Sample Sip Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules))
* `flow_whitelist5_g_overlap_rules` (Attributes) - Map Flow Whitelist 5g Overlap Rule match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist5_g_overlap_rules))
* `flow_whitelist5_g_rules` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist5_g_rules))
* `flow_whitelist_overlap_rules` (Attributes) - Map Flow Whitelist Overlap Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules))
* `flow_whitelist_rules` (Attributes) - Map Flow Whitelist Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules))
* `fstype` (Attributes) - Type of Mobility Flowsampling & properties (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--fstype))
* `gs_rules` (Attributes) - Map GigaSMART Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--gs_rules))
* `gsop` (String) - Alias of referenced GSOP. Applicable to 'regular' and 'secondLevel' map types
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--health_state_reasons))
* `inline_traffic_path` (String) - Only applicable for 'inline' map types, in which case defaults to 'normal'. When set to 'bypass', the 'dstPorts' must be empty
* `inline_traffic_type` (String) - Only applicable for 'inline/passAll' map types, in which case defaults to 'symmetric'. For 'asymmetric' maps, 'srcPort' must be of type 'inline-net'
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--ip_rewrite))
* `mod_time` (Number) - Last modification time of the map in milliseconds since the epoch
* `null_dst_port` (Boolean) - enabled when the dstPort is null
* `order` (Number) - relative order within per-source port map chain
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rewrite))
* `roles` (Attributes) - Map Roles Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--roles))
* `rule_matching` (String) - If set to 'blacklist', packet are passed when no 'drop' rules are matched. This field is only valid for 'regular/byRule' map types. (maps into 'no-rule-match' CLI command)
* `rules` (Attributes) - Map Rules Container. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules))
* `rx_cluster_ports` (List of String)
* `src_ports` (List of String) - list of the 'from' ports. Only port number is supported
* `sub_type` (String) - 'byRule' is applicable to all map types; 'collector' is applicable to 'regular', 'inline', 'flexInline' and 'secondLevel' maps; 'passAll' is applicable to 'regular' and 'inline' maps; 'flowFilter', 'flowSample', 'flowWhitelist', 'flowSampleSip', 'flowWhitelistSip', 'flowSampleDiameter', 'flowWhitelistDiameter', 'flowSampleOverlap', 'flowWhitelistOverlap', 'flowSample5g','flowWhitelist5g','flowSample5gOverlap' and 'flowWhitelist5gOverlap' are applicable to 'secondLevel' maps
* `traffic_type` (String) - Only applicable for 'firstLevel/byRule' map types, in which case defaults to 'user'
* `tx_cluster_ports` (List of String)
* `type` (String) - 'regular' maps are from network/hybrid ports to tool/hybrid/gigastream; 'inline' maps are from inline ports to inline/tool/hybrid/gigastream; 'firstLevel' are from network/hybrid ports to vPorts/tool/hybrid/gigastream; 'secondLevel' maps are from vPorts to tool/hybrid/gigastream; 'inlineFirstLevel is from inline network to vport; 'inlineSecondLevel' is from vport to inline tool; 'flexInline' is from inline network to tools; 'transitLevel' is from vport to vport
* `updated_time` (Number) - Last Updated time of the fabric map
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--vlan_tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--ap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.ap_rules`

Read-Only:

* `drop_rules` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--ap_rules--drop_rules))
* `pass_rules` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--ap_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--ap_rules--drop_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.ap_rules.drop_rules`

Read-Only:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--ap_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.ap_rules.pass_rules`

Read-Only:

* `application_profile` (String) - application profile alias
* `rule_id` (Number) - application profile rule Id, should not have same id as gsRules
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flex_inline`

Read-Only:

* `a_to_b` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--a_to_b))
* `b_to_a` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--b_to_a))
* `oob_copy` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--oob_copy))
* `svt_mode` (Boolean)
* `svt_tag` (Number) - only applicable when svtMode is enabled
* `tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--a_to_b"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flex_inline.a_to_b`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--b_to_a"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flex_inline.b_to_a`

Read-Only:

* `ib_pathway` (String) - ibPathway alias. Only applicable when type is 'ibPathway'
* `tools` (List of String) - ordered list of inline tools or vports. Only applicable when 'type' is 'tools'
* `type` (String) - when set to 'tools', the list of the processing inline tools have to be provided. when set to 'ibPathway', ibPathway alias should be provided
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--oob_copy"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flex_inline.oob_copy`

Read-Only:

* `direction` (String)
* `dst_ports` (List of String) - list of destination tool ports
* `src_ports` (List of String) - inline network or an item from a-to-b and b-to-a lists
* `tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--oob_copy--tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--oob_copy--tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flex_inline.oob_copy.tag`

Read-Only:

* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flex_inline--tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flex_inline.tag`

Read-Only:

* `tag_protocol_id` (String) - When tool VLAN tag is added , this protocol Id will be added which egress out the traffic
* `type` (String)
* `vlan_id` (Number) - only applicable when type is 'vlan'
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules--drop_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_rules.drop_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules--drop_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules--drop_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_rules.drop_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_rules.pass_rules`

Read-Only:

* `gtp` (Attributes) - Map Flow Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules--pass_rules--gtp))
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_rules.pass_rules.gtp`

Read-Only:

* `imei` (String) - mutually exclusive with 'imsi' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - mutually exclusive with 'imei' and 'msisdn'. If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - mutually exclusive with 'imsi' and 'imei'. If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Overlap Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_rules.pass_rules`

Read-Only:

* `comment` (String)
* `flow5_g` (Attributes) - Map Flow Sample Rule 5g match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules--pass_rules--flow5_g))
* `percentage` (Number)
* `priority` (Number)
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample5_g_rules--pass_rules--flow5_g"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample5_g_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `gpsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nsiid` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `pei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `supi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_diameter_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_diameter_rules.pass_rules`

Read-Only:

* `diameter` (Attributes) - Map Flow Sample Diameter Rule Definition (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules--pass_rules--diameter))
* `interface` (String) - interface type
* `percentage` (Number)
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_diameter_rules--pass_rules--diameter"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_diameter_rules.pass_rules.diameter`

Read-Only:

* `user_name` (String) - If '\*' is added at the end of the value, it is treated as prefix
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_overlap_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_overlap_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_rules.pass_rules`

Read-Only:

* `comment` (String)
* `gtp` (Attributes) - Map Flow Sample Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules--pass_rules--gtp))
* `percentage` (Number)
* `periodic_recalc` (Boolean) - Enable Periodic Recalc for rotational sampling. Map look up in the data path based on this flag
* `priority` (Number) - maximum value is equal to the number of rules upon completion of the request
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `eci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imei` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `imsi` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `interface` (String) - interface type. Mutually exclusive with version
* `msisdn` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `nas_5_qi` (String) - 5G QoS Indicator. Valid 5qi value <1 - 255>
* `nci` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `plmn_id` (String) - If '\*' is added at the end of the MNC value, it is treated as prefix
* `qci` (Number) - QoS Class Indicator
* `snssai` (String) - If '\*' is added at the end of the SD value, it is treated as prefix
* `tac` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `tac_5_g` (String) - If '\*' is added at the end of the value, it is treated as prefix
* `version` (String) - mutually exclusive with interface
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules`

Read-Only:

* `percentage` (Number)
* `rule_id` (Number)
* `sip` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules.sip`

Read-Only:

* `callee_id` (String) - sip callee id
* `callee_id_range` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range))
* `caller_id` (String) - sip caller id
* `caller_id_range` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range))
* `id_range` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--id_range))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--callee_id_range"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules.sip.callee_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--caller_id_range"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules.sip.caller_id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_sample_sip_rules--pass_rules--sip--id_range"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_sample_sip_rules.pass_rules.sip.id_range`

Read-Only:

* `max_value` (String)
* `value` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist5_g_overlap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist5_g_overlap_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist5_g_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist5_g_rules`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--sip))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--flow5_g"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_overlap_rules--pass_rules--sip"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_overlap_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules`

Read-Only:

* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules.pass_rules`

Read-Only:

* `flow5_g` (Attributes) - Map Flow Whitelist 5g Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--flow5_g))
* `gtp` (Attributes) - Map Flow Whitelist Rule GTP match Definition. Private class (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--gtp))
* `rule_id` (Number)
* `sip` (Attributes) - Map Flow Whitelist Rule Sip match definition (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--sip))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--flow5_g"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules.pass_rules.flow5_g`

Read-Only:

* `dnn` (String) - Domain Network Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `type` (String) - Set 5G WL-DB lookup type
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--gtp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules.pass_rules.gtp`

Read-Only:

* `apn` (String) - Access Point Name pattern.  Alphanumeric, '.', '-', and '\*' allowed.
* `interface` (String) - interface type. Mutually exclusive with version. required till H 5.6
* `type` (String) - Set GTP WL-DB lookup type
* `version` (String) - mutually exclusive with interface
* `whitelist_databases` (List of String) - Attach whitelist databases to the map
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--flow_whitelist_rules--pass_rules--sip"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.flow_whitelist_rules.pass_rules.sip`

Read-Only:

* `type` (String) - all:Whitelist based on caller/callee/source/destination IP address, bothAddr: Whitelist source/destination IP address, bothId:Whitelist Caller/Callee Id's, calleeId: Whitelist Callee ID, callerId: Whitelist Caller ID, destIp: Whitelist based on destination IP address, srcIp: Whitelist based on Source IP address
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--fstype"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.fstype`

Read-Only:

* `offset` (Number) - Offset for Mobility Rotational Flowsampling
* `timer` (Number) - Timer for Mobility Rotational Flowsampling in minutes
* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--gs_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.gs_rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--gs_rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--gs_rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--gs_rules--drop_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.gs_rules.drop_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--gs_rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.gs_rules.pass_rules`

Read-Only:

* `comment` (String)
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MAY be used more than once, However, their matching positions MUST be unique
* `rule_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--health_state_reasons"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--ip_rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--roles"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.roles`

Read-Only:

* `editors` (List of String)
* `listeners` (List of String)
* `owners` (List of String)
* `viewers` (List of String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rules`

Read-Only:

* `drop_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules))
* `pass_rules` (Attributes Set) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rules.drop_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--vlan_tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--ip_rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rules.drop_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rules.drop_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--drop_rules--vlan_tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rules.drop_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rules.pass_rules`

Read-Only:

* `bidi` (Boolean)
* `comment` (String)
* `ip_rewrite` (Attributes) - IpRewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--ip_rewrite))
* `matches` (Set of Dynamic) - Set of rule's matching elements. Within a rule, matching elements of the the same type MUST NOT be repeated. The 'position' property of each matching element is not relevant for this rule type as only the outer headers are matched
* `rewrite` (Attributes) - Rewrite options on the packets (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--rewrite))
* `rule_id` (Number)
* `vlan_tag` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--vlan_tag))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--ip_rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rules.pass_rules.ip_rewrite`

Read-Only:

* `dst_ip` (String)
* `src_ip` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--rewrite"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rules.pass_rules.rewrite`

Read-Only:

* `dst_mac` (String)
* `src_mac` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--rules--pass_rules--vlan_tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.rules.pass_rules.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--ingress_traffic_configs--ingress_traffic_map--vlan_tag"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.ingress_traffic_configs.ingress_traffic_map.vlan_tag`

Read-Only:

* `tag_protocol_id` (String)
* `vlan_action` (String)
* `vlan_id` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config`

Read-Only:

* `action` (String)
* `cluster_id` (String) - cluster id
* `config_status` (String)
* `error_message` (String)
* `exporter_alias` (String) - exporter alias associated with monitor
* `exporter_config` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config))
* `gs_group` (String) - gsgroup alias associated with monitor
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--health_state_reasons))
* `monitor_ip_interface` (String) - ip interface alias associated with monitor
* `ref_sols` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--ref_sols))
* `solution_alias` (String) - solution alias
* `user_defined_application_profile` (Dynamic)
* `user_defined_applications` (List of List of String)
* `vport_alias` (String) - vport alias associated with monitor
* `vport_config` (Attributes) - GigaSMART vPort (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--vport_config))
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.exporter_config`

Read-Only:

* `alias` (String)
* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--cef))
* `description` (String)
* `destination` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--mobility_sam))
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--monitor))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--snmp))
* `source` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--source))
* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--cef"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.exporter_config.cef`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--destination"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.exporter_config.destination`

Read-Only:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--mobility_sam"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.exporter_config.mobility_sam`

Read-Only:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--mobility_sam--event_enable))
* `trigger` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--mobility_sam--event_enable"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.exporter_config.mobility_sam.event_enable`

Read-Only:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--monitor"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.exporter_config.monitor`

Read-Only:

* `timeout` (Number) - how often to export in seconds
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--netflow"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.exporter_config.netflow`

Read-Only:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--snmp"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.exporter_config.snmp`

Read-Only:

* `enabled` (Boolean) - snmp reverse lookup enable/disable
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--exporter_config--source"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.exporter_config.source`

Read-Only:

* `ip_interface` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--health_state_reasons"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--ref_sols"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.ref_sols`

Read-Only:

* `alias` (String)
* `associated_map` (String)
* `type` (String)
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--vport_config"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.vport_config`

Read-Only:

* `alias` (String)
* `deferred_binding` (Boolean) - enable/disable deferred-binding
* `fail_over_action` (String)
* `gs_group` (String) - Alias of referenced managing GsGroup
* `health_state` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--vport_config--health_state_reasons))
* `inline_status` (String)
* `inner_traffic_path` (String) - Similar to inline-network traffic-path, applicable for inner map
* `metadata_monitoring` (Attributes) (see [below for nested schema](#nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--vport_config--metadata_monitoring))
* `mode` (String)
* `outer_traffic_path` (String) - Similar to inline-tool/inline-tool-group flex-traffic path, applicable for outer map
* `sa_apf_profile` (String) - ASF session profile
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--vport_config--health_state_reasons"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.vport_config.health_state_reasons`

Read-Only:

* `message` (String) - Read-only. Describes the reason for component's health state
* `severity` (String) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `traffic_health_state_computation_type` (String) - Traffic Health State Computation Type
<a id="nestedatt--engine_level_app_vz_configs--app_vz_config--monitor_solution_config--vport_config--metadata_monitoring"></a>
### Nested Schema for `engine_level_app_vz_configs.app_vz_config.monitor_solution_config.vport_config.metadata_monitoring`

Read-Only:

* `action` (String) - metadata monitoring action
* `exporters` (List of String)
<a id="nestedatt--ip_interface_config"></a>
### Nested Schema for `ip_interface_config`

Read-Only:

* `alias` (String) - ip interface name
* `attach` (List of String) - network ports ,tool ports or circuit ports
* `comment` (String)
* `gateway` (String) - gateway ipv4 or ipv6 address
* `gs_groups` (List of String) - Gs Groups associated with the IP Interface
* `hw_address` (String)
* `ip_address` (String) - ipv4/ipv6 address
* `ip_mask` (String) - ipAddress netmask required with ipAddress
* `ip_type` (String)
* `mtu` (Number)
* `netflow_exporters` (List of String) - Netflow Exporters associated with the IP Interface
* `tags` (Attributes List) (see [below for nested schema](#nestedatt--ip_interface_config--tags))
<a id="nestedatt--ip_interface_config--tags"></a>
### Nested Schema for `ip_interface_config.tags`

Read-Only:

* `tag_key` (String) - Name of the tag
* `tag_values` (List of String) - All possible values of the tag
<a id="nestedatt--location"></a>
### Nested Schema for `location`

Read-Only:

* `cluster_id` (String)
* `engine_ports` (List of String)

