---
page_title: "gigavuecore_load_all_metadata_cache Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Metadata Cache
---

# gigavuecore_load_all_metadata_cache Data Source

Load All Metadata Cache

## Example Usage

```terraform
data "gigavuecore_load_all_metadata_cache" "example" {
  page = null
  sort = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `page` (String, optional) - parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned
* `sort` (String, optional) - parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)

### Attributes

In addition to all arguments above, the following attributes are exported:

* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))
* `metadata_caches` (Attributes List, computed) (see [below for nested schema](#nestedatt--metadata_caches))

<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type
<a id="nestedatt--metadata_caches"></a>
### Nested Schema for `metadata_caches`

Read-Only:

* `advance_hash` (Boolean) - When scheduling flows for processing, use external encapsulate packet header if disable, use inner packet header if enable.
* `alias` (String)
* `description` (String)
* `dpi_inject_limit` (Number)
* `event` (String)
* `exporters` (List of String) - alias of metadata exporters to attach this cache
* `flow_behavior` (String) - direction for flow identification
* `match` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match))
* `multi_collect` (Boolean) - Collect all attributes as it is discovered when enable. It will export the same record once when disable.
* `network_profiles` (List of String) - alias of metadata network profiles to attach this cache
* `observation_domain_id` (Number)
* `sampling` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--sampling))
* `size` (Attributes) - size of the flows (see [below for nested schema](#nestedatt--metadata_caches--size))
* `timeout` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--timeout))
<a id="nestedatt--metadata_caches--match"></a>
### Nested Schema for `metadata_caches.match`

Read-Only:

* `datalink` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--datalink))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv6))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--transport))
<a id="nestedatt--metadata_caches--match--datalink"></a>
### Nested Schema for `metadata_caches.match.datalink`

Read-Only:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--metadata_caches--match--interface"></a>
### Nested Schema for `metadata_caches.match.interface`

Read-Only:

* `in_name_width` (Number)
* `in_physical_width` (Number)
<a id="nestedatt--metadata_caches--match--ip"></a>
### Nested Schema for `metadata_caches.match.ip`

Read-Only:

* `version` (Boolean)
<a id="nestedatt--metadata_caches--match--ipv4"></a>
### Nested Schema for `metadata_caches.match.ipv4`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--metadata_caches--match--ipv4--destination"></a>
### Nested Schema for `metadata_caches.match.ipv4.destination`

Read-Only:

* `prefix_min_mask` (String) - ipv4 destination prefix minimum-mask - netmask or mask length
<a id="nestedatt--metadata_caches--match--ipv4--fragmentation"></a>
### Nested Schema for `metadata_caches.match.ipv4.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--metadata_caches--match--ipv4--section"></a>
### Nested Schema for `metadata_caches.match.ipv4.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--metadata_caches--match--ipv4--source"></a>
### Nested Schema for `metadata_caches.match.ipv4.source`

Read-Only:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--metadata_caches--match--ipv6"></a>
### Nested Schema for `metadata_caches.match.ipv6`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--metadata_caches--match--ipv6--destination"></a>
### Nested Schema for `metadata_caches.match.ipv6.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--metadata_caches--match--ipv6--fragmentation"></a>
### Nested Schema for `metadata_caches.match.ipv6.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--metadata_caches--match--ipv6--length"></a>
### Nested Schema for `metadata_caches.match.ipv6.length`

Read-Only:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--metadata_caches--match--ipv6--section"></a>
### Nested Schema for `metadata_caches.match.ipv6.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--metadata_caches--match--ipv6--source"></a>
### Nested Schema for `metadata_caches.match.ipv6.source`

Read-Only:

* `prefix_min_mask` (String) - ipv6 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--metadata_caches--match--transport"></a>
### Nested Schema for `metadata_caches.match.transport`

Read-Only:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--metadata_caches--match--transport--udp))
<a id="nestedatt--metadata_caches--match--transport--icmp"></a>
### Nested Schema for `metadata_caches.match.transport.icmp`

Read-Only:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--metadata_caches--match--transport--tcp"></a>
### Nested Schema for `metadata_caches.match.transport.tcp`

Read-Only:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--metadata_caches--match--transport--udp"></a>
### Nested Schema for `metadata_caches.match.transport.udp`

Read-Only:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--metadata_caches--sampling"></a>
### Nested Schema for `metadata_caches.sampling`

Read-Only:

* `mode` (String)
* `single_sampling_rate` (Number) - Packet interval window size. Valid values: 10-16000 (in packets)
<a id="nestedatt--metadata_caches--size"></a>
### Nested Schema for `metadata_caches.size`

Read-Only:

* `flows` (Number) - size of flows in millions
<a id="nestedatt--metadata_caches--timeout"></a>
### Nested Schema for `metadata_caches.timeout`

Read-Only:

* `idle` (Number) - idle timeout in seconds. max value 7days. default 30 min

