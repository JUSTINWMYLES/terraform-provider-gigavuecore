---
page_title: "gigavuecore_load_all_metadata_application_profiles Data Source - gigavuecore"
subcategory: ""
description: |-
  Load All Metadata Application Profiles
---

# gigavuecore_load_all_metadata_application_profiles Data Source

Load All Metadata Application Profiles

## Example Usage

```terraform
data "gigavuecore_load_all_metadata_application_profiles" "example" {
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

* `items` (Attributes List, computed) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

* `alias` (String) - application profile alias
* `application_id` (Boolean) - only valid with 'export' type
* `applications` (Attributes List) - application and attributes. (see [below for nested schema](#nestedatt--items--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--items--counter))
* `datalink` (Attributes) (see [below for nested schema](#nestedatt--items--datalink))
* `description` (String)
* `flow` (Attributes) (see [below for nested schema](#nestedatt--items--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--items--gtpu))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--items--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--items--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--items--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--items--ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--items--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--items--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--items--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--items--transport))
* `type` (String)
<a id="nestedatt--items--applications"></a>
### Nested Schema for `items.applications`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--items--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name
<a id="nestedatt--items--applications--attributes"></a>
### Nested Schema for `items.applications.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value
<a id="nestedatt--items--counter"></a>
### Nested Schema for `items.counter`

Read-Only:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `inner_byte` (Boolean)
* `inner_byte_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)
<a id="nestedatt--items--datalink"></a>
### Nested Schema for `items.datalink`

Read-Only:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--items--flow"></a>
### Nested Schema for `items.flow`

Read-Only:

* `end_reason` (Boolean)
<a id="nestedatt--items--gtpu"></a>
### Nested Schema for `items.gtpu`

Read-Only:

* `qfi` (Boolean)
* `teid` (Boolean)
<a id="nestedatt--items--interface"></a>
### Nested Schema for `items.interface`

Read-Only:

* `in_name_width` (Number)
* `in_physical_width` (Number)
* `out_physical_width` (Number)
<a id="nestedatt--items--ip"></a>
### Nested Schema for `items.ip`

Read-Only:

* `version` (Boolean)
<a id="nestedatt--items--ipv4"></a>
### Nested Schema for `items.ipv4`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--items--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--items--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--items--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--items--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--items--ipv4--destination"></a>
### Nested Schema for `items.ipv4.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--items--ipv4--fragmentation"></a>
### Nested Schema for `items.ipv4.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--items--ipv4--section"></a>
### Nested Schema for `items.ipv4.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--items--ipv4--source"></a>
### Nested Schema for `items.ipv4.source`

Read-Only:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--items--ipv6"></a>
### Nested Schema for `items.ipv6`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--items--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--items--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--items--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--items--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--items--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--items--ipv6--destination"></a>
### Nested Schema for `items.ipv6.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--items--ipv6--fragmentation"></a>
### Nested Schema for `items.ipv6.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--items--ipv6--length"></a>
### Nested Schema for `items.ipv6.length`

Read-Only:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--items--ipv6--section"></a>
### Nested Schema for `items.ipv6.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--items--ipv6--source"></a>
### Nested Schema for `items.ipv6.source`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--items--outer_ipv4"></a>
### Nested Schema for `items.outer_ipv4`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--items--outer_ipv6"></a>
### Nested Schema for `items.outer_ipv6`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--items--timestamp"></a>
### Nested Schema for `items.timestamp`

Read-Only:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
* `sys_up_time_first` (Boolean)
* `sys_up_time_last` (Boolean)
<a id="nestedatt--items--transport"></a>
### Nested Schema for `items.transport`

Read-Only:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--items--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--items--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--items--transport--udp))
<a id="nestedatt--items--transport--icmp"></a>
### Nested Schema for `items.transport.icmp`

Read-Only:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--items--transport--tcp"></a>
### Nested Schema for `items.transport.tcp`

Read-Only:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--items--transport--udp"></a>
### Nested Schema for `items.transport.udp`

Read-Only:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)

