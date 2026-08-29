---
page_title: "gigavuecore_app_visibility Resource - gigavuecore"
subcategory: ""
description: |-
  Load an Application Intelligence Solution by alias
---

# gigavuecore_app_visibility Resource

Load an Application Intelligence Solution by alias

## Example Usage

```terraform
resource "gigavuecore_app_visibility" "example" {
  app_export_config = {
    cache_config = {
      advance_hash     = true
      alias            = "example"
      description      = "example"
      dpi_inject_limit = 0
      event            = "example"
      exporters        = [ "example" ]
      flow_behavior    = "example"
      match = {
        datalink = {
          mac_dst = true
          mac_src = true
          vlan    = true
        }
        interface = {
          in_name_width     = 0
          in_physical_width = 0
        }
        ip = {
          version = true
        }
        ipv4 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp = true
          fragmentation = {
            flags  = true
            offset = true
          }
          header_len = true
          option_map = true
          precedence = true
          protocol   = true
          section = {
            header_size  = 0
            payload_size = 0
          }
          source = {
            prefix_min_mask = "example"
          }
          tos          = true
          total_length = true
          ttl          = true
        }
        ipv6 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp          = true
          extension_map = true
          flow_label    = true
          fragmentation = {
            flags  = true
            offset = true
          }
          hop_limit = true
          length = {
            header  = true
            payload = true
            total   = true
          }
          next_header = true
          precedence  = true
          section = {
            header_size  = 0
            payload_size = 0
          }
          source = {
            prefix_min_mask = "example"
          }
          traffic_class = true
        }
        transport = {
          dst_port = true
          icmp = {
            ipv4_code = true
            ipv4_type = true
            ipv6_code = true
            ipv6_type = true
          }
          src_port = true
          tcp = {
            ack_number  = true
            dst_port    = true
            flags       = true
            header_len  = true
            seq_number  = true
            src_port    = true
            urgent_ptr  = true
            window_size = true
          }
          udp = {
            dst_port = true
            msg_len  = true
            src_port = true
          }
        }
      }
      multi_collect         = true
      network_profiles      = [ "example" ]
      observation_domain_id = 0
      sampling = {
        mode                 = "example"
        single_sampling_rate = 0
      }
      size = {
        flows = 0
      }
      timeout = {
        idle = 0
      }
    }
    destination_configs = [{
      application_names = [{
        attributes = [{
          name  = "example"
          value = "example"
        }]
        is_user_defined = true
        name            = "example"
      }]
      destination_name = "example"
      export_meta_app_profile = {
        alias          = "example"
        application_id = true
        applications = [{
          attributes = [{
            name  = "example"
            value = "example"
          }]
          is_user_defined = true
          name            = "example"
        }]
        counter = {
          bytes           = true
          bytes_long      = true
          inner_byte      = true
          inner_byte_long = true
          packets         = true
          packets_long    = true
        }
        datalink = {
          mac_dst = true
          mac_src = true
          vlan    = true
        }
        description = "example"
        flow = {
          end_reason = true
        }
        gtpu = {
          qfi  = true
          teid = true
        }
        interface = {
          in_name_width      = 0
          in_physical_width  = 0
          out_physical_width = 0
        }
        ip = {
          version = true
        }
        ipv4 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp = true
          fragmentation = {
            flags  = true
            offset = true
          }
          header_len = true
          option_map = true
          precedence = true
          protocol   = true
          section = {
            header_size  = 0
            payload_size = 0
          }
          source = {
            prefix_min_mask = "example"
          }
          tos          = true
          total_length = true
          ttl          = true
        }
        ipv6 = {
          destination = {
            prefix_min_mask = "example"
          }
          dscp          = true
          extension_map = true
          flow_label    = true
          fragmentation = {
            flags  = true
            offset = true
          }
          hop_limit = true
          length = {
            header  = true
            payload = true
            total   = true
          }
          next_header = true
          precedence  = true
          section = {
            header_size  = 0
            payload_size = 0
          }
          source = {
            prefix_min_mask = "example"
          }
          traffic_class = true
        }
        outer_ipv4 = {
          destination = true
          source      = true
        }
        outer_ipv6 = {
          destination = true
          source      = true
        }
        timestamp = {
          flow_end_msec     = true
          flow_endsec       = true
          flow_start_msec   = true
          flow_startsec     = true
          sys_up_time_first = true
          sys_up_time_last  = true
        }
        transport = {
          dst_port = true
          icmp = {
            ipv4_code = true
            ipv4_type = true
            ipv6_code = true
            ipv6_type = true
          }
          src_port = true
          tcp = {
            ack_number  = true
            dst_port    = true
            flags       = true
            header_len  = true
            seq_number  = true
            src_port    = true
            urgent_ptr  = true
            window_size = true
          }
          udp = {
            dst_port = true
            msg_len  = true
            src_port = true
          }
        }
        type = "example"
      }
      export_meta_app_profile_alias = "example"
      exporter_alias                = "example"
      exporter_config = {
        alias                = "example"
        application_profiles = [ "example" ]
        cef = {
          active_timeout   = 0
          inactive_timeout = 0
        }
        description = "example"
        destination = {
          dscp         = 0
          ipv4_address = "example"
          l4_port_dst  = 0
          l4_port_src  = 0
          l4_protocol  = "example"
          ttl          = 0
        }
        max_pkt_size = 0
        mobility_sam = {
          encoding        = "example"
          encoding_format = "example"
          event_enable = {
            modify = true
            update = true
          }
          trigger = "example"
        }
        monitor = {
          timeout = 0
        }
        netflow = {
          active_timeout   = 0
          inactive_timeout = 0
          template_refresh = 0
          template_type    = "example"
          version          = "example"
        }
        snmp = {
          enabled = true
        }
        source = {
          ip_interface = "example"
        }
        type = "example"
      }
      iface = "example"
    }]
  }
  app_filter_config = {
    egress_traffic_configs = [{
      drop_application_names = [{
        attributes = [{
          name  = "example"
          value = "example"
        }]
        is_user_defined = true
        name            = "example"
      }]
      pass_application_names = [{
        attributes = [{
          name  = "example"
          value = "example"
        }]
        is_user_defined = true
        name            = "example"
      }]
      tunnel_aliases = [ "example" ]
    }]
  }
  conn_id            = "example"
  distribute_traffic = true
  dynamic_scale_unit = true
  env_id             = "example"
  ingress_traffic_configs = [{
    source_selector_aliases      = [ "example" ]
    src_raw_end_point_interfaces = [ "example" ]
    tunnel_aliases               = [ "example" ]
    tunnel_interface_mappings = [{
      iface        = "example"
      tunnel_alias = "example"
    }]
  }]
  monitor_solution_config = {
    destination_config = {
      iface = "example"
    }
    mgmt_interface = "example"
  }
  scale_unit     = 0
  solution_alias = "example"
  solution_desc  = "example"
}
```

