---
page_title: "gigavuecore_cache Resource - gigavuecore"
subcategory: ""
description: |-
  Load Metadata Cache by Alias
---

# gigavuecore_cache Resource

Load Metadata Cache by Alias

## Example Usage

```terraform
resource "gigavuecore_cache" "example" {
  advance_hash = null
  alias = null
  description = null
  dpi_inject_limit = null
  event = null
  exporters = []
  flow_behavior = null
  match = {}
  multi_collect = null
  network_profiles = []
  observation_domain_id = null
  sampling = {}
  size = {}
  timeout = {}
}
```

## Schema

### Arguments

The following arguments are supported:

* `advance_hash` (Bool, optional) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
* `alias` (String, required)
* `description` (String, optional)
* `dpi_inject_limit` (Number, optional)
* `event` (String, optional)
* `exporters` (List(String), optional) - alias of metadata exporters to attach this cache
* `flow_behavior` (String, optional) - direction for flow identification
* `match` (Object({datalink, interface, ip, ipv4, ipv6, transport}), optional)
  * `datalink` (Object({mac_dst, mac_src, vlan}), optional)
    * `mac_dst` (Bool, optional)
    * `mac_src` (Bool, optional)
    * `vlan` (Bool, optional)
  * `interface` (Object({in_name_width, in_physical_width}), optional)
    * `in_name_width` (Number, optional)
    * `in_physical_width` (Number, optional)
  * `ip` (Object({version}), optional)
    * `version` (Bool, optional)
  * `ipv4` (Object({destination, dscp, fragmentation, header_len, option_map, precedence, protocol, section, source, tos, total_length, ttl}), optional)
    * `destination` (Object({prefix_min_mask}), optional)
      * `prefix_min_mask` (String, optional) - ipv4 destination prefix minimum-mask - netmask or mask length
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
      * `prefix_min_mask` (String, optional) - ipv6 source prefix minimum-mask - netmask or mask length
    * `traffic_class` (Bool, optional)
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
* `multi_collect` (Bool, optional) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
* `network_profiles` (List(String), optional) - alias of metadata network profiles to attach this cache
* `observation_domain_id` (Number, optional)
* `sampling` (Object({mode, single_sampling_rate}), optional)
  * `mode` (String, optional)
  * `single_sampling_rate` (Number, optional) - Packet interval window size. Valid values: 10-16000 (in packets)
* `size` (Object({flows}), optional) - size of the flows
  * `flows` (Number, optional) - size of flows in millions
* `timeout` (Object({idle}), optional)
  * `idle` (Number, optional) - idle timeout in seconds. max value 7days. default 30 min

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `advance_hash` (Bool, computed) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
* `description` (String, computed)
* `dpi_inject_limit` (Number, computed)
* `event` (String, computed)
* `exporters` (List(String), computed) - alias of metadata exporters to attach this cache
* `flow_behavior` (String, computed) - direction for flow identification
* `match` (Object({datalink, interface, ip, ipv4, ipv6, transport}), computed)
  * `datalink` (Object({mac_dst, mac_src, vlan}), optional)
    * `mac_dst` (Bool, optional)
    * `mac_src` (Bool, optional)
    * `vlan` (Bool, optional)
  * `interface` (Object({in_name_width, in_physical_width}), optional)
    * `in_name_width` (Number, optional)
    * `in_physical_width` (Number, optional)
  * `ip` (Object({version}), optional)
    * `version` (Bool, optional)
  * `ipv4` (Object({destination, dscp, fragmentation, header_len, option_map, precedence, protocol, section, source, tos, total_length, ttl}), optional)
    * `destination` (Object({prefix_min_mask}), optional)
      * `prefix_min_mask` (String, optional) - ipv4 destination prefix minimum-mask - netmask or mask length
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
      * `prefix_min_mask` (String, optional) - ipv6 source prefix minimum-mask - netmask or mask length
    * `traffic_class` (Bool, optional)
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
* `multi_collect` (Bool, computed) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
* `network_profiles` (List(String), computed) - alias of metadata network profiles to attach this cache
* `observation_domain_id` (Number, computed)
* `sampling` (Object({mode, single_sampling_rate}), computed)
  * `mode` (String, optional)
  * `single_sampling_rate` (Number, optional) - Packet interval window size. Valid values: 10-16000 (in packets)
* `size` (Object({flows}), computed) - size of the flows
  * `flows` (Number, optional) - size of flows in millions
* `timeout` (Object({idle}), computed)
  * `idle` (Number, optional) - idle timeout in seconds. max value 7days. default 30 min

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_cache.example {alias}
```
