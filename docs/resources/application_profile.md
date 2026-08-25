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
  alias          = null
  application_id = null
  applications   = []
  counter        = {}
  datalink       = {}
  description    = null
  flow           = {}
  gtpu           = {}
  interface      = {}
  ip             = {}
  ipv4           = {}
  ipv6           = {}
  outer_ipv4     = {}
  outer_ipv6     = {}
  timestamp      = {}
  transport      = {}
  type           = null
}
```

## Schema

### Arguments

The following arguments are supported:

* `alias` (String, required) - application profile alias
* `application_id` (Boolean, optional) - only valid with 'export' type
* `applications` (Attributes List, optional) - application and attributes. (see [below for nested schema](#nestedatt--applications))
* `counter` (Attributes, optional) (see [below for nested schema](#nestedatt--counter))
* `datalink` (Attributes, optional) (see [below for nested schema](#nestedatt--datalink))
* `description` (String, optional)
* `flow` (Attributes, optional) (see [below for nested schema](#nestedatt--flow))
* `gtpu` (Attributes, optional) (see [below for nested schema](#nestedatt--gtpu))
* `interface` (Attributes, optional) (see [below for nested schema](#nestedatt--interface))
* `ip` (Attributes, optional) (see [below for nested schema](#nestedatt--ip))
* `ipv4` (Attributes, optional) (see [below for nested schema](#nestedatt--ipv4))
* `ipv6` (Attributes, optional) (see [below for nested schema](#nestedatt--ipv6))
* `outer_ipv4` (Attributes, optional) (see [below for nested schema](#nestedatt--outer_ipv4))
* `outer_ipv6` (Attributes, optional) (see [below for nested schema](#nestedatt--outer_ipv6))
* `timestamp` (Attributes, optional) (see [below for nested schema](#nestedatt--timestamp))
* `transport` (Attributes, optional) (see [below for nested schema](#nestedatt--transport))
* `type` (String, optional)

### Attributes

In addition to all arguments above, the following computed attributes are exported:

* `application_id` (Boolean, computed) - only valid with 'export' type
* `applications` (Attributes List, computed) - application and attributes. (see [below for nested schema](#nestedatt--applications))
* `counter` (Attributes, computed) (see [below for nested schema](#nestedatt--counter))
* `datalink` (Attributes, computed) (see [below for nested schema](#nestedatt--datalink))
* `description` (String, computed)
* `flow` (Attributes, computed) (see [below for nested schema](#nestedatt--flow))
* `gtpu` (Attributes, computed) (see [below for nested schema](#nestedatt--gtpu))
* `interface` (Attributes, computed) (see [below for nested schema](#nestedatt--interface))
* `ip` (Attributes, computed) (see [below for nested schema](#nestedatt--ip))
* `ipv4` (Attributes, computed) (see [below for nested schema](#nestedatt--ipv4))
* `ipv6` (Attributes, computed) (see [below for nested schema](#nestedatt--ipv6))
* `outer_ipv4` (Attributes, computed) (see [below for nested schema](#nestedatt--outer_ipv4))
* `outer_ipv6` (Attributes, computed) (see [below for nested schema](#nestedatt--outer_ipv6))
* `timestamp` (Attributes, computed) (see [below for nested schema](#nestedatt--timestamp))
* `transport` (Attributes, computed) (see [below for nested schema](#nestedatt--transport))
* `type` (String, computed)

<a id="nestedatt--applications"></a>
### Nested Schema for `applications`

Required:

* `name` (String) - application name
Optional:

* `attributes` (Attributes List) (see [below for nested schema](#nestedatt--applications--attributes))
* `is_user_defined` (Boolean) - Default value is false. Must be set to true only for App Intel solution while configuring user defined apps
<a id="nestedatt--applications--attributes"></a>
### Nested Schema for `applications.attributes`

Required:

* `name` (String) - attribute name
Optional:

* `value` (String) - application's attribute value
<a id="nestedatt--counter"></a>
### Nested Schema for `counter`

Optional:

* `bytes` (Boolean)
* `bytes_long` (Boolean)
* `inner_byte` (Boolean)
* `inner_byte_long` (Boolean)
* `packets` (Boolean)
* `packets_long` (Boolean)
<a id="nestedatt--datalink"></a>
### Nested Schema for `datalink`

Optional:

* `mac_dst` (Boolean)
* `mac_src` (Boolean)
* `vlan` (Boolean)
<a id="nestedatt--flow"></a>
### Nested Schema for `flow`

Optional:

* `end_reason` (Boolean)
<a id="nestedatt--gtpu"></a>
### Nested Schema for `gtpu`

Optional:

* `qfi` (Boolean)
* `teid` (Boolean)
<a id="nestedatt--interface"></a>
### Nested Schema for `interface`

Optional:

* `in_name_width` (Number)
* `in_physical_width` (Number)
* `out_physical_width` (Number)
<a id="nestedatt--ip"></a>
### Nested Schema for `ip`

Optional:

* `version` (Boolean)
<a id="nestedatt--ipv4"></a>
### Nested Schema for `ipv4`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--ipv4--destination))
* `dscp` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--ipv4--fragmentation))
* `header_len` (Boolean)
* `option_map` (Boolean)
* `precedence` (Boolean)
* `protocol` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--ipv4--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--ipv4--source))
* `tos` (Boolean)
* `total_length` (Boolean)
* `ttl` (Boolean)
<a id="nestedatt--ipv4--destination"></a>
### Nested Schema for `ipv4.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--ipv4--fragmentation"></a>
### Nested Schema for `ipv4.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--ipv4--section"></a>
### Nested Schema for `ipv4.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--ipv4--source"></a>
### Nested Schema for `ipv4.source`