## Schema

### Arguments

The following arguments are supported:

* `app_export_config` (Attributes, optional) (see [below for nested schema](#nestedatt--app_export_config))
* `app_filter_config` (Attributes, optional) (see [below for nested schema](#nestedatt--app_filter_config))
* `conn_id` (String, required) - ID of the Connection Domain
* `distribute_traffic` (Boolean, optional) - Indicates Traffic Distribution enabled or not
* `dynamic_scale_unit` (Boolean, optional) - When set as true, FM automatically sets the optimal scaleUnit
* `env_id` (String, required) - ID of the Environment
* `ingress_traffic_configs` (Attributes List, optional) (see [below for nested schema](#nestedatt--ingress_traffic_configs))
* `monitor_solution_config` (Attributes, required) (see [below for nested schema](#nestedatt--monitor_solution_config))
* `scale_unit` (Number, optional) - Number of units of memory needed for each Application
* `solution_alias` (String, required) - Alias of the solution
* `solution_desc` (String, optional) - Description of the solution

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `app_env_id` (String, computed) - ID of the Application Environment
* `app_exporter_config` (Attributes, computed) (see [below for nested schema](#nestedatt--app_exporter_config))
* `app_filter_config` (Attributes, computed) (see [below for nested schema](#nestedatt--app_filter_config))
* `config_status` (String, computed) - configuration status
* `config_status_reasons` (String, computed) - configuration status reasons
* `dedup` (Attributes, computed) (see [below for nested schema](#nestedatt--dedup))
* `dynamic_scale_unit` (Boolean, computed) - When set as true, FM automatically sets the optimal scaleUnit
* `health_state` (String, computed) - Read-only. 'green' indicates healthy state; 'yellow'  indicates warning state; 'orange'  indicates error state; 'red'  indicates critical state;
* `health_state_reasons` (Attributes List, computed) (see [below for nested schema](#nestedatt--health_state_reasons))
* `ingress_traffic_configs` (Attributes List, computed) (see [below for nested schema](#nestedatt--ingress_traffic_configs))
* `scale_unit` (Number, computed) - Number of units of memory needed for each Application
* `solution_created_timestamp` (String, computed)
* `solution_desc` (String, computed) - Description of the solution
* `solution_modified_timestamp` (String, computed)
* `traffic_policy_graph_alias` (String, computed) - Alias of the Traffic Policy Graph

<a id="nestedatt--app_export_config"></a>
### Nested Schema for `app_export_config`

Optional:

* `cache_config` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config))
* `destination_configs` (Attributes List) (see [below for nested schema](#nestedatt--app_export_config--destination_configs))
<a id="nestedatt--app_export_config--cache_config"></a>
### Nested Schema for `app_export_config.cache_config`

Required:

* `alias` (String)
Optional:

* `advance_hash` (Boolean) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
* `description` (String)
* `dpi_inject_limit` (Number)
* `event` (String)
* `exporters` (List of String) - alias of metadata exporters to attach this cache
* `flow_behavior` (String) - direction for flow identification
* `match` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match))
* `multi_collect` (Boolean) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
* `network_profiles` (List of String) - alias of metadata network profiles to attach this cache
* `observation_domain_id` (Number)
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--sampling))
* `size` (Attributes) - size of the flows (see [below for nested schema](#nestedatt--app_export_config--cache_config--size))
* `timeout` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--timeout))
<a id="nestedatt--app_export_config--cache_config--match"></a>
### Nested Schema for `app_export_config.cache_config.match`

Optional:

* `datalink` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--datalink))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--transport))
<a id="nestedatt--app_export_config--cache_config--match--datalink"></a>
### Nested Schema for `app_export_config.cache_config.match.datalink`

Optional:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--interface"></a>
### Nested Schema for `app_export_config.cache_config.match.interface`

Optional:

* `in_name_width` (Number)
* `in_physical_width` (Number)
<a id="nestedatt--app_export_config--cache_config--match--ip"></a>
### Nested Schema for `app_export_config.cache_config.match.ip`

Optional:

* `version` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv4"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv4--destination"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4.destination`

Optional:

* `prefix_min_mask` (String) - ipv4 destination prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_export_config--cache_config--match--ipv4--fragmentation"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv4--section"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_export_config--cache_config--match--ipv4--source"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv4.source`

Optional:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_export_config--cache_config--match--ipv6"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--destination"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--fragmentation"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--length"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.length`

Optional:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--section"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_export_config--cache_config--match--ipv6--source"></a>
### Nested Schema for `app_export_config.cache_config.match.ipv6.source`

Optional:

* `prefix_min_mask` (String) - ipv6 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_export_config--cache_config--match--transport"></a>
### Nested Schema for `app_export_config.cache_config.match.transport`

Optional:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--cache_config--match--transport--udp))
<a id="nestedatt--app_export_config--cache_config--match--transport--icmp"></a>
### Nested Schema for `app_export_config.cache_config.match.transport.icmp`

Optional:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--transport--tcp"></a>
### Nested Schema for `app_export_config.cache_config.match.transport.tcp`

Optional:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--app_export_config--cache_config--match--transport--udp"></a>
### Nested Schema for `app_export_config.cache_config.match.transport.udp`

Optional:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--app_export_config--cache_config--sampling"></a>
### Nested Schema for `app_export_config.cache_config.sampling`

Optional:

* `mode` (String)
* `single_sampling_rate` (Number) - Packet interval window size. Valid values: 10-16000 (in packets)
<a id="nestedatt--app_export_config--cache_config--size"></a>
### Nested Schema for `app_export_config.cache_config.size`

Optional:

* `flows` (Number) - size of flows in millions
<a id="nestedatt--app_export_config--cache_config--timeout"></a>
### Nested Schema for `app_export_config.cache_config.timeout`

Optional:

* `idle` (Number) - idle timeout in seconds. max value 7days. default 30 min
<a id="nestedatt--app_export_config--destination_configs"></a>
### Nested Schema for `app_export_config.destination_configs`

Optional:

* `application_names` (Attributes List) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--application_names))
* `destination_name` (String)
* `export_meta_app_profile` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile))
* `export_meta_app_profile_alias` (String)
* `exporter_alias` (String)
* `exporter_config` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config))
* `iface` (String) - Interface Mapping for Egress Tunnel
<a id="nestedatt--app_export_config--destination_configs--application_names"></a>
### Nested Schema for `app_export_config.destination_configs.application_names`

Required:

* `name` (String) - application name
Optional:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
<a id="nestedatt--app_export_config--destination_configs--application_names--attributes"></a>
### Nested Schema for `app_export_config.destination_configs.application_names.attributes`

Required:

* `name` (String) - attribute name
Optional:

* `value` (String) - application's attribute value
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile`

Required:

* `alias` (String) - application profile alias
Optional:

* `application_id` (Boolean) - only valid with 'export' type
* `applications` (Attributes List) - application and attributes. (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--counter))
* `datalink` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--datalink))
* `description` (String)
* `flow` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--gtpu))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport))
* `type` (String)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--applications"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.applications`

Required:

* `name` (String) - application name
Optional:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--applications--attributes"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.applications.attributes`

Required:

* `name` (String) - attribute name
Optional:

* `value` (String) - application's attribute value
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--counter"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.counter`

Optional:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `inner_byte` (Boolean)
* `inner_byte_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--datalink"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.datalink`

Optional:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--flow"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.flow`

Optional:

* `end_reason` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--gtpu"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.gtpu`

Optional:

* `qfi` (Boolean)
* `teid` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--interface"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.interface`

Optional:

* `in_name_width` (Number)
* `in_physical_width` (Number)
* `out_physical_width` (Number)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ip"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ip`

Optional:

* `version` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--destination"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--fragmentation"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--section"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv4--source"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv4.source`

Optional:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--destination"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--fragmentation"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--length"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.length`

Optional:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--section"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--ipv6--source"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.ipv6.source`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--outer_ipv4"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.outer_ipv4`

Optional:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--outer_ipv6"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.outer_ipv6`

Optional:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--timestamp"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.timestamp`

Optional:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
* `sys_up_time_first` (Boolean)
* `sys_up_time_last` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.transport`

Optional:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--udp))
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--icmp"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.transport.icmp`

Optional:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--tcp"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.transport.tcp`

Optional:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--export_meta_app_profile--transport--udp"></a>
### Nested Schema for `app_export_config.destination_configs.export_meta_app_profile.transport.udp`

Optional:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--exporter_config"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config`

Required:

* `alias` (String)
Optional:

* `application_profiles` (List of String) - application profile aliases to attach to the exporter
* `cef` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--cef))
* `description` (String)
* `destination` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--destination))
* `max_pkt_size` (Number)
* `mobility_sam` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--mobility_sam))
* `monitor` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--monitor))
* `netflow` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--netflow))
* `snmp` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--snmp))
* `source` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--source))
* `type` (String)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--cef"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.cef`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
<a id="nestedatt--app_export_config--destination_configs--exporter_config--destination"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.destination`

Optional:

* `dscp` (Number)
* `ipv4_address` (String) - ipv4 address
* `l4_port_dst` (Number)
* `l4_port_src` (Number)
* `l4_protocol` (String)
* `ttl` (Number)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--mobility_sam"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.mobility_sam`

