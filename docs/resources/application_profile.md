---
page_title: "gigavuecore_application_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Load Metadata Application Profile by alias
---

# gigavuecore_application_profile Resource

Load Metadata Application Profile by alias

## Example Usage

```terraform
resource "gigavuecore_application_profile" "example" {
  alias = null
  application_id = null
  applications = []
  counter = {}
  datalink = {}
  description = null
  flow = {}
  gtpu = {}
  interface = {}
  ip = {}
  ipv4 = {}
  ipv6 = {}
  outer_ipv4 = {}
  outer_ipv6 = {}
  timestamp = {}
  transport = {}
  type = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - application profile alias
* `application_id` (Bool, optional) - only valid with 'export' type
* `applications` (List(Object({attributes, is_user_defined, name})), optional) - application and attributes.
* `counter` (Object({bytes, bytes_long, inner_byte, inner_byte_long, packets, packets_long}), optional)
  * `bytes` (Bool, optional)
  * `bytes_long` (Bool, optional)
  * `inner_byte` (Bool, optional)
  * `inner_byte_long` (Bool, optional)
  * `packets` (Bool, optional)
  * `packets_long` (Bool, optional)
* `datalink` (Object({mac_dst, mac_src, vlan}), optional)
  * `mac_dst` (Bool, optional)
  * `mac_src` (Bool, optional)
  * `vlan` (Bool, optional)
* `description` (String, optional)
* `flow` (Object({end_reason}), optional)
  * `end_reason` (Bool, optional)
* `gtpu` (Object({qfi, teid}), optional)
  * `qfi` (Bool, optional)
  * `teid` (Bool, optional)
* `interface` (Object({in_name_width, in_physical_width, out_physical_width}), optional)
  * `in_name_width` (Number, optional)
  * `in_physical_width` (Number, optional)
  * `out_physical_width` (Number, optional)
* `ip` (Object({version}), optional)
  * `version` (Bool, optional)
* `ipv4` (Object({destination, dscp, fragmentation, header_len, option_map, precedence, protocol, section, source, tos, total_length, ttl}), optional)
  * `destination` (Object({prefix_min_mask}), optional)
    * `prefix_min_mask` (String, optional)
  * `dscp` (Bool, optional)
  * `fragmentation` (Object({flags, offset}), optional)
    * `flags` (Bool, optional)
    * `offset` (Bool, optional)
  * `header_len` (Bool, optional)
  * `option_map` (Bool, optional)
  * `precedence` (Bool, optional)
  * `protocol` (Bool, optional)
  * `section` (Object({header_size, payload_size}), optional)
    * `header_size` (Number, optional)
    * `payload_size` (Number, optional)
  * `source` (Object({prefix_min_mask}), optional)
    * `prefix_min_mask` (String, optional) - ipv4 source prefix minimum-mask - netmask or mask length
  * `tos` (Bool, optional)
  * `total_length` (Bool, optional)
  * `ttl` (Bool, optional)
* `ipv6` (Object({destination, dscp, extension_map, flow_label, fragmentation, hop_limit, length, next_header, precedence, section, source, traffic_class}), optional)
  * `destination` (Object({prefix_min_mask}), optional)
    * `prefix_min_mask` (String, optional)
  * `dscp` (Bool, optional)
  * `extension_map` (Bool, optional)
  * `flow_label` (Bool, optional)
  * `fragmentation` (Object({flags, offset}), optional)
    * `flags` (Bool, optional)
    * `offset` (Bool, optional)
  * `hop_limit` (Bool, optional)
  * `length` (Object({header, payload, total}), optional)
    * `header` (Bool, optional)
    * `payload` (Bool, optional)
    * `total` (Bool, optional)
  * `next_header` (Bool, optional)
  * `precedence` (Bool, optional)
  * `section` (Object({header_size, payload_size}), optional)
    * `header_size` (Number, optional)
    * `payload_size` (Number, optional)
  * `source` (Object({prefix_min_mask}), optional)
    * `prefix_min_mask` (String, optional)
  * `traffic_class` (Bool, optional)
* `outer_ipv4` (Object({destination, source}), optional)
  * `destination` (Bool, optional)
  * `source` (Bool, optional)
* `outer_ipv6` (Object({destination, source}), optional)
  * `destination` (Bool, optional)
  * `source` (Bool, optional)
* `timestamp` (Object({flow_end_msec, flow_endsec, flow_start_msec, flow_startsec, sys_up_time_first, sys_up_time_last}), optional)
  * `flow_end_msec` (Bool, optional)
  * `flow_endsec` (Bool, optional)
  * `flow_start_msec` (Bool, optional)
  * `flow_startsec` (Bool, optional)
  * `sys_up_time_first` (Bool, optional)
  * `sys_up_time_last` (Bool, optional)
* `transport` (Object({dst_port, icmp, src_port, tcp, udp}), optional)
  * `dst_port` (Bool, optional)
  * `icmp` (Object({ipv4_code, ipv4_type, ipv6_code, ipv6_type}), optional)
    * `ipv4_code` (Bool, optional)
    * `ipv4_type` (Bool, optional)
    * `ipv6_code` (Bool, optional)
    * `ipv6_type` (Bool, optional)
  * `src_port` (Bool, optional)
  * `tcp` (Object({ack_number, dst_port, flags, header_len, seq_number, src_port, urgent_ptr, window_size}), optional)
    * `ack_number` (Bool, optional)
    * `dst_port` (Bool, optional)
    * `flags` (Bool, optional)
    * `header_len` (Bool, optional)
    * `seq_number` (Bool, optional)
    * `src_port` (Bool, optional)
    * `urgent_ptr` (Bool, optional)
    * `window_size` (Bool, optional)
  * `udp` (Object({dst_port, msg_len, src_port}), optional)
    * `dst_port` (Bool, optional)
    * `msg_len` (Bool, optional)
    * `src_port` (Bool, optional)
* `type` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `application_id` (Bool, computed) - only valid with 'export' type
* `applications` (List(Object({attributes, is_user_defined, name})), computed) - application and attributes.
* `counter` (Object({bytes, bytes_long, inner_byte, inner_byte_long, packets, packets_long}), computed)
  * `bytes` (Bool, optional)
  * `bytes_long` (Bool, optional)
  * `inner_byte` (Bool, optional)
  * `inner_byte_long` (Bool, optional)
  * `packets` (Bool, optional)
  * `packets_long` (Bool, optional)
* `datalink` (Object({mac_dst, mac_src, vlan}), computed)
  * `mac_dst` (Bool, optional)
  * `mac_src` (Bool, optional)
  * `vlan` (Bool, optional)
* `description` (String, computed)
* `flow` (Object({end_reason}), computed)
  * `end_reason` (Bool, optional)
* `gtpu` (Object({qfi, teid}), computed)
  * `qfi` (Bool, optional)
  * `teid` (Bool, optional)
* `interface` (Object({in_name_width, in_physical_width, out_physical_width}), computed)
  * `in_name_width` (Number, optional)
  * `in_physical_width` (Number, optional)
  * `out_physical_width` (Number, optional)
* `ip` (Object({version}), computed)
  * `version` (Bool, optional)
* `ipv4` (Object({destination, dscp, fragmentation, header_len, option_map, precedence, protocol, section, source, tos, total_length, ttl}), computed)
  * `destination` (Object({prefix_min_mask}), optional)
    * `prefix_min_mask` (String, optional)
  * `dscp` (Bool, optional)
  * `fragmentation` (Object({flags, offset}), optional)
    * `flags` (Bool, optional)
    * `offset` (Bool, optional)
  * `header_len` (Bool, optional)
  * `option_map` (Bool, optional)
  * `precedence` (Bool, optional)
  * `protocol` (Bool, optional)
  * `section` (Object({header_size, payload_size}), optional)
    * `header_size` (Number, optional)
    * `payload_size` (Number, optional)
  * `source` (Object({prefix_min_mask}), optional)
    * `prefix_min_mask` (String, optional) - ipv4 source prefix minimum-mask - netmask or mask length
  * `tos` (Bool, optional)
  * `total_length` (Bool, optional)
  * `ttl` (Bool, optional)
* `ipv6` (Object({destination, dscp, extension_map, flow_label, fragmentation, hop_limit, length, next_header, precedence, section, source, traffic_class}), computed)
  * `destination` (Object({prefix_min_mask}), optional)
    * `prefix_min_mask` (String, optional)
  * `dscp` (Bool, optional)
  * `extension_map` (Bool, optional)
  * `flow_label` (Bool, optional)
  * `fragmentation` (Object({flags, offset}), optional)
    * `flags` (Bool, optional)
    * `offset` (Bool, optional)
  * `hop_limit` (Bool, optional)
  * `length` (Object({header, payload, total}), optional)
    * `header` (Bool, optional)
    * `payload` (Bool, optional)
    * `total` (Bool, optional)
  * `next_header` (Bool, optional)
  * `precedence` (Bool, optional)
  * `section` (Object({header_size, payload_size}), optional)
    * `header_size` (Number, optional)
    * `payload_size` (Number, optional)
  * `source` (Object({prefix_min_mask}), optional)
    * `prefix_min_mask` (String, optional)
  * `traffic_class` (Bool, optional)
* `outer_ipv4` (Object({destination, source}), computed)
  * `destination` (Bool, optional)
  * `source` (Bool, optional)
* `outer_ipv6` (Object({destination, source}), computed)
  * `destination` (Bool, optional)
  * `source` (Bool, optional)
* `timestamp` (Object({flow_end_msec, flow_endsec, flow_start_msec, flow_startsec, sys_up_time_first, sys_up_time_last}), computed)
  * `flow_end_msec` (Bool, optional)
  * `flow_endsec` (Bool, optional)
  * `flow_start_msec` (Bool, optional)
  * `flow_startsec` (Bool, optional)
  * `sys_up_time_first` (Bool, optional)
  * `sys_up_time_last` (Bool, optional)
* `transport` (Object({dst_port, icmp, src_port, tcp, udp}), computed)
  * `dst_port` (Bool, optional)
  * `icmp` (Object({ipv4_code, ipv4_type, ipv6_code, ipv6_type}), optional)
    * `ipv4_code` (Bool, optional)
    * `ipv4_type` (Bool, optional)
    * `ipv6_code` (Bool, optional)
    * `ipv6_type` (Bool, optional)
  * `src_port` (Bool, optional)
  * `tcp` (Object({ack_number, dst_port, flags, header_len, seq_number, src_port, urgent_ptr, window_size}), optional)
    * `ack_number` (Bool, optional)
    * `dst_port` (Bool, optional)
    * `flags` (Bool, optional)
    * `header_len` (Bool, optional)
    * `seq_number` (Bool, optional)
    * `src_port` (Bool, optional)
    * `urgent_ptr` (Bool, optional)
    * `window_size` (Bool, optional)
  * `udp` (Object({dst_port, msg_len, src_port}), optional)
    * `dst_port` (Bool, optional)
    * `msg_len` (Bool, optional)
    * `src_port` (Bool, optional)
* `type` (String, computed)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_application_profile.example {alias}
```
