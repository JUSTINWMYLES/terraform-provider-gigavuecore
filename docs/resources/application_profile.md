---
page_title: "gigavuecore_application_profile Resource - gigavuecore"
subcategory: ""
description: |-
  Create a Application Profile
---

# gigavuecore_application_profile Resource

Create a Application Profile

## Example Usage

```terraform
resource "gigavuecore_application_profile" "example" {
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
    in_name_width      = 1
    in_physical_width  = 2
    out_physical_width = 2
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
      header_size  = 1
      payload_size = 1
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
      header_size  = 1
      payload_size = 1
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
  type = "export"
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

### Nested Blocks

* `timeouts` (Block Single) (see [below for nested schema](#nestedatt--timeouts))

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
<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

* `create` (String) - A create timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `read` (String) - A read timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).
* `update` (String) - An update timeout for this operation, e.g. "20m0s". Overrides the generator default (20m0s).
* `delete` (String) - A delete timeout for this operation, e.g. "10m0s". Overrides the generator default (10m0s).

## Import

Import is supported using the following syntax:

```shell
terraform import gigavuecore_application_profile.example {alias}
```