Optional:

* `encoding` (String)
* `encoding_format` (String)
* `event_enable` (Attributes) (see [below for nested schema](#nestedatt--app_export_config--destination_configs--exporter_config--mobility_sam--event_enable))
* `trigger` (String)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--mobility_sam--event_enable"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.mobility_sam.event_enable`

Optional:

* `modify` (Boolean)
* `update` (Boolean)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--monitor"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.monitor`

Optional:

* `timeout` (Number) - how often to export in seconds
<a id="nestedatt--app_export_config--destination_configs--exporter_config--netflow"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.netflow`

Optional:

* `active_timeout` (Number) - in seconds
* `inactive_timeout` (Number) - in seconds
* `template_refresh` (Number) - template refresh interval in seconds
* `template_type` (String)
* `version` (String)
<a id="nestedatt--app_export_config--destination_configs--exporter_config--snmp"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.snmp`

Optional:

* `enabled` (Boolean) - snmp reverse lookup enable/disable
<a id="nestedatt--app_export_config--destination_configs--exporter_config--source"></a>
### Nested Schema for `app_export_config.destination_configs.exporter_config.source`

Optional:

* `ip_interface` (String)
<a id="nestedatt--app_filter_config"></a>
### Nested Schema for `app_filter_config`

Optional:

* `egress_traffic_configs` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs))
Read-Only:

