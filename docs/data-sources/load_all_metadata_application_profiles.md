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

* `application_profiles` (Attributes List, computed) (see [below for nested schema](#nestedatt--application_profiles))
* `context` (Attributes, computed) - Gigamon query result context (see [below for nested schema](#nestedatt--context))

<a id="nestedatt--application_profiles"></a>
### Nested Schema for `application_profiles`

Read-Only:

* `alias` (String) - application profile alias
* `application_id` (Boolean) - only valid with 'export' type
* `applications` (Attributes List) - application and attributes. (see [below for nested schema](#nestedatt--application_profiles--applications))
* `counter` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--counter))
* `datalink` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--datalink))
* `description` (String)
* `flow` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--flow))
* `gtpu` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--gtpu))
* `interface` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--interface))
* `ip` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ip))
* `ipv4` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv4))
* `ipv6` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv6))
* `outer_ipv4` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--outer_ipv4))
* `outer_ipv6` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--outer_ipv6))
* `timestamp` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--timestamp))
* `transport` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--transport))
* `type` (String)
<a id="nestedatt--application_profiles--applications"></a>
### Nested Schema for `application_profiles.applications`

Read-Only:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--application_profiles--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
* `name` (String) - application name
<a id="nestedatt--application_profiles--applications--attributes"></a>
### Nested Schema for `application_profiles.applications.attributes`

Read-Only:

* `name` (String) - attribute name
* `value` (String) - application's attribute value
<a id="nestedatt--application_profiles--counter"></a>
### Nested Schema for `application_profiles.counter`

Read-Only:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `inner_byte` (Boolean)
* `inner_byte_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)
<a id="nestedatt--application_profiles--datalink"></a>
### Nested Schema for `application_profiles.datalink`

Read-Only:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--application_profiles--flow"></a>
### Nested Schema for `application_profiles.flow`

Read-Only:

* `end_reason` (Boolean)
<a id="nestedatt--application_profiles--gtpu"></a>
### Nested Schema for `application_profiles.gtpu`

Read-Only:

* `qfi` (Boolean)
* `teid` (Boolean)
<a id="nestedatt--application_profiles--interface"></a>
### Nested Schema for `application_profiles.interface`

Read-Only:

* `in_name_width` (Number)
* `in_physical_width` (Number)
* `out_physical_width` (Number)
<a id="nestedatt--application_profiles--ip"></a>
### Nested Schema for `application_profiles.ip`

Read-Only:

* `version` (Boolean)
<a id="nestedatt--application_profiles--ipv4"></a>
### Nested Schema for `application_profiles.ipv4`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--application_profiles--ipv4--destination"></a>
### Nested Schema for `application_profiles.ipv4.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--application_profiles--ipv4--fragmentation"></a>
### Nested Schema for `application_profiles.ipv4.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--application_profiles--ipv4--section"></a>
### Nested Schema for `application_profiles.ipv4.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--application_profiles--ipv4--source"></a>
### Nested Schema for `application_profiles.ipv4.source`

Read-Only:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--application_profiles--ipv6"></a>
### Nested Schema for `application_profiles.ipv6`

Read-Only:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--application_profiles--ipv6--destination"></a>
### Nested Schema for `application_profiles.ipv6.destination`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--application_profiles--ipv6--fragmentation"></a>
### Nested Schema for `application_profiles.ipv6.fragmentation`

Read-Only:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--application_profiles--ipv6--length"></a>
### Nested Schema for `application_profiles.ipv6.length`

Read-Only:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--application_profiles--ipv6--section"></a>
### Nested Schema for `application_profiles.ipv6.section`

Read-Only:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--application_profiles--ipv6--source"></a>
### Nested Schema for `application_profiles.ipv6.source`

Read-Only:

* `prefix_min_mask` (String)
<a id="nestedatt--application_profiles--outer_ipv4"></a>
### Nested Schema for `application_profiles.outer_ipv4`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--application_profiles--outer_ipv6"></a>
### Nested Schema for `application_profiles.outer_ipv6`

Read-Only:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--application_profiles--timestamp"></a>
### Nested Schema for `application_profiles.timestamp`

Read-Only:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
* `sys_up_time_first` (Boolean)
* `sys_up_time_last` (Boolean)
<a id="nestedatt--application_profiles--transport"></a>
### Nested Schema for `application_profiles.transport`

Read-Only:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--application_profiles--transport--udp))
<a id="nestedatt--application_profiles--transport--icmp"></a>
### Nested Schema for `application_profiles.transport.icmp`

Read-Only:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--application_profiles--transport--tcp"></a>
### Nested Schema for `application_profiles.transport.tcp`

Read-Only:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--application_profiles--transport--udp"></a>
### Nested Schema for `application_profiles.transport.udp`

Read-Only:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)
<a id="nestedatt--context"></a>
### Nested Schema for `context`

Read-Only:

* `page_no` (Number) - page number of the returned result set
* `page_size` (Number) - page size of the returned result set
* `sort` (List of String) - sorting info of the returned result set. list of fields in the array indicate sorting order
* `total_items` (Number) - total number of items in the queried entity type