Optional:

* `prefix_min_mask` (String) - ipv4 source prefix minimum-mask - netmask or mask length
<a id="nestedatt--ipv6"></a>
### Nested Schema for `ipv6`

Optional:

* `destination` (Attributes) (see [below for nested schema](#nestedatt--ipv6--destination))
* `dscp` (Boolean)
* `extension_map` (Boolean)
* `flow_label` (Boolean)
* `fragmentation` (Attributes) (see [below for nested schema](#nestedatt--ipv6--fragmentation))
* `hop_limit` (Boolean)
* `length` (Attributes) (see [below for nested schema](#nestedatt--ipv6--length))
* `next_header` (Boolean)
* `precedence` (Boolean)
* `section` (Attributes) (see [below for nested schema](#nestedatt--ipv6--section))
* `source` (Attributes) (see [below for nested schema](#nestedatt--ipv6--source))
* `traffic_class` (Boolean)
<a id="nestedatt--ipv6--destination"></a>
### Nested Schema for `ipv6.destination`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--ipv6--fragmentation"></a>
### Nested Schema for `ipv6.fragmentation`

Optional:

* `flags` (Boolean)
* `offset` (Boolean)
<a id="nestedatt--ipv6--length"></a>
### Nested Schema for `ipv6.length`

Optional:

* `header` (Boolean)
* `payload` (Boolean)
* `total` (Boolean)
<a id="nestedatt--ipv6--section"></a>
### Nested Schema for `ipv6.section`

Optional:

* `header_size` (Number)
* `payload_size` (Number)
<a id="nestedatt--ipv6--source"></a>
### Nested Schema for `ipv6.source`

Optional:

* `prefix_min_mask` (String)
<a id="nestedatt--outer_ipv4"></a>
### Nested Schema for `outer_ipv4`

Optional:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--outer_ipv6"></a>
### Nested Schema for `outer_ipv6`

Optional:

* `destination` (Boolean)
* `source` (Boolean)
<a id="nestedatt--timestamp"></a>
### Nested Schema for `timestamp`

Optional:

* `flow_end_msec` (Boolean)
* `flow_endsec` (Boolean)
* `flow_start_msec` (Boolean)
* `flow_startsec` (Boolean)
* `sys_up_time_first` (Boolean)
* `sys_up_time_last` (Boolean)
<a id="nestedatt--transport"></a>
### Nested Schema for `transport`

Optional:

* `dst_port` (Boolean)
* `icmp` (Attributes) (see [below for nested schema](#nestedatt--transport--icmp))
* `src_port` (Boolean)
* `tcp` (Attributes) (see [below for nested schema](#nestedatt--transport--tcp))
* `udp` (Attributes) (see [below for nested schema](#nestedatt--transport--udp))
<a id="nestedatt--transport--icmp"></a>
### Nested Schema for `transport.icmp`

Optional:

* `ipv4_code` (Boolean)
* `ipv4_type` (Boolean)
* `ipv6_code` (Boolean)
* `ipv6_type` (Boolean)
<a id="nestedatt--transport--tcp"></a>
### Nested Schema for `transport.tcp`

Optional:

* `ack_number` (Boolean)
* `dst_port` (Boolean)
* `flags` (Boolean)
* `header_len` (Boolean)
* `seq_number` (Boolean)
* `src_port` (Boolean)
* `urgent_ptr` (Boolean)
* `window_size` (Boolean)
<a id="nestedatt--transport--udp"></a>
### Nested Schema for `transport.udp`

Optional:

* `dst_port` (Boolean)
* `msg_len` (Boolean)
* `src_port` (Boolean)

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_application_profile.example {alias}
```