* `app_filter_tiered_batch_id` (String) - ID of the appFilterConfig tiered batch
* `map_alias` (String) - Alias of the appFiltering map
* `sapf_profile_alias` (String) - Alias of the SAPF Profile
<a id="nestedatt--app_filter_config--egress_traffic_configs"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs`

Optional:

* `drop_application_names` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--drop_application_names))
* `pass_application_names` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--pass_application_names))
* `tunnel_aliases` (List of String)
Read-Only:

* `priority` (Number)
<a id="nestedatt--app_filter_config--egress_traffic_configs--drop_application_names"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.drop_application_names`

Required:

* `name` (String) - application name
Optional:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--drop_application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
<a id="nestedatt--app_filter_config--egress_traffic_configs--drop_application_names--attributes"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.drop_application_names.attributes`

Required:

* `name` (String) - attribute name
Optional:

* `value` (String) - application's attribute value
<a id="nestedatt--app_filter_config--egress_traffic_configs--pass_application_names"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.pass_application_names`

Required:

* `name` (String) - application name
Optional:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--app_filter_config--egress_traffic_configs--pass_application_names--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
<a id="nestedatt--app_filter_config--egress_traffic_configs--pass_application_names--attributes"></a>
### Nested Schema for `app_filter_config.egress_traffic_configs.pass_application_names.attributes`

Required:

* `name` (String) - attribute name
Optional:

* `value` (String) - application's attribute value
<a id="nestedatt--ingress_traffic_configs"></a>
### Nested Schema for `ingress_traffic_configs`

Optional:

* `source_selector_aliases` (List of String)
* `src_raw_end_point_interfaces` (List of String)
* `tunnel_aliases` (List of String)
* `tunnel_interface_mappings` (Attributes List) (see [below for nested schema](#nestedatt--ingress_traffic_configs--tunnel_interface_mappings))
<a id="nestedatt--ingress_traffic_configs--tunnel_interface_mappings"></a>
### Nested Schema for `ingress_traffic_configs.tunnel_interface_mappings`

Optional:

* `iface` (String)
* `tunnel_alias` (String)
<a id="nestedatt--monitor_solution_config"></a>
### Nested Schema for `monitor_solution_config`

Required:

* `destination_config` (Attributes) (see [below for nested schema](#nestedatt--monitor_solution_config--destination_config))
Optional:

* `mgmt_interface` (String) - Pass this value if you would like to use the management interface of the Node to export App Monitoring metadata to FM
Read-Only:

* `app_instance_config_id` (String)
* `app_instance_name` (String)
* `app_viz_tiered_batch_id` (String) - ID of the monitoringSolutionConfig tiered batch
* `gs_param_config_id` (String) - Gs param name on which the solution is applied
* `user_defined_application_profile` (Dynamic)
* `user_defined_applications` (List of List of String)
<a id="nestedatt--monitor_solution_config--destination_config"></a>
### Nested Schema for `monitor_solution_config.destination_config`

Optional:

* `iface` (String) - Interface Mapping for Egress Tunnel
Read-Only:

* `destination_name` (String)
* `exporter_alias` (String)
<a id="nestedatt--app_exporter_config"></a>
### Nested Schema for `app_exporter_config`

Read-Only:

* `app_instance_name` (String) - The name of the AppMetadata app instance
* `app_metadata_tiered_batch_id` (String) - ID of the appExporterConfig tiered batch
* `cache_config_alias` (String) - Alias of the cacheConfig
* `destination_configs` (Attributes List) (see [below for nested schema](#nestedatt--app_exporter_config--destination_configs))
<a id="nestedatt--app_exporter_config--destination_configs"></a>
### Nested Schema for `app_exporter_config.destination_configs`

Read-Only:

* `destination_name` (String) - Alias of the destination tool
* `exporter_alias` (String) - Alias of the exporter
* `iface` (String) - Interface Mapping for Egress Tunnel
* `template_name` (String) - Name of the App Profile Template Used
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

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_app_visibility.example {solution_alias}
```
